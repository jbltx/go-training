package main

func longestPalindrome(s string) string {
    
    length := len(s)
    result := ""
    
    inc := -1
    for center := length / 2; center < length; center+=inc {
        a := center
        b := center
        even := 0
        
        for ; ; a,b = a-1,b+1 {
                    
            if a < 0 || b >= length || s[a] != s[b] {
                if even == 0 {
                    a = center + 1
                    b = center
                    even++
                } else {
                    break
                }
            } else {
                if b-a >= len(result) {
                    result = s[a:b+1]
                }
            }
        }
               
        if center == 0 {
            center = length / 2
            inc = 1
        }
    }
    
    return result
}