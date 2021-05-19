package main

import "testing"

func BenchmarkInitDataSmallSize(b *testing.B) {
	colors := []string{
		"\033[31m", // red
		"\033[33m", // yellow
		"\033[34m", // blue
		"\033[35m", // purple
	}
	var col, row int
	col, row = 20, 20
	matrix := make([][]Cell, col)
	resultMatrix, tempMatrix := make(map[string][][]int), make(map[string][][]int)
	for i := 0; i < b.N; i++ {
		initData(colors, matrix, resultMatrix, tempMatrix, col, row)
	}
}

func BenchmarkInitDataBigSize(b *testing.B) {
	colors := []string{
		"\033[31m", // red
		"\033[33m", // yellow
		"\033[34m", // blue
		"\033[35m", // purple
	}
	var col, row int
	col, row = 100, 100
	matrix := make([][]Cell, col)
	resultMatrix, tempMatrix := make(map[string][][]int), make(map[string][][]int)
	for i := 0; i < b.N; i++ {
		initData(colors, matrix, resultMatrix, tempMatrix, col, row)
	}
}

func BenchmarkAdjacentMarksSmallSize(b *testing.B) {
	colors := []string{
		"\033[31m", // red
		"\033[33m", // yellow
		"\033[34m", // blue
		"\033[35m", // purple
	}
	var col, row int
	col, row = 20, 20
	matrix := make([][]Cell, col)
	resultMatrix, tempMatrix := make(map[string][][]int), make(map[string][][]int)
	initData(colors, matrix, resultMatrix, tempMatrix, col, row)
	for i := 0; i < b.N; i++ {
		adjacentMarks(matrix, resultMatrix, tempMatrix, col, row)
	}
}

func BenchmarkAdjacentMarksBigSize(b *testing.B) {
	colors := []string{
		"\033[31m", // red
		"\033[33m", // yellow
		"\033[34m", // blue
		"\033[35m", // purple
	}
	var col, row int
	col, row = 100, 100
	matrix := make([][]Cell, col)
	resultMatrix, tempMatrix := make(map[string][][]int), make(map[string][][]int)
	initData(colors, matrix, resultMatrix, tempMatrix, col, row)
	for i := 0; i < b.N; i++ {
		adjacentMarks(matrix, resultMatrix, tempMatrix, col, row)
	}
}
