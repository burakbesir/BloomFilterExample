package bloom_filter

import (
	"github.com/spaolacci/murmur3"
	"hash/fnv"
)

const vectorSize = 1_000_000

var bitVector [vectorSize]int

func Search(s string) bool {
	murmurIndex := murmurHashIndex(s)
	fnvIndex := fnvHashIndex(s)

	return bitVector[murmurIndex] == 1 && bitVector[fnvIndex] == 1
}

func Insert(s string) {
	murmurIndex := murmurHashIndex(s)
	fnvIndex := fnvHashIndex(s)

	bitVector[murmurIndex] = 1
	bitVector[fnvIndex] = 1
}

func murmurHashIndex(s string) uint64 {
	m64 := murmur3.New64()
	_, err := m64.Write([]byte(s))

	if err != nil {
		panic(err)
	}

	return m64.Sum64() % vectorSize
}

func fnvHashIndex(s string) uint64 {
	f64 := fnv.New64()

	_, err := f64.Write([]byte(s))
	if err != nil {
		panic(err)
	}

	return f64.Sum64() % vectorSize
}
