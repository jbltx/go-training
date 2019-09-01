package main

func duplicateZeros(arr []int)  {
    
    count := len(arr)
    offset := count - 1
    i := 0
    
    // Count the offset generated by the insertion
    // of 0 values
    for ; i < count; i++ {
        if arr[i] == 0 {
            offset++
        }
    }
    
    // Read the array in reverse to apply values at the right place
    // Note : we also check for each iteration if the offset is equal
    // to the current index i to stop the loop, since all previous values
    // will be at the same place.
    for i = count - 1; i >= 0 && i != offset; i,offset=i-1,offset-1 {
        
        // When a 0 is reached, decrement more the offset for next values
        if arr[i] == 0 {
            // swap the value if the offset is valid in the array
            if offset < count {
                arr[offset] = arr[i]
            }
            offset--
        }
        
        // swap the value if the offset is valid in the array
        // Since we have already swapped values if we reached a 0,
        // and decremented the offset, this new swap will add a new 0
        // at the left of the existing one.
        if offset < count {
            arr[offset] = arr[i]
        }
                
    } 
}