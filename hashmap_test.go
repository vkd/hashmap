package hashmap

import (
	"reflect"
	"testing"
)

func TestHashMap_Set(t *testing.T) {
	tests := []struct {
		name    string
		key     Key
		value   interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base", "key", 5, false},
		{"base", "key", 6, false},
		{"base", "key", 7, false},
		{"base", "key0", 8, false},
		{"base", "key1", 9, false},
		{"base", "key2", 10, false},
		{"base", "key3", 11, false},
	}
	h := NewHashMap(0, zeroHash)
	for _, tt := range tests {
		if err := h.Set(tt.key, tt.value); (err != nil) != tt.wantErr {
			t.Errorf("HashMap.Set() error = %v, wantErr %v", err, tt.wantErr)
		}
		v, err := h.Get(tt.key)
		if err != nil {
			t.Errorf("Error on get key: %v", err)
		}
		if !reflect.DeepEqual(v, tt.value) {
			t.Errorf("HashMap.Get() = %v, want = %v", v, tt.value)
		}
	}

	if h.Count() != 5 {
		t.Errorf("Wrong count hashmap: %d (want: %d)", h.Count(), 5)
	}
}

func zeroHash(uint, Key) (uint, error) {
	return 0, nil
}

func TestHashMap_Get(t *testing.T) {
	tests := []struct {
		name    string
		key     Key
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"first", "1", 1, false},
		{"last", "5", 5, false},
		{"not found", "6", 6, true},
	}
	h := NewHashMap(0, zeroHash)
	h.Set("1", 1)
	h.Set("2", 2)
	h.Set("3", 3)
	h.Set("4", 4)
	h.Set("5", 5)

	for _, tt := range tests {
		got, err := h.Get(tt.key)
		if (err != nil) != tt.wantErr {
			t.Errorf("Error on key: %v", tt.key)
			t.Errorf("HashMap.Get() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if err != nil {
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("HashMap.Get() = %v, want %v", got, tt.want)
		}
	}
}

func TestHashMap_Unset(t *testing.T) {
	tests := []struct {
		name    string
		key     Key
		wantErr bool
	}{
		// TODO: Add test cases.
		{"first", "1", false},
		{"first second", "1", true},
		{"last", "5", false},
		{"last second", "5", true},
		{"not found", "6", true},
	}

	h := NewHashMap(0, zeroHash)
	h.Set("1", 1)
	h.Set("2", 2)
	h.Set("3", 3)
	h.Set("4", 4)
	h.Set("5", 5)

	for _, tt := range tests {
		if err := h.Unset(tt.key); (err != nil) != tt.wantErr {
			t.Errorf("HashMap.Unset() error = %v, wantErr %v", err, tt.wantErr)
		}
	}

	if h.Count() != 3 {
		t.Errorf("Wrong hashmap count: %d (want: %d)", h.Count(), 3)
	}
}
