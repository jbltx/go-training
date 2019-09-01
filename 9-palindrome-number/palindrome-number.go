package main

import "strconv"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    s := []rune(strconv.Itoa(x))
    slen := len(s)
    for i,j := 0, slen - 1; i < slen; i,j = i+1,j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}