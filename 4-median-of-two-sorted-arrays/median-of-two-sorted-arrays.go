package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
        
    merged := make([]int, len(nums1) + len(nums2))
    
    if len(merged) == 0 {
        return 0
    }
    
    if len(nums1) == 0 {
        merged = nums2
    } else if len(nums2) == 0 {
        merged = nums1
    } else {
        i,j,k := 0,0,0

        for ; i < len(nums1) && j < len(nums2) ; k++ {
            if nums1[i] < nums2[j] {
                merged[k] = nums1[i]
                i++
            } else {
                merged[k] = nums2[j]
                j++
            }
        }
        for ; i < len(nums1); i,k = i+1,k+1 {
            merged[k] = nums1[i]
        }
        for ; j < len(nums2); j,k = j+1,k+1 {
            merged[k] = nums2[j]
        }
    }

    result := float64(merged[len(merged) / 2])
    if len(merged) % 2 == 0 {
        result = (result + float64(merged[(len(merged) / 2) - 1])) / 2.0
    }
    
    return result
}