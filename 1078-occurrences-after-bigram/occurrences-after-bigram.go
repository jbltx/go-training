package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
    
    length := len(text)
    result := make([]string, 0)
    query := first + " " + second
    queryLength := len(query)
    
    for s := 0 ; s < length; s++ {
        for j := 0 ; s < length && j < queryLength ; s,j=s+1,j+1 {
            
            if text[s] != query[j] {
                s -= j
                break
            }
            
            if j == queryLength - 1 {
                marker := s + 1
                s += 2
                if s < length {
                    var b strings.Builder
                    for ; s < length && text[s] != ' ' ; s++ {
                        b.WriteByte(text[s])
                    }
                    if b.Len() > 0 {
                        result = append(result, b.String())
                    }
                    s = marker
                }
                break
            }
        }
    }
    
    return result
}