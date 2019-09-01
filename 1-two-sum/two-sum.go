package main

func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    result := make([]int, 2)
    for i := 0; i < len(nums); i++ {
        if j, ok := hashMap[target - nums[i]]; ok {
            result[0] = j
            result[1] = i
            break
        }
        hashMap[nums[i]] = i
    }
    return result
}