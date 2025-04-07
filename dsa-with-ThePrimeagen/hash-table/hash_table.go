package hashtable

func BucketIndex(key string, size int) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % size
}

type keyValue struct {
	Key   string
	Value string
}

type hashTable struct {
	buckets [][]keyValue
}

func New(size int) *hashTable {
	return &hashTable{
		buckets: make([][]keyValue, size),
	}
}

func (ht *hashTable) Insert(key, value string) {
	hash := BucketIndex(key, len(ht.buckets))
	for i, kv := range ht.buckets[hash] {
		if kv.Key == key {
			ht.buckets[hash][i].Value = value
			return
		}
	}

	ht.buckets[hash] = append(ht.buckets[hash], keyValue{Key: key, Value: value})
}

func (ht *hashTable) Get(key string) (string, bool) {
	hash := BucketIndex(key, len(ht.buckets))
	for _, kv := range ht.buckets[hash] {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return "", false
}
