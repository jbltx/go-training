package main

import "strconv"

func fizzBuzz(n int) []string {
    
    if n <= 0 {
        return []string{}
    }
    
    result := make([]string, n)
    fizz := "Fizz"
    buzz := "Buzz"
    fizzbuzz := "FizzBuzz"
    
    for i := 1; i <= n; i++ {
        
        if i % 15 == 0 {
            result[i-1] = fizzbuzz
        } else if i % 3 == 0 {
            result[i-1] = fizz
        } else if i % 5 == 0 {
            result[i-1] = buzz
        } else {
            result[i-1] = strconv.Itoa(i)
        }
    }
    
    return result
}