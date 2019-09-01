package main

import "math"

func reverse(x int) int {
	rev := 0
    const intMax10 int = math.MaxInt32 / 10
    const intMin10 int = math.MinInt32 / 10
    for ;x != 0; {
        pop := x % 10
        x /= 10
        if rev < intMin10 || (rev == intMin10 && pop < -8) || rev > intMax10 || (rev == intMax10 && pop > 7) {
            return 0
        }
        rev = rev * 10 + pop
    }
	return rev  
}