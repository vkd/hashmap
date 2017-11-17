package hashmap

// NativeMap - hashmap of native go map
type NativeMap struct {
	m map[interface{}]interface{}
}

var _ HashMaper = (*NativeMap)(nil)

// NewNativeMap - create native hashmap
func NewNativeMap() HashMaper {
	return &NativeMap{
		m: make(map[interface{}]interface{}),
	}
}

// Set - set new key|value pair
func (n *NativeMap) Set(key Key, value interface{}) error {
	n.m[key] = value
	return nil
}

// Get - get value by key
func (n *NativeMap) Get(key Key) (interface{}, error) {
	v, ok := n.m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	return v, nil
}

// Unset - delete value by key
func (n *NativeMap) Unset(key Key) error {
	delete(n.m, key)
	return nil
}

// Count - return count keys
func (n *NativeMap) Count() int {
	return len(n.m)
}
