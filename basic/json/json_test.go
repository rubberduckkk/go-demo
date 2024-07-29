package json

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"

	jsoniter "github.com/json-iterator/go"
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

func genTestSlice(sliceLen int, itemLen int) []string {
	s := make([]string, sliceLen)
	for i := range s {
		s[i] = string(randSeq(itemLen))
	}
	return s
}

/*
	BenchmarkEncodingJSON/case_5
	BenchmarkEncodingJSON/case_5-11         	 4573008	       230.3 ns/op
	BenchmarkEncodingJSON/case_10
	BenchmarkEncodingJSON/case_10-11        	 3086535	       387.0 ns/op
	BenchmarkEncodingJSON/case_20
	BenchmarkEncodingJSON/case_20-11        	 1698301	       705.7 ns/op
	BenchmarkEncodingJSON/case_50
	BenchmarkEncodingJSON/case_50-11        	  717511	      1652 ns/op
	BenchmarkEncodingJSON/case_100
	BenchmarkEncodingJSON/case_100-11       	  371694	      3337 ns/op
	BenchmarkEncodingJSON/case_200
	BenchmarkEncodingJSON/case_200-11       	  186453	      6417 ns/op
	BenchmarkEncodingJSON/case_500
	BenchmarkEncodingJSON/case_500-11       	   76377	     15436 ns/op
*/
func BenchmarkEncodingJSON(b *testing.B) {
	for _, n := range benchmarkCases {
		b.Run("case "+strconv.Itoa(n), func(b *testing.B) {
			data := genTestSlice(n, 36)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = json.Marshal(data)
			}
		})
	}
}

/*
	BenchmarkDecodingJSON/case_5
	BenchmarkDecodingJSON/case_5-11         	 1064298	      1078 ns/op
	BenchmarkDecodingJSON/case_10
	BenchmarkDecodingJSON/case_10-11        	  636854	      1929 ns/op
	BenchmarkDecodingJSON/case_20
	BenchmarkDecodingJSON/case_20-11        	  347272	      3434 ns/op
	BenchmarkDecodingJSON/case_50
	BenchmarkDecodingJSON/case_50-11        	  148662	      8108 ns/op
	BenchmarkDecodingJSON/case_100
	BenchmarkDecodingJSON/case_100-11       	   76050	     15889 ns/op
	BenchmarkDecodingJSON/case_200
	BenchmarkDecodingJSON/case_200-11       	   38496	     31084 ns/op
	BenchmarkDecodingJSON/case_500
	BenchmarkDecodingJSON/case_500-11       	   15558	     76981 ns/op
*/
func BenchmarkDecodingJSON(b *testing.B) {
	for _, n := range benchmarkCases {
		b.Run("case "+strconv.Itoa(n), func(b *testing.B) {
			data := genTestSlice(n, 36)
			raw, _ := json.Marshal(data)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var data []string
				_ = json.Unmarshal(raw, &data)
			}
		})
	}
}

/*
	BenchmarkEncodingJSONIter/case_5
	BenchmarkEncodingJSONIter/case_5-11         	 5274819	       207.4 ns/op
	BenchmarkEncodingJSONIter/case_10
	BenchmarkEncodingJSONIter/case_10-11        	 3438835	       348.9 ns/op
	BenchmarkEncodingJSONIter/case_20
	BenchmarkEncodingJSONIter/case_20-11        	 1869126	       643.9 ns/op
	BenchmarkEncodingJSONIter/case_50
	BenchmarkEncodingJSONIter/case_50-11        	  781975	      1496 ns/op
	BenchmarkEncodingJSONIter/case_100
	BenchmarkEncodingJSONIter/case_100-11       	  405064	      2912 ns/op
	BenchmarkEncodingJSONIter/case_200
	BenchmarkEncodingJSONIter/case_200-11       	  204106	      5808 ns/op
	BenchmarkEncodingJSONIter/case_500
	BenchmarkEncodingJSONIter/case_500-11       	   85036	     14032 ns/op
*/
func BenchmarkEncodingJSONIter(b *testing.B) {
	for _, n := range benchmarkCases {
		b.Run("case "+strconv.Itoa(n), func(b *testing.B) {
			data := genTestSlice(n, 36)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = jsoniter.Marshal(data)
			}
		})
	}
}

/*
	BenchmarkDecodingJSONIter/case_5
	BenchmarkDecodingJSONIter/case_5-11         	 3418740	       327.5 ns/op
	BenchmarkDecodingJSONIter/case_10
	BenchmarkDecodingJSONIter/case_10-11        	 2159852	       552.3 ns/op
	BenchmarkDecodingJSONIter/case_20
	BenchmarkDecodingJSONIter/case_20-11        	 1221991	       965.9 ns/op
	BenchmarkDecodingJSONIter/case_50
	BenchmarkDecodingJSONIter/case_50-11        	  557569	      2093 ns/op
	BenchmarkDecodingJSONIter/case_100
	BenchmarkDecodingJSONIter/case_100-11       	  303124	      3912 ns/op
	BenchmarkDecodingJSONIter/case_200
	BenchmarkDecodingJSONIter/case_200-11       	  156946	      7586 ns/op
	BenchmarkDecodingJSONIter/case_500
	BenchmarkDecodingJSONIter/case_500-11       	   65990	     18179 ns/op
*/
func BenchmarkDecodingJSONIter(b *testing.B) {
	for _, n := range benchmarkCases {
		b.Run("case "+strconv.Itoa(n), func(b *testing.B) {
			data := genTestSlice(n, 36)
			raw, _ := jsoniter.Marshal(data)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var data []string
				_ = jsoniter.Unmarshal(raw, &data)
			}
		})
	}
}
