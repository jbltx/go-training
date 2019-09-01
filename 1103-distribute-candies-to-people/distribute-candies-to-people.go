package main

func distributeCandies(candies int, num_people int) []int {
    result := make([]int, num_people)
    
    for dist := 1 ; candies > 0; {
        for i := 0; candies > 0 && i < num_people; i,dist,candies=i+1,dist+1,candies-dist {
            if candies - dist < 0 {
                result[i] += candies
            } else {
                result[i] += dist
            }
        }
    }
    
    return result
}