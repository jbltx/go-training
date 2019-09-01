package main

import (
	"fmt"
)

func main() {
	var num uint32 = 43261596
	result := reverseBits(num)
	fmt.Printf("%v\n", result)
}