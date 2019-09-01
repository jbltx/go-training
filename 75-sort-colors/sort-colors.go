package main

// Counting sort algorithm
/*func sortColors(nums []int)  {
    
    count := len(nums)
    if count <= 1 {
        return
    }
    
    colors := make([]int, 3)
    
    for i := 0; i < count; i++ {
        colors[nums[i]]++
    }
    
    for i,j := 0,0; i < 3; i++ {
        for ; colors[i] > 0 ; colors[i],j=colors[i]-1,j+1 {
            nums[j] = i
        }
    }
}*/

// Bubble sort algorithm
func bubbleSort(nums []int, n int) {
    
    newn := 0
    for i := 0; i < n; i++ {
        if nums[i] > nums[i + 1] {
            newn = i + 1
            cache := nums[i]
            nums[i] = nums[i + 1]
            nums[i + 1] = cache
        }
    }
    
    if newn > 0 {
        bubbleSort(nums, n)
    }
    
}

func sortColors(nums []int)  {
    
    count := len(nums)
    if count <= 1 {
        return
    }
    
    bubbleSort(nums, count - 1)
}