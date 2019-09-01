package main

// Abs returns the absolute value of n
func Abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func threeSumClosest(nums []int, target int) int {
    
    count := len(nums)
    
    // The input array should have at least 3 values
    if count < 3 {
        return 0
    }
    
    // We first sort the array
    // This will be useful to skip a lot of 
    // iterations in the loop
    sort.Ints(nums)
    
    // Initialize the result with a valid sum from
    // the input array
    result := nums[0] + nums[1] + nums[2]
    
    // Iterate through n-2 values in the array
    // Since we check for a sum of 3 values, we need
    // to skip the 2 last ones
    for i := 0; i < count-2; i++ {
        
        // Since the array is sorted, if our current
        // value of nums[i] is greater than the target,
        // all next values will be too large so 
        // we can stop here
        if nums[i] >= 0 && nums[i] > target {
            break
        }
        
        // If our current value nums[i] is equal to the value 
        // from the last iteration, all the next tests will give 
        // the same results, so we can skip this iteration
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        // If all values in the array are negatives (or 0)
        // nums[i] should be at least equal to target to
        // make valid tests, if not we can skip it
        if nums[count-1] <= 0 && nums[i] < target {
            continue
        }
        
        // We initialize our 2 pointers for the others
        // 2 values to test against nums[i]
        // At the begining, the left pointer is the pointer located next to i
        // and the right pointer is the pointer of the last value of the array
        left := i+1
        right := count-1
        
        // We will iterate through this range of values between the 
        // left and right pointer
        // At each iteration, depdending of the sum found, the left pointer can
        // be moved to the next value, or the right pointer moved to the previous value
        // in the array (it's an optimization to find quickly the exact or closest match with the target)
        for ;left < right; {
            
            // Calculate the current sum
            sum := nums[i] + nums[left] + nums[right]
            
            // If the exact value of target is found
            // there's no better solution so we can 
            // return it
            if sum == target {
                return sum
            }
            
            // If the current found sum is closest to the target than the main result,
            // set the main result value as the current sum
            if Abs(target - result) > Abs(target - sum) {
                result = sum
            }

            // Depending of the value of the current sum,
            // move the left or right pointer
            if sum > target {
                // While the right pointer will point to the same value,
                // decrement it (optimization to skip tests)
                for ;right < left && nums[right] == nums[right-1]; {
                    right--
                }
                right--
            } else {
                // While the left pointer will point to the same value,
                // increment it (optimization to skip tests)
                for ;left < right && nums[left] == nums[left+1]; {
                    left++
                }
                left++
            }
        }
    }
    return result
}