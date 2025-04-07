package hashtable_test

import (
	hashtable "dsa-with-ThePrimeagen/hash-table"
	"testing"
)

func TestHashTable(t *testing.T) {
	hashTable := hashtable.New(10)

	hashTable.Insert("key1", "value1")
	hashTable.Insert("key1", "value3")
	hashTable.Insert("key2", "value2")

	value, _ := hashTable.Get("key1")

	if value != "value3" {
		t.Fatalf("expected %s rec %s", "value1", value)
	}
}
