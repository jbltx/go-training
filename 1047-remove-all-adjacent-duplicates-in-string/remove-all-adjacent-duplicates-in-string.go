package main

func removeDuplicates(S string) string {
    
    count := len(S)
    
    for i := 0; i < count - 1; i++ {
        if S[i] == S[i+1] {
            return removeDuplicates(S[:i]+S[i+2:])
        }
    }
    
    return S
}