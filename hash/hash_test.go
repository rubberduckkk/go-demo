package hash

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}

func BenchmarkMD5(b *testing.B) {
	sequence := randSeq(1024 * 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MD5(sequence)
	}
}

func BenchmarkCRC32(b *testing.B) {
	sequence := randSeq(1024 * 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CRC32(sequence)
	}
}

func BenchmarkCRC64(b *testing.B) {
	sequence := randSeq(1024 * 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CRC64(sequence)
	}
}
