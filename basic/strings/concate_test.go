package strings

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var (
	letters        = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	benchmarkCases = []int{5, 10, 20, 50, 100, 200, 500}
)

func randSeq(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}

/*
	BenchmarkConcate/case_5
	BenchmarkConcate/case_5-11         	125744928	         9.436 ns/op
	BenchmarkConcate/case_10
	BenchmarkConcate/case_10-11        	136409209	         8.760 ns/op
	BenchmarkConcate/case_20
	BenchmarkConcate/case_20-11        	60157286	        19.24 ns/op
	BenchmarkConcate/case_50
	BenchmarkConcate/case_50-11        	55018215	        21.75 ns/op
	BenchmarkConcate/case_100
	BenchmarkConcate/case_100-11       	44455627	        27.65 ns/op
	BenchmarkConcate/case_200
	BenchmarkConcate/case_200-11       	28230920	        42.22 ns/op
	BenchmarkConcate/case_500
	BenchmarkConcate/case_500-11       	13327545	        89.72 ns/op
*/
func BenchmarkConcate(b *testing.B) {
	for _, n := range benchmarkCases {
		b.Run("case "+strconv.Itoa(n), func(b *testing.B) {
			str1, str2 := string(randSeq(n)), string(randSeq(n))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = str1 + str2
			}
		})
	}
}

/*
	BenchmarkFormat/case_5
	BenchmarkFormat/case_5-11         	15018075	        71.04 ns/op
	BenchmarkFormat/case_10
	BenchmarkFormat/case_10-11        	16916617	        70.22 ns/op
	BenchmarkFormat/case_20
	BenchmarkFormat/case_20-11        	16968574	        71.52 ns/op
	BenchmarkFormat/case_50
	BenchmarkFormat/case_50-11        	16100709	        76.81 ns/op
	BenchmarkFormat/case_100
	BenchmarkFormat/case_100-11       	15213085	        80.57 ns/op
	BenchmarkFormat/case_200
	BenchmarkFormat/case_200-11       	12717628	        96.94 ns/op
	BenchmarkFormat/case_500
	BenchmarkFormat/case_500-11       	 8677011	       139.0 ns/op
*/
func BenchmarkFormat(b *testing.B) {
	for _, n := range benchmarkCases {
		b.Run("case "+strconv.Itoa(n), func(b *testing.B) {
			str1, str2 := string(randSeq(n)), string(randSeq(n))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = fmt.Sprintf("%s%s", str1, str2)
			}
		})
	}
}
