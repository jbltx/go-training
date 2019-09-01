package main

func findTheDifference(s string, t string) byte {
    
    count := len(t)
    result := t[count - 1]
    
    for i := 0; i < count - 1; i++ {
        result ^= s[i]
        result ^= t[i]
    }
    
    return result
}