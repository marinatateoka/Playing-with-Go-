package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Create a function that receives a string “1,4,5,6,3,5, 6” as param, and returns a reverse list of integers

func main() {
	l, err := ReturnIntList("1, 2, 3, 4, 5, 6")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)

	x := Reverse(l)
	fmt.Println(x)
}

func ReturnIntList(myStr string) ([]int, error) {
	newStr := strings.Replace(myStr, " ", "", -1)
	res1 := strings.Split(newStr, ",")
	res2 := []int{}

	for _, i := range res1 {
		j, err := strconv.Atoi(i)
		if err != nil {
			return nil, (err)
		}
		res2 = append(res2, j)
	}
	return res2, nil
}

func Reverse(arrayInitial []int) []int {
	arrayFinal := []int{}

	for i := len(arrayInitial) - 1; i >= 0; i-- {
		arrayFinal = append(arrayFinal, arrayInitial[i])
	}
	return arrayFinal
}
