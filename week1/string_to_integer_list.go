package main

import (
	"fmt"
	"strings"
)

func main() {
	res1 := strings.SplitN("1, 2, 3, 4, 5, 6", ",", 6)

	fmt.Printf("\nSlice 1: %s", res1)
}
