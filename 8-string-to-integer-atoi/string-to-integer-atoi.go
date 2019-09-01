package main

import "math"

func myAtoi(str string) int {
    
    count := len(str)
    result := 0
    
    if count == 0 {
        return result
    }
    
    i := 0
    for ; i < count && str[i] == ' ' ; i++ {
        
    }
    
    if i == count {
        return result
    }
    
    sign := 1
    if str[i] == '-' {
        sign = -1
        i++
    } else if str[i] == '+' {
        i++
    }
    
    max10 := math.MaxInt32 / 10
    min10 := math.MinInt32 / 10
    
    for ; i < count && str[i] >= '0' && str[i] <= '9' ; i++ {
        num := int(str[i] - '0')
        
        if sign == 1 && (result > max10 || (result == max10 && num > 7)) {
            return math.MaxInt32
        } 
        if sign == -1 && (-result < min10 || (-result == min10 && -num < -8)) {
            return math.MinInt32
        } 
        result = result * 10 + num
    }
    
    return result * sign
}