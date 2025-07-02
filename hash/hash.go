package hash

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"hash/crc64"
)

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func CRC32(str []byte) uint32 {
	return crc32.ChecksumIEEE(str)
}

func CRC64(str []byte) uint64 {
	return crc64.Checksum(str, crc64.MakeTable(crc64.ISO))
}

func HashCode(src string, bucketCount uint32) int {
	var hash uint32 = 0
	length := len(src)
	for i := 0; i < length; i++ {
		if i&1 == 0 {
			hash ^= (hash << 7) ^ uint32(src[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint32(src[i]) ^ (hash >> 5))
		}
	}
	return int(hash % bucketCount)
}

type Node struct {
	Addr   string
	Weight int
}

type Hash interface {
	Add(nodes ...*Node)
	Rebuild(nodes []*Node)
	Get(key string) (addr string, ok bool)
	Remove(addr ...string)
	Clear()
}
