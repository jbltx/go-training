package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
    
    strLen1 := len(str1)
    strLen2 := len(str2)
    
    for s := 0 ; s < strLen1; s++ {
        for t := 0; s < strLen1 && t < strLen2; t++ {
            if str1[s] != str2[t] {
                return ""
            }
            if t < strLen2 - 1 {
                s++
            }
        }
    }
    
    for s := strLen2; s >= 1; s-- {
        if strLen1 % s == 0 && strLen2 % s == 0 {
            return str2[:s]
        }
    }
    
    return ""
}