package main

func maxSubArray(nums []int) int {
    
    count := len(nums)
    
    if count == 0 {
        return 0
    }
    
    result := nums[0]
    curr := result

    for i := 1; i < count; i++ {
        if nums[i] < curr + nums[i] {
            curr += nums[i]
        } else {
            curr = nums[i]
        }
        if result < curr {
            result = curr
        }
    }
    
    return result
}