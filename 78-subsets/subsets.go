package main

func subsets(nums []int) [][]int {
    
    result := [][]int{ []int{} }
    count := len(nums)
    
    for i := 0; i < count; i++ {
        size := len(result)
        for j := 0; j < size; j++ {
            subset := make([]int, len(result[j]))
            copy(subset, result[j])
            subset = append(subset, nums[i])
            result = append(result, subset)
        }
    }
    
    return result
}