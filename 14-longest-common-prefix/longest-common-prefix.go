package main

import "math"
import "strings"

func longestCommonPrefix(strs []string) string {
      
    strlen := math.MaxInt32
    strslen := len(strs)
    
    if strslen == 0 {
        return ""
    }
    
    for _, s := range strs {
        if l := len(s); l < strlen {
            strlen = l
        }
    }
    
    var result strings.Builder
    result.Grow(strlen)
    
    for i := 0; i < strlen; i++ {
        for s := 1; s < strslen; s++ {
            if strs[0][i] != strs[s][i] {
                return result.String()
            }
        }
        result.WriteByte(strs[0][i])
    }
    
    return result.String()
}