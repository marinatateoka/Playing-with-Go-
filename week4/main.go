package main

import (
	"fmt"
)

//Create a function that receives a string “1,4,5,6,3,5, 6” as param, and returns a reverse list of string

func main() {
	l := StringToSlice("marina")
	fmt.Println(l)

	x := Reverse(l)
	fmt.Println(x)
}

func StringToSlice(myStr string) []string {

	arrayString := []string{}
	for _, letra := range myStr {
		arrayString = append(arrayString, string(letra))
	}
	return arrayString
}

func Reverse(arrayInitial []string) []string {
	arrayFinal := []string{}

	for i := len(arrayInitial) - 1; i >= 0; i-- {
		arrayFinal = append(arrayFinal, arrayInitial[i])
	}
	return arrayFinal
}
