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
