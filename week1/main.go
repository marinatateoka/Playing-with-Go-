package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l, err := ReturnIntList("1, 2, 3,4, 5, 6")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)

	l, err = ReturnIntList("1, 2, 3,4, 5, 6")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
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
