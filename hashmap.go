package hashmap

// HashMap - implement HashMap structure
type HashMap struct {
	blockSize uint
	hashFn    HashFunc

	count   int
	buckets []*entry
}

var _ HashMaper = (*HashMap)(nil)

// NewHashMap - create new hashmap
//
// blockSize - count of buckets (default 128)
// fn - hash func
func NewHashMap(blockSize uint, fn HashFunc) HashMaper {
	if blockSize <= 0 {
		blockSize = 128
	}
	if fn == nil {
		fn = defaultHashFunc
	}
	e := &HashMap{
		blockSize: blockSize,
		hashFn:    fn,
	}
	e.buckets = make([]*entry, blockSize)
	return e
}

// Set - set value by key
func (h *HashMap) Set(key Key, value interface{}) error {
	hash, err := h.hashFn(h.blockSize, key)
	if err != nil {
		return err
	}
	bindex := h.backetIndex(hash)

	etr := h.buckets[bindex]
	if etr == nil {
		h.count++
		h.buckets[bindex] = &entry{key: key, val: value}
		return nil
	}

	var prev *entry
	for etr != nil {
		// key found - update value
		if equals(etr.key, key) {
			etr.val = value
			return nil
		}

		prev = etr
		etr = etr.next
	}

	// key not found - new value
	h.count++
	etr = &entry{key: key, val: value}
	prev.next = etr
	return nil
}

// Get - get value by key
func (h *HashMap) Get(key Key) (interface{}, error) {
	hash, err := h.hashFn(h.blockSize, key)
	if err != nil {
		return nil, err
	}
	bindex := h.backetIndex(hash)

	etr := h.buckets[bindex]
	for etr != nil {
		// key found
		if equals(etr.key, key) {
			return etr.val, nil
		}

		etr = etr.next
	}

	// key not found
	return nil, ErrKeyNotFound
}

// Unset - delete value by key
func (h *HashMap) Unset(key Key) error {
	hash, err := h.hashFn(h.blockSize, key)
	if err != nil {
		return err
	}
	bindex := h.backetIndex(hash)

	etr := h.buckets[bindex]
	if etr == nil {
		return ErrKeyNotFound
	}
	if equals(etr.key, key) {
		h.buckets[bindex] = etr.next
		h.count--
		return nil
	}
	var prev = etr
	etr = etr.next

	for etr != nil {
		// key found
		if equals(etr.key, key) {
			prev.next = etr.next
			h.count--
			return nil
		}

		prev = etr
		etr = etr.next
	}

	return ErrKeyNotFound
}

// Count - count values in HashMap
func (h *HashMap) Count() int {
	return h.count
}

func (h *HashMap) backetIndex(hash uint) uint {
	return hash % h.blockSize
}

func equals(a, b interface{}) bool {
	return a == b
}

type entry struct {
	key  Key
	val  interface{}
	next *entry
}
