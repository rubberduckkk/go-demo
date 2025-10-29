package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// HashFunc 定义了哈希函数类型
type HashFunc func(data []byte) uint32

// Map 一致性哈希结构
type Map struct {
	hash     HashFunc       // 哈希函数
	replicas int            // 每个节点的虚拟节点数
	keys     []int          // 哈希环（存储所有虚拟节点的哈希值并排序）
	hashMap  map[int]string // 虚拟节点到真实节点的映射
	mu       sync.RWMutex
}

// New 创建一个一致性哈希实例
func New(replicas int, fn HashFunc) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 添加真实节点
func (m *Map) Add(nodes ...string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, node := range nodes {
		for i := 0; i < m.replicas; i++ {
			// 虚拟节点名 = 节点名 + 序号
			hash := int(m.hash([]byte(strconv.Itoa(i) + node)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = node
		}
	}
	sort.Ints(m.keys)
}

// Get 根据 key 获取对应的节点
func (m *Map) Get(key string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	// 二分查找顺时针第一个大于 consistenthash 的虚拟节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	// 如果超过了最大值，回绕到第一个节点
	if idx == len(m.keys) {
		idx = 0
	}
	return m.hashMap[m.keys[idx]]
}

// Delete 删除节点（可选）
func (m *Map) Delete(nodes ...string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, node := range nodes {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + node)))
			delete(m.hashMap, hash)
			// 删除 keys 切片中的 consistenthash
			for j, k := range m.keys {
				if k == hash {
					m.keys = append(m.keys[:j], m.keys[j+1:]...)
					break
				}
			}
		}
	}
	sort.Ints(m.keys)
}

func example() {
	ch := New(3, nil)
	ch.Add("NodeA", "NodeB", "NodeC")

	keys := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	for _, key := range keys {
		fmt.Printf("%s => %s\n", key, ch.Get(key))
	}
}
