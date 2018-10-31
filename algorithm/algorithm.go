package algorithm

import (
    "fmt"
)

const (
    stars  = "****************************************************************"
)

func copyMatrix(baseMatrix [][]int) [][]int {
    matrix := make([][]int, len(baseMatrix))
    for i := range matrix {
        matrix[i] = make([]int, len(baseMatrix[i]))
    }

    for i, row := range baseMatrix {
        for j, val := range row {
            matrix[i][j] = val
        }
    }

    return matrix
}

func cityNames(matrix [][]int) ([]rune, []rune) {
    namesByRow := make([]rune, len(matrix))
    namesByColumn := make([]rune, len(matrix[0]))

    for i := range namesByRow {
        namesByRow[i] = rune('A') + rune(i)
    }
    for i := range namesByColumn {
        namesByColumn[i] = rune('A') + rune(i)
    }

    return namesByRow, namesByColumn
}

func printMatrix(matrix [][]int, solutions []coordinate) {
    fmt.Println(stars)
    fmt.Println()
    for i, row := range matrix {
        if !canProceed(i, solutions, true) {
            continue
        }
        for j, val := range row {
            if !canProceed(j, solutions, false) {
                continue
            }
            fmt.Printf("%v\t", val)
        }
        fmt.Println()
    }
    fmt.Println()
    fmt.Println(stars)
}

func printCalculations(calculations []string) {
    for _, val := range calculations {
        fmt.Println(val)
    }
}

func printPaths(namesByRow, namesByColumn []rune, solutions []coordinate) {
    for _, val := range solutions {
        fmt.Println(string(namesByRow[val.row]) + " -> " + 
            string(namesByColumn[val.column]))
    }
}

func Calculate() {
//    baseMatrix := [][]int{{-1, 7, 6, 8, 4},
  //                        {7, -1, 8, 5, 6},
    //                      {6, 8, -1, 9, 7},
      //                    {8, 5, 9, -1, 8},
        //                  {4, 6, 7, 8, -1}}

    baseMatrix := [][]int{{-1, 5, 4, 9, 17},
                          {10, -1, 13, 19, 8},
                          {20, 18, -1, 15, 23},
                          {3, 16, 12, -1, 17},
                          {11, 36, 13, 7, -1}}

    matrix := copyMatrix(baseMatrix)
    namesByRow, namesByColumn := cityNames(matrix)
    //_ = namesByRow
    //_ = namesByColumn

    fmt.Println("Base matrix")
    solutions := make([]coordinate, 0)
    printMatrix(matrix, solutions)
    for len(solutions) != len(baseMatrix) {
        fmt.Println("Row minimization")
        rowMinimization(matrix, solutions)
        printMatrix(matrix, solutions)
        fmt.Println("Column minimization")
        columnMinimization(matrix, solutions)
        printMatrix(matrix, solutions)
        fmt.Println("Calculate Panelties")
        kindlyPanelties := calculatePanelties(matrix, solutions)
        fmt.Println(kindlyPanelties)
        coord := rowAndColumnToDestroy(matrix, kindlyPanelties)
        solutions = append(solutions, coord)

        matrix[coord.column][coord.row] = -1
        fmt.Println(solutions)
        printMatrix(matrix, solutions)
    }

    printPaths(namesByRow, namesByColumn, solutions)
}

func canProceed(index int, coordinates []coordinate, row bool) bool {
    for _, coord := range coordinates {
        val := 0
        if row {
            val = coord.row
        } else {
            val = coord.column
        }
        if index == val {
            return false
        }
    }
    return true
}

func rowMinimization(matrix [][]int, coordinates []coordinate) {
    minimals := make([]int, 0)
    for i, row := range matrix {
        if !canProceed(i, coordinates, true) {
            continue
        }
        min := 1000000000

        for j, val := range row {
            if !canProceed(j, coordinates, false) {
                continue
            }
            if val != -1 && min > val {
                min = val
            }
        }
        minimals = append(minimals, min)
    }

    index := 0
    for i, row := range matrix {
        if !canProceed(i, coordinates, true) {
            index--
            continue
        }
        for j := range row {
            if !canProceed(j, coordinates, false) {
                continue
            }
            if row[j] != -1 {
                row[j] -= minimals[i + index]
            }
        }
    }
}

func columnMinimization(matrix [][]int, coordinates []coordinate) {
    minimals := make([]int, 0)
    for i := 0; i < len(matrix[0]); i++ {
        if !canProceed(i, coordinates, false) {
            continue
        }
        min := 1000000000
        for j := 0; j < len(matrix); j++ {
            if !canProceed(j, coordinates, true) {
                continue
            }
            if matrix[j][i] != -1 && matrix[j][i] < min {
                min = matrix[j][i]
            }
        }

        minimals = append(minimals, min)
    }

    index := 0
    for i := 0; i < len(matrix[0]); i++ {
        if !canProceed(i, coordinates, false) {
            index--
            continue
        }
        for j := 0; j < len(matrix); j++ {
            if !canProceed(j, coordinates, true) {
                continue
            }
            if matrix[j][i] != -1 {
                matrix[j][i] -= minimals[i + index]
            }
        }
    }
}

type coordinate struct {
    row, column int
}

func calculatePanelties(matrix [][]int, solutions []coordinate) map[coordinate]int {
    panelties := make(map[coordinate]int, 0)

    coordinates := make([]coordinate, 0)
    for i, row := range matrix {
        if !canProceed(i, solutions, true) {
            continue
        }
        for j, val := range row {
            if !canProceed(j, solutions, false) {
                continue
            }
            if val == 0 {
                coordinates = append(coordinates, coordinate{i, j})
            }
        }
    }

    for _, val := range coordinates {
        minByRow := 1000000000
        for j := 0; j < len(matrix[val.row]); j++ {
            if !canProceed(j, solutions, false) {
                continue
            }
            if matrix[val.row][j] != -1 && j != val.column {
                if minByRow > matrix[val.row][j] {
                    minByRow = matrix[val.row][j]
                }
            }
        }
        minByColumn := 1000000000
        for i := 0; i < len(matrix); i++ {
            if !canProceed(i, solutions, true) {
                continue
            }
            if matrix[i][val.column] != -1 && i != val.row {
                if minByColumn > matrix[i][val.column] {
                    minByColumn = matrix[i][val.column]
                }
            }
        }

        panelties[coordinate{val.row, val.column}] = minByRow + minByColumn
    }
    return panelties
}

func rowAndColumnToDestroy(matrix [][]int,
    panelties map[coordinate]int) coordinate {//, solutions []coordinate) coordinate {
    max := 0
    key := coordinate{0, 0}
    for k, val := range panelties {
        if max < val {
            max = val
            key = k
        }
    }

    return key
}
