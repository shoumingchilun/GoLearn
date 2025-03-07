package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < 1000; i++ {
		num := uint64(rand.Int63())
		PopCount(num)
	}
}
func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < 1000; i++ {
		num := uint64(rand.Int63())
		PopCount2(num)
	}

}
