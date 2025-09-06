package main

import "fmt"

// func checkXMatrix(grid [][]int) bool {
//     r := len(grid)

//     for i, _ := range grid {
//         for j, _ := range grid[i] {
//             if i != j && (i+j+1) != r {
//                 if grid[i][j] != 0 {
//                     return false
//                 }
//             } else {
//                 if grid[i][j] == 0 {
//                     return false
//                 }
//             }
//         }
//     }

//     return true
// }

func checkXMatrix(grid [][]int) bool {
	for i, row := range grid {
		for j, v := range row {
			if v == 0 == (i == j || i+j == len(grid)-1) {
				return false
			}
		}
	}

	return true
}

func main() {
	grid := [][]int{
		{2, 0, 0, 1},
		{0, 3, 1, 0},
		{0, 5, 2, 0},
		{4, 0, 0, 2},
	}
	fmt.Println(checkXMatrix(grid))
}
