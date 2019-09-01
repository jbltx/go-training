package main

/*func factorial(x int) int {
	if x == 0 {
		return 1
	}
	
	return x * factorial(x-1)
}

func combi(p int, n int) int {
    return factorial(n) / (factorial(p) * factorial(n - p))
}*/

func climbStairs(n int) int {
    
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    
    first := 1
    second := 2
    for i := 3; i <= n; i++ {
        third := first + second
        first = second
        second = third
    }
    
    return second
    
    /*
    00
     1
    
    000
     10
     01

    0000
     100
     010
     001
      11
    
    00000
     1000
     0100
     0010
     0001
      101
      011
      110
    
    
    
    result := 1
    target := n / 2
    
    facts := make(map[int]int)
    
    for i := 0; i < target; i++ {
        var factN, factP, factNP int
        
        if fn,ok := facts[n-i-1]; ok {
            factN = fn
        } else {
            factN = factorial(n-i-1)
            facts[n-i-1] = factN
        }
        
        if fp,ok := facts[i+1]; ok {
            factP = fp
        } else {
            factP = factorial(i+1)
            facts[i+1] = factP
        }
        
        if fnp,ok := facts[n-i-1-i-1]; ok {
            factNP = fnp
        } else {
            factNP = factorial(n-i-1-i-1)
            facts[n-i-1-i-1] = factNP
        }
        
        result += factN / (factP * factNP) // combi(i+1, n-i-1)
    }
    
    return result*/
}