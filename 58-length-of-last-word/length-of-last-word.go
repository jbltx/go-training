package main

func lengthOfLastWord(s string) int {
    result := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' ' && result > 0 { 
            break
        } else if s[i] != ' ' {
            result++
        }
    }
    return result
}