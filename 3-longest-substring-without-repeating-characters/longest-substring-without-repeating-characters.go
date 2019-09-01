package main

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    hashMap := make(map[rune]bool)
    buffer := []rune{}
    result := 1
    for _, c := range s {
        if val, _ := hashMap[c]; val {
            index := 0
            for ; index < len(buffer); index++ {
                hashMap[buffer[index]] = false
                if buffer[index] == c {
                    break
                }
            }
            buffer = append(buffer, c)
            buffer = buffer[index + 1:]
        } else {
            buffer = append(buffer, c)
            if len(buffer) > result {
                result = len(buffer)
            }
        }
        hashMap[c] = true
    }
    return result
}