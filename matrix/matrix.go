package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Cell struct {
	color  string
	row    int
	col    int
	isMark bool
}

func getColor(colors []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	n := rand.Intn(len(colors))
	return colors[n]
}

//Вычесление смежных одноцветных ячеек
func CustomAdjacency(inTrend int, matrix [][]Cell, col int, row int, color string, result map[string][][]int, maxNum *int) {
	switch inTrend {
	case 0:
		row++
		if row > (len(matrix[0])-1) || matrix[col][row].isMark {
			return
		}
		if matrix[col][row].color == color {
			result[color][col][row] = 1
			matrix[col][row].isMark = true
			CustomAdjacency(0, matrix, col, row, color, result, maxNum)
			CustomAdjacency(3, matrix, col, row, color, result, maxNum)
			CustomAdjacency(2, matrix, col, row, color, result, maxNum)
			*maxNum++
		}
	case 1:
		row--
		if row < 0 || matrix[col][row].isMark {
			return
		}
		if matrix[col][row].color == color {
			result[color][col][row] = 1
			matrix[col][row].isMark = true
			CustomAdjacency(1, matrix, col, row, color, result, maxNum)
			CustomAdjacency(2, matrix, col, row, color, result, maxNum)
			CustomAdjacency(3, matrix, col, row, color, result, maxNum)
			*maxNum++
		}
	case 2:
		col--
		if col < 0 || matrix[col][row].isMark {
			return
		}
		if matrix[col][row].color == color {
			result[color][col][row] = 1
			matrix[col][row].isMark = true
			CustomAdjacency(2, matrix, col, row, color, result, maxNum)
			CustomAdjacency(1, matrix, col, row, color, result, maxNum)
			CustomAdjacency(0, matrix, col, row, color, result, maxNum)
			*maxNum++
		}
	case 3:
		col++
		if col > (len(matrix)-1) || matrix[col][row].isMark {
			return
		}
		if matrix[col][row].color == color {
			result[color][col][row] = 1
			matrix[col][row].isMark = true
			CustomAdjacency(3, matrix, col, row, color, result, maxNum)
			CustomAdjacency(0, matrix, col, row, color, result, maxNum)
			CustomAdjacency(1, matrix, col, row, color, result, maxNum)
			*maxNum++
		}
	}
}

func initData(colors []string, matrix [][]Cell, resultMatrix map[string][][]int, tempMatrix map[string][][]int, col int, row int) {
	for _, value := range colors {
		resultMatrix[value] = make([][]int, col)
		tempMatrix[value] = make([][]int, col)
		for i := 0; i < col; i++ {
			resultMatrix[value][i] = make([]int, row)
			tempMatrix[value][i] = make([]int, row)
		}
	}
	for i := 0; i < col; i++ {
		matrix[i] = make([]Cell, row)
		for j := 0; j < row; j++ {
			val := Cell{
				col:    i,
				row:    j,
				color:  getColor(colors),
				isMark: false,
			}
			matrix[i][j] = val
		}
	}
}

//Выборка результатов выполнения функции CustomAdjacency по наибольшему кол-ву необходимых элементов
//Если в выборке будет присутствовать промежуточные матрицы с одинаковым кол-вом необходимых эл-то, то результатом будет первая матрица
func adjacentMarks(matrix [][]Cell, resultMatrix map[string][][]int, tempMatrix map[string][][]int, col int, row int) {
	resultNum := 0
	for i, x := range matrix {
		for y := range x {
			maxNum := 0
			CustomAdjacency(0, matrix, i, y-1, matrix[i][y].color, tempMatrix, &maxNum)
			if resultNum < maxNum {
				resultNum = maxNum
				resultMatrix[matrix[i][y].color] = resultMatrix[matrix[i][y].color][:0][:0]
				resultMatrix[matrix[i][y].color] = make([][]int, col)
				for k := 0; k < col; k++ {
					resultMatrix[matrix[i][y].color][k] = make([]int, row)
				}
				for k := 0; k < col; k++ {
					for m := 0; m < row; m++ {
						resultMatrix[matrix[i][y].color][k][m] = tempMatrix[matrix[i][y].color][k][m]
					}
				}
			}
			tempMatrix[matrix[i][y].color] = tempMatrix[matrix[i][y].color][:0][:0]
			tempMatrix[matrix[i][y].color] = make([][]int, col)
			for k := 0; k < col; k++ {
				tempMatrix[matrix[i][y].color][k] = make([]int, row)
			}
		}
	}
}

func main() {
	//Цвета для маркировки ячеек
	colors := []string{
		"\033[31m", // red
		"\033[33m", // yellow
		"\033[34m", // blue
		"\033[35m", // purple
	}
	//Цвет по-умолчанию
	colorReset := "\033[0m"
	var col, row int
	//Маркировка ячеек исходной таблицы
	symbol := "0"
	//Ввод кол-ва столбцов и строк через пробел
	fmt.Println("Введите количестов столбцов и строк:")
	fmt.Scanf("%d %d", &col, &row)
	//Генерируемая матрица
	matrix := make([][]Cell, col)
	//Результирующая и промежуточные матрицы
	resultMatrix, tempMatrix := make(map[string][][]int), make(map[string][][]int)
	initData(colors, matrix, resultMatrix, tempMatrix, col, row)
	//Печать начальной матрицы
	for _, x := range matrix {
		for y := range x {
			value := x[y].color + symbol + colorReset + " "
			fmt.Print(value)
		}
		fmt.Println()
	}
	adjacentMarks(matrix, resultMatrix, tempMatrix, col, row)
	//Печать результирующей матрицы
	fmt.Printf("\n\n")
	resultNum := 0
	maxColor := ""
	for valueColor, doubleSlice := range resultMatrix {
		tempNum := 0
		for _, slice := range doubleSlice {
			for _, x := range slice {
				if x > 0 {
					tempNum++
				}
			}
		}
		if resultNum < tempNum {
			resultNum = tempNum
			maxColor = valueColor
		}
	}
	for _, x := range resultMatrix[maxColor] {
		for y := range x {
			var valueColor string
			if x[y] > 0 {
				valueColor = maxColor
			} else {
				valueColor = colorReset
			}
			fmt.Printf("%s%s%s ", valueColor, strconv.Itoa(x[y]), colorReset)
		}
		fmt.Println()
	}
}
