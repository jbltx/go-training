package main

func singleNumber(nums []int) int {
    count := len(nums)
    result := 0
    
    for i := 0; i < count; i++ {
        result ^= nums[i]
    }
    
    return result
}