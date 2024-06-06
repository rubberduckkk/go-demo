package hash

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
)

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func CRC32(str []byte) uint32 {
	return crc32.ChecksumIEEE(str)
}
