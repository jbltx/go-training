package main

import "sort"

func threeSum(nums []int) [][]int {
    
    count := len(nums)
    result := [][]int{}
    
    sort.Ints(nums)
    
    for i := 0; i < count-2; i++ {
        
        if nums[i] > 0 {
            break
        }
        
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
            
        for left,right := i+1,count-1; left < right; {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{ nums[i], nums[left], nums[right]})
                    
                for ;left < right && nums[left] == nums[left+1]; {
                    left++
                }
                    
                for ;right > left && nums[right] == nums[right-1] ; {
                    right--
                }
                    
                right--
                left++
            } else if sum > 0 {
                right--
            } else {
                left++
            }
        }
        
    }
    
    return result
    
}