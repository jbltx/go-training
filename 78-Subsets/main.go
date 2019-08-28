package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,3}
	result := subsets(nums)
	fmt.Printf("%v \n", result)
}