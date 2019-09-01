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

func solve(board [][]byte, startX int, startY int) bool {    
    for y := startY; y < 9; y,startX=y+1,0 {
        for x := startX; x < 9; x++ {
            if board[y][x] == '.' {
                for n := byte('1'); n <= byte('9'); n++ {
                    if isValid(board, x, y, n) {
                        board[y][x] = n
                        if solve(board, x + 1, y) {
                            return true
                        } else {
                            board[y][x] = '.'
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}

func solveSudoku(board [][]byte) {
    if board == nil || len(board) == 0 {
        return
    }
    solve(board, 0, 0)
}