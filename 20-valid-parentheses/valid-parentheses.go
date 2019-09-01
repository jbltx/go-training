package main

func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    
    str := []rune(s)
    var scopeStack []rune
    
    for i := 0; i < len(str); i++ {
        if str[i] == '(' {
            scopeStack = append(scopeStack, ')')
        } else if str[i] == '{' {
            scopeStack = append(scopeStack, '}')
        } else if str[i] == '[' {
            scopeStack = append(scopeStack, ']')
        } else if len(scopeStack) <= 0 {
            return false
        } else if scopeStack[len(scopeStack) - 1] != str[i] {
            return false
        } else {
            scopeStack = scopeStack[:len(scopeStack) - 1]
        }
    }
    
    if len(scopeStack) > 0 {
        return false
    }
    
    return true
}