package main

func islandPerimeter(grid [][]int) int {
    
    ilen := len(grid)
    
    if ilen == 0 {
        return ilen
    }
    
    jlen := len(grid[0])
    
    if jlen == 0 {
        return jlen
    }
    
    result := 0
    
    for i := 0; i < ilen; i++ {
        for j := 0; j < jlen; j++ {
            if grid[i][j] == 1 {
                
                val := 4
                
                if i > 0 && grid[i-1][j] == 1 {
                    val--
                }
                if i < ilen-1 && grid[i+1][j] == 1 {
                    val--
                }
                if j > 0 && grid[i][j-1] == 1 {
                    val--
                }
                if j < jlen-1 && grid[i][j+1] == 1 {
                    val--
                }
                result += val
                
            }  
        }
    } 
    
    return result
}