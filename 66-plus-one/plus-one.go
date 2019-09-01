package main

func plusOne(digits []int) []int {
    
    count := len(digits)
    for i := count -1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i] += 1
            return digits
        } else {
            digits[i] = 0
        }
    }
    
    digits = make([]int, count+1)
    digits[0] = 1
    
    return digits
}