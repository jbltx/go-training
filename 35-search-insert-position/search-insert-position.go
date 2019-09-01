package main

func searchInsert(nums []int, target int) int {
    count := len(nums)
    for i:= 0; i < count; i++ {
        if target <= nums[i] {
            return i
        }
    }
    return count
}