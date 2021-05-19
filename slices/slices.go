package main

import (
	"fmt"
	"sort"
	"strings"
)

//Подсчет повторений и дальнейшая смена метсами ключа и значения
func converKeyValue(str []string) map[int]string {
	tempMap := make(map[string]int, len(str))
	for _, value := range str {
		tempMap[value]++
	}
	result := make(map[int]string, len(str))
	for key, value := range tempMap {
		result[value] = key
	}
	return result
}

func a() (j int) {
	i := 0
	defer func() {
		j++
		fmt.Printf("inside defer i = %d\n", j)
	}()
	i++
	fmt.Printf("past i = %d\n", i)
	return i
}
func main() {
	str := strings.Split("hello ward hello path tree hello path", " ")
	result := converKeyValue(str)
	//сортирвка числа повторений
	keys := make([]int, 0, len(result))
	for key, _ := range result {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	//берётся 2 самых часто повторяющихся слова
	keys = append(keys[len(keys)-2:len(keys)-1], keys[len(keys)-1:]...)
	for _, value := range keys {
		fmt.Printf("word = %s count = %d\n", result[value], value)
	}
	fmt.Printf("main fmt = %d\n", a())
}
