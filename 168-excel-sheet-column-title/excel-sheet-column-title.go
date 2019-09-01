package main

func convertToTitle(n int) string {
    
    result := "" 
    if n == 0 {
        return result
    }
    
    return convertToTitle((n-1) / 26) + string(byte((n-1) % 26 + 'A'))
    
}