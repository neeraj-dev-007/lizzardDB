package storage

import "fmt"

type KeyValueStore struct {
	store map[string]string
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{store: make(map[string]string)}
}

func (kv *KeyValueStore) Insert(key, value string) {
	kv.store[key] = value
	fmt.Println("Inserted: ", key, value)
}

func (kv *KeyValueStore) Select(key string) {
	value, exists := kv.store[key]
	if(exists) {
		fmt.Println("Value: ", value)
	} else{
		fmt.Println("Key not found")
	}
}

func (kv *KeyValueStore) Delete(key string) {
	_, exists := kv.store[key]
	if(exists) {
		delete(kv.store, key)
		fmt.Println("Deleted: ", key)
	} else {
		fmt.Println("Key not found")
	}
}