package main

func titleToNumber(s string) int {
    
    count := len(s)   
    result := 0
    
    for i := 0; i < count; i++ {
        result = result * 26 + int(s[i] - 'A' + 1)
    }
    
    return result
}