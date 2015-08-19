package main 

import (
        "math"
        "./sudoku"
)

func getSize(bstr string) int {
        return int(math.Sqrt(float64(len(bstr))))
}

func main() {
        text := "1...79.8." + ".5.413.6." + ".6......." +
                        "..1......" + "8.5......" + "....3...7" +
                        ".94.5...." + "7..8.63.9" + ".....74.."
//      text := "......54." + ".278...31" + "........2" +
//                      ".7862...." + "........6" + ".....37.." +
//                      ".432.7..8" + "1..45...3" + "..5.8...."

        // Create the sudoku board
        sz := getSize(text)
        board := sudoku.NewBoard(sz)

        for i,ch := range text {
                r := i/sz + 1
                c := i%sz + 1
                if ch >= '0' && ch <= '9' {
                        board.Set(r,c, ch - '0')
                }
        }
        board.Print()
        if !board.Solved() {
                board.PrintAll()
        }
}
