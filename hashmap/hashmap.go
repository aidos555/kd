package hashmap

import (
	"bytes"
	"hash/fnv"
)

var defaultCap uint64 = 1 << 10

type node struct {
	Key   []byte
	Value int
}

type hashMap struct {
	Capacity uint64
	Size     uint64
	Table    []*node
}

func newNode(key []byte, value int) *node {
	return &node{
		Key:   key,
		Value: value,
	}
}

func NewHashMap() *hashMap {
	return &hashMap{
		Capacity: defaultCap,
		Table:    make([]*node, defaultCap),
	}
}

func (hm *hashMap) Get(key []byte) int {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.Value
	}

	return 0
}

func (hm *hashMap) getNodeByHash(hash uint64) *node {
	return hm.Table[hash]
}

func (hm *hashMap) Set(key []byte, value int) int {
	return hm.setValue(hm.hash(key), key, value)
}

func (hm *hashMap) setValue(hash uint64, key []byte, value int) int {
	node := hm.getNodeByHash(hash)

	if node == nil {
		hm.Table[hash] = newNode(key, value)
	} else if bytes.Equal(node.Key, key) {
		hm.Table[hash].Value = value
		return value
	} else {
		hm.resize()
		hm.setValue(hash, key, value)
	}

	hm.Size++
	return value
}

func (hm *hashMap) Contains(key []byte) bool {
	node := hm.getNodeByHash(hm.hash(key))
	if node != nil {
		return true
	}
	return false
}

func (hm *hashMap) resize() {
	hm.Capacity <<= 1

	tempTable := hm.Table
	hm.Table = make([]*node, hm.Capacity)

	for i := 0; i < len(tempTable); i++ {
		node := tempTable[i]
		if node == nil {
			continue
		}

		hm.Table[hm.hash(node.Key)] = node
	}
}

func (hm *hashMap) hash(key []byte) uint64 {
	h := fnv.New32()
	h.Write(key)
	return uint64(h.Sum32()) % hm.Capacity
}
