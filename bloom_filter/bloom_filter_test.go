package bloom_filter

import "testing"

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insert("11111111_0000000")
	}
}
