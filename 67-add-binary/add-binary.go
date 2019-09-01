package main

import "strings"

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func addBinary(a string, b string) string {
   
    result := make([]rune, max(len(a), len(b)) + 1) 
    
    carry := byte('0') // 0x30
    for i,j,k := len(a)-1,len(b)-1,len(result)-1; i >=0 || j >=0 ; i,j,k = i-1,j-1,k-1 {
        
        x := byte('0')
        y := byte('0')
        
        if i >= 0 {
            x = a[i]
        }
        
        if j >= 0 {
            y = b[j]
        }
        
        bytes := x + y + carry
        // 0x30 + 0x30 + 0x30 = 0x90 => carry = 0, val = 0
        // 0x30 + 0x31 + 0x30 = 0x91 => carry = 0, val = 1
        // 0x31 + 0x31 + 0x30 = 0x92 => carry = 1, val = 0
        // 0x31 + 0x31 + 0x31 = 0x93 => carry = 1, val = 1
        
        carry = byte('0')
        if bytes > 0x91 {
            carry = byte('1')
        }
        
        val := '0'
        if bytes == 0x91 || bytes == 0x93 {
            val = '1'
        }
        
        result[k] = val
    }
    
    if carry == 0x31 {
        result[0] = '1'
    } else {
        result = result[1:]
    }
    
    return string(result)
}