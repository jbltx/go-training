package main

func romanToInt(s string) int {
    
    arr := []rune(s)
    result := 0
    rlen := len(arr)
    
    for i := 0; i < rlen; i++ {
        switch arr[i] {
        case 'I':
            if i < rlen - 1 && (arr[i + 1] == 'V' || arr[i + 1] == 'X') {
                result--
            } else {
                result++
            }
        case 'V':
            result += 5
        case 'X':
            if i < rlen - 1 && (arr[i + 1] == 'L' || arr[i + 1] == 'C') {
                result -= 10
            } else {
                result += 10
            }
        case 'L':
            result += 50
        case 'C':
            if i < rlen - 1 && (arr[i + 1] == 'D' || arr[i + 1] == 'M') {
                result -= 100
            } else {
                result += 100
            }
        case 'D':
            result += 500
        case 'M':
            result += 1000
        }
    } 
    
    return result
}