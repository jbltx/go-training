package main

func rotate(matrix [][]int)  {
    
    n := len(matrix)
    
    /*
     * Edge cases
     */
    if n == 0 || n == 1 {
        return
    }
    
    /*
     * By analyzing the input and wanted output
     * We can clearly see a rotate by 90 clockwise can
     * be interpreted as : For each row r in the given
     * 2D array, r values will become the values of the column
     * c where c = n - r - 1.
     * When we apply the result, we should first cache old data
     * from this column to be used in other row iterations.
     * For fast lookup i will use a hashmap with key equals 
     * to y * n + x (easy way to have unique ids for each value of 2D array)
     */
    
    cache := make(map[int]int)
    
    // y is the main index for each row processed
    for y := 0; y < n; y++ {
        
        // x is the target column
        x := n - y - 1
        
        // i is the index for each value processed for the row
        for i := 0; i < n; i++ {
            
            // put in cache old data first
            cache[i * n + x] = matrix[i][x]
            
            // now retrieve the valid value to put in
            // the target column
            val := matrix[y][i]
            if v,ok := cache[y * n + i]; ok {
                val = v
            }
            
            matrix[i][x] = val
        }
    }
}