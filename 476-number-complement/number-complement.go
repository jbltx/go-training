package main

import "math/bits"

func findComplement(num int) int {
    
    if num == 0 {
        return 1
    }
    
    val := uint32(num)
    zeros := uint(bits.LeadingZeros32(val))
    return int(((^val)<<zeros)>>zeros)
}