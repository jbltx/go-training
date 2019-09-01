package main

func removeElement(nums []int, val int) int {
    
    result := 0
    valLength := len(nums)
    
    for i := 0; i < valLength; i++ {
        if nums[i] != val {
            nums[result] = nums[i]
            result++
        }
    }
    
    return result
}