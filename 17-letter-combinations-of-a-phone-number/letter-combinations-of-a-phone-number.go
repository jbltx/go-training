package main

func letterCombinations(digits string) []string {
    
    result := []string{}
    
    if len(digits) == 0 {
        return result
    }
    
    phone := map[byte][]string{}
    phone[byte('2')] = []string{"a","b","c"}    
    phone[byte('3')] = []string{"d","e","f"}
    phone[byte('4')] = []string{"g","h","i"}
    phone[byte('5')] = []string{"j","k","l"}
    phone[byte('6')] = []string{"m","n","o"}
    phone[byte('7')] = []string{"p","q","r","s"}
    phone[byte('8')] = []string{"t","u","v"}
    phone[byte('9')] = []string{"w","x","y","z"}
    
    // init
    result = append(result, phone[digits[0]]...)
    digits = digits[1:]
    count := len(digits)
    
    if count == 0 {
        return result
    }

    for c := 0; c < count; c++ {
        
        if digits[c] == '1' {
            return []string{}
        }
        
        currentTotal := len(result)
        total := currentTotal * len(phone[digits[c]])
        delta := total - currentTotal
        
        for k := 0; k < delta; {
            for i := 0; i < currentTotal; i,k=i+1,k+1 {
                result = append(result, result[i])
            }
        }
        
        for i := 0; i < total; i++ {
            result[i] += phone[digits[c]][i/currentTotal]
        }
        
    }
    
    return result
}