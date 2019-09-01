package main



func Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func Min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxArea(height []int) int {
    
    max := 0
    count := len(height)
    
    for i,j := 0,count-1; i < j; {
        
        max = Max(max, (j-i) * Min(height[i],height[j]))
        
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
        
    }
    
    return max
}