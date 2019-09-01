package main

import "sort"

func heightChecker(heights []int) int {
    
    count := len(heights)
    sorted := make([]int, count)
    copy(sorted, heights)
    sort.Ints(sorted)
    
    result := 0
    
    for i := 0; i < count ; i++ {
        if sorted[i] != heights[i] {
            result++
        }
    }
    
    return result
}