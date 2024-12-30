package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string 4.5 to float64
	parseFloat, err := strconv.ParseFloat("4.5", 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}

	fmt.Println(parseFloat)
}