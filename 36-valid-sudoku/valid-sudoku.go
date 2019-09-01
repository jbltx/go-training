package main

func isValid(board [][]byte, x int, y int, n byte) bool {
    sqrX,sqrY := (x/3)*3,(y/3)*3
    for i:= 0; i < 9; i++ {
        if board[y][i] == n || board[i][x] == n || board[sqrY+i/3][sqrX+i%3] == n {
            return false
        }
    }
    return true
}

func isValidSudoku(board [][]byte) bool {
    
    if board == nil || len(board) != 9 {
        return false
    }
    
    for y := 0; y < 9; y++ {
        for x := 0; x < 9; x++ {
            if board[y][x] != '.' {
                n := board[y][x]
                board[y][x] = '.'
                if !isValid(board, x, y, n) {
                    return false
                }
                board[y][x] = n
            }
        }
    }   
    return true
}