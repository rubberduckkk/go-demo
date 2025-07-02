package hash

import (
	"crypto/md5"
	"encoding/binary"
	"sort"
	"strconv"
)

type HashFunc func([]byte) uint32

const (
	DefaultReplica = 100
)

type Ketama struct {
	hash     HashFunc
	replicas int
	keys     []int
	hashmap  map[int]string
}

func (k *Ketama) Add(nodes ...*Node) {
	for _, n := range nodes {
		for i := 0; i < k.replicas; i++ {
			key := int(k.hash([]byte(strconv.Itoa(i) + "_" + n.Addr)))
			if _, ok := k.hashmap[key]; !ok {
				k.keys = append(k.keys, key)
			}
			k.hashmap[key] = n.Addr
		}
	}
	sort.Ints(k.keys)
}

func (k *Ketama) Rebuild(nodes []*Node) {
	k.Clear()
	k.Add(nodes...)
}

func (k *Ketama) Get(key string) (addr string, ok bool) {
	if k.IsEmpty() {
		return "", false
	}

	h := int(k.hash([]byte(key)))
	idx := sort.Search(len(k.keys), func(i int) bool {
		return k.keys[i] >= h
	})

	if idx == len(k.keys) {
		idx = 0
	}

	addr, ok = k.hashmap[k.keys[idx]]
	return addr, ok
}

func (k *Ketama) Remove(addrs ...string) {
	deleteKeys := make([]int, 0)
	for _, addr := range addrs {
		for i := 0; i < k.replicas; i++ {
			key := int(k.hash([]byte(strconv.Itoa(i) + "_" + addr)))
			if _, ok := k.hashmap[key]; ok {
				deleteKeys = append(deleteKeys, key)
				delete(k.hashmap, key)
			}
		}
	}

	k.deleteKeys(deleteKeys)
}

func (k *Ketama) deleteKeys(keys []int) {
	if len(keys) == 0 {
		return
	}

	// TODO: finish implementation
}

func (k *Ketama) Clear() {
	k.keys = nil
	k.hashmap = make(map[int]string)
}

func (k *Ketama) IsEmpty() bool {
	return len(k.keys) == 0
}

func NewKetama(replicas int, fn HashFunc) *Ketama {
	h := &Ketama{
		hash:     fn,
		replicas: replicas,
		hashmap:  make(map[int]string),
	}

	if h.replicas <= 0 {
		h.replicas = DefaultReplica
	}

	if h.hash == nil {
		h.hash = MD5Ketama
	}
	return h
}

func MD5Ketama(data []byte) uint32 {
	s := md5.Sum(data)
	return binary.LittleEndian.Uint32(s[:4])
}
