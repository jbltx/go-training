package main

func mySqrt(x int) int {
    
    if x == 0 {
        return 0
    }

    a := float64(x)
    result := float64(1)
    
    for i := 0; i < x; i++ {
        newResult := 0.5 * (result + a / result)
        
        if int(newResult) == int(result) {
            break
        }
        result = newResult
    }
    
    return int(result)
}