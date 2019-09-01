package main

func strStr(haystack string, needle string) int {
    
    hLen := len(haystack)
    nLen := len(needle)
    
    if len(needle) == 0 {
        return 0
    }
    
    j := 0
    
    for i := 0 ; i < hLen; i++ {
        
        if haystack[i] == needle[j] {
            if j == nLen - 1 {
                return i - j
            } else {
                j++
            }
        } else {
            i -= j
            j = 0
        }
        
    }
    
    return -1
}