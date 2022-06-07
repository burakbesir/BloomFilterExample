package main

import (
	"BloomFilterExample/bloom_filter"
	"fmt"
)

func main() {

	bloom_filter.Insert("11111")
	bloom_filter.Insert("22222")
	bloom_filter.Insert("33333")

	fmt.Println(bloom_filter.Search("11111"))
	fmt.Println(bloom_filter.Search("2222"))

}
