package hash

import (
	"fmt"
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

func TestCRC32(t *testing.T) {
	sequence1 := randSeq(1024)
	res1 := CRC32(sequence1)
	sequence2 := randSeq(1024)
	res2 := CRC32(sequence2)
	fmt.Printf("res1: %x, res2: %x\n", res1, res2)
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
