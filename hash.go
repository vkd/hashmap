package hashmap

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"hash/crc32"
)

var (
	defaultHashFunc HashFunc = hash
)

func hash(size uint, s Key) (uint, error) {
	var bs []byte
	switch s := s.(type) {
	case string:
		bs = []byte(s)
	case int:
		return uint(s), nil
	case uint:
		return s, nil
	default:
		var err error
		bs, err = serialize(reflect.ValueOf(s))
		if err != nil {
			return 0, fmt.Errorf("error on serialize: %v", err)
		}
	}

	// f := fnv.New64a()
	// f.Write([]byte(s))
	// return uint(f.Sum64())

	// return uint(crc64.Checksum([]byte(s), crc64.MakeTable(1000)))
	return uint(crc32.ChecksumIEEE(bs)), nil
}

func serialize(v reflect.Value) ([]byte, error) {
	var bs bytes.Buffer
	err := objToBytes(&bs, v)
	if err != nil {
		return nil, err
	}
	return bs.Bytes(), nil
}

func objToBytes(bs *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		bs.WriteString("S\"" + v.String() + "\"")
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		bs.WriteString("I\"" + strconv.FormatInt(v.Int(), 10) + "\"")
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		bs.WriteString("UI\"" + strconv.FormatUint(v.Uint(), 10) + "\"")
	case reflect.Float32, reflect.Float64:
		bs.WriteString("F\"" + strconv.FormatFloat(v.Float(), 'f', 0, 64) + "\"")
	case reflect.Bool:
		if v.Bool() {
			bs.WriteString("B\"True\"")
		} else {
			bs.WriteString("B\"False\"")
		}
	case reflect.Ptr:
		if v.IsNil() {
			bs.WriteString("nil")
		} else {
			bs.WriteRune('*')
			err := objToBytes(bs, v.Elem())
			if err != nil {
				return fmt.Errorf("error on serialize ptr (name: %s): %v", v.Type().Name(), err)
			}
		}
	case reflect.Array, reflect.Slice:
		bs.WriteRune('[')
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				bs.WriteRune(',')
			}
			err := objToBytes(bs, v.Index(i))
			if err != nil {
				return fmt.Errorf("error on serialize array/slice (name: %s): %v", v.Type().Name(), err)
			}
		}
		bs.WriteRune(']')
	case reflect.Map:
		keys := v.MapKeys()
		items := make([]mapItem, len(keys))
		var err error
		var raw []byte
		for i, key := range keys {
			raw, err = serialize(key)
			if err != nil {
				return fmt.Errorf("error on serialize map key (raw: %v), %v", key, err)
			}
			items[i].key = string(raw)
			items[i].value = v.MapIndex(key)
		}

		sort.Sort(sortNameItems(items))

		bs.WriteString("M{")
		for _, item := range items {
			bs.WriteRune('"')
			bs.WriteString(item.key)
			bs.WriteString("\":")
			err = objToBytes(bs, item.value)
			if err != nil {
				return fmt.Errorf("error on serialize map value (key: %s): %v", item.key, err)
			}
		}
		bs.WriteRune('}')
	case reflect.Struct:
		tp := v.Type()
		fieldsLen := tp.NumField()

		items := make([]mapItem, 0, fieldsLen)
		var field reflect.StructField
		for i := 0; i < fieldsLen; i++ {
			field = tp.Field(i)
			items = append(items, mapItem{
				key:   field.Name,
				value: v.Field(i),
			})
		}

		sort.Sort(sortNameItems(items))

		var err error

		bs.WriteString("S{")
		for _, item := range items {
			bs.WriteRune('"')
			bs.WriteString(item.key)
			bs.WriteString("\":")
			err = objToBytes(bs, item.value)
			if err != nil {
				return fmt.Errorf("error on serialize struct value (key: %s): %v", item.key, err)
			}
		}
		bs.WriteRune('}')
	case reflect.Interface:
		bs.WriteString("IFACE(")
		err := objToBytes(bs, reflect.ValueOf(v.Interface()))
		if err != nil {
			return fmt.Errorf("error on serialize interface: %v", err)
		}
		bs.WriteRune(')')
	default:
		// case reflect.Uintptr, reflect.Complex64, reflect.Complex128, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		bs.WriteString("UNKNOWN\"" + v.String() + "\"")
	}
	return nil
}

type mapItem struct {
	key   string
	value reflect.Value
}

type sortNameItems []mapItem

var _ sort.Interface = (sortNameItems)(nil)

func (s sortNameItems) Len() int           { return len(s) }
func (s sortNameItems) Less(i, j int) bool { return s[i].key < s[j].key }
func (s sortNameItems) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
