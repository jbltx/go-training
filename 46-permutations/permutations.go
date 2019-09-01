package main

func permute(nums []int) [][]int {
    
    result := [][]int{}
    
    if nums == nil {
        return result
    }
    
    n := len(nums)
    
    if n == 0 {
        return result
    } else if n == 1 {
        result = append(result, nums)
        return result
    }
    
    for i := 0; i < n; i++ {
        tp := []int{}
        tp = append(tp, nums[:i]...)
        tp = append(tp, nums[i+1:]...)
        tmp := permute(tp)
        for j := 0; j < len(tmp); j++ {
            p := []int{}
            p = append(p, nums[i])
            p = append(p, tmp[j]...)
            result = append(result, p)
        }
    }

    return result
}
