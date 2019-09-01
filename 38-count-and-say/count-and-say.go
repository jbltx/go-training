package main

import "strings"

func countAndSay(n int) string {
    
    if n == 1 {
        return "1"
    }
    
    cas := countAndSay(n - 1)
    length := len(cas)
    
    c := cas[0]
    j := byte(0x31)
    var result strings.Builder
    result.Grow(length)
    
    for i := 1; i < length; i++ {
        if cas[i] == c {
            j++
        } else {
            // write in result
            result.WriteByte(j)
            result.WriteByte(c)
            c = cas[i]
            j = byte(0x31)
        }
    }
    
    result.WriteByte(j)
    result.WriteByte(c)
    
    return result.String()
}