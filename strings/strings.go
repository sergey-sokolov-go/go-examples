package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func delElement(slice []int) []int {
	fmt.Printf("PRE_FUNC len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Println("address of 0th element:", &slice[0])
	fmt.Printf("PRE_FUNC slice_1=%v slice_2=%v\n", slice[:1], slice[2:])
	slice = append(slice[:1], slice[2:]...)
	fmt.Printf("POST_FUNC len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Println("address of 0th element:", &slice[0])
	//slice = append(slice, slice[2:]...)
	return slice
}

//TODO пересмотреть лекцию и обратить внимание на работу состроками!
func main() {
	slice := []int{8, 1, 2, 3}
	fmt.Printf("PRE len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Println("address of 0th element:", &slice[0])
	slice = delElement(slice)
	fmt.Printf("POST len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Println("address of 0th element:", &slice[0])
	str := `ab2c\2dr3`
	n := 0
	//var resultString []rune
	var tempRune []rune
	for len(str) > 0 {
		if n == 0 && str[n] > 47 && str[n] < 58 {
			log.Fatal("data incorrect")
		}
		n++
		runeValue, size := utf8.DecodeRuneInString(str)
		tempRune = append(tempRune, runeValue)
		//fmt.Printf("str = %s\n", str)
		//fmt.Printf("rune = %c \n", runeValue)
		if n%2 == 0 {
			if tempRune[0] > 47 && tempRune[0] < 58 && tempRune[1] > 47 && tempRune[1] < 58 {
				log.Fatal("data incorrect")
			}
			if !(tempRune[0] > 47 && tempRune[0] < 58) {
				if tempRune[1] > 47 && tempRune[1] < 58 {
					//i := int(tempRune[1])
					//fmt.Printf("str_int = %d\n", i)
					/*for i := 0; i < int(tempRune[1]); i++{

					}*/
				}
			}
			tempRune = tempRune[2:]
		}
		str = str[size:]
	}
}
