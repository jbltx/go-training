package main

import "strings"

func convert(s string, numRows int) string {
    
    if len(s) == 0 {
        return ""
    }
    
    if numRows == 1 {
        return s
    }
    
    count := len(s)
    var result strings.Builder
    result.Grow(count)
    
    /*
     * The phase length defines the number of character
     * to make a zigzag phase (V shape)
     */
    phaseLength := (numRows - 1) + numRows
    
    
    // Iterate for each row...
    for n := 0; n < numRows; n++ {
        
        /*
         * For first and last row, each next character
         * is the limit of the phase, so we can increment
         * by phaseLength-1 to get next character
         */
        if n == 0 || n == numRows - 1 {
            for i := n ; i < count; i+=phaseLength-1 {
                result.WriteByte(s[i]) 
            }
        } else {
            /*
             * For others rows, the phase will vary
             * between even and odd iterations
             * The first iteration will have a phase equals
             * to phaseLength - (n * 2), and the second one
             * equals to the difference between the original
             * phase length and the value found for the first
             * iteration. Increment value are the phase value
             * minus one because we start at the first element 
             * of the phase so we don't take it in account
             */
            increment1 := phaseLength - (2 * n) - 1
            increment2 := phaseLength - increment1 - 1
            for i,j := n,0 ; i < count; j++ {                
                result.WriteByte(s[i]) 
                
                if j % 2 == 0 {
                    i += increment1
                } else {
                    i += increment2
                }
            }
        }        
    }
    
    return result.String()
}