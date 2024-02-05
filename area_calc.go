package main

import (
	"fmt"
	"math"
	"strconv"
)

var multiplier float64 = 4.0

func main() {
	println("Type 'exit' or 'CTRL + C' To Quit.")
	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	var input string
	for {
		print("Enter Material Thickness: ")
		var _, err = fmt.Scan(&input)
		if err != nil {
			panic(err)
		}

		var thickness, err2 = strconv.ParseFloat(input, 16)

		if err2 != nil {
			panic(err2)
		}

		area := multiplier * math.Pi * math.Pow(thickness, 2)

		println(fmt.Sprintf("Calculated Output Area: %.4f", area))
		println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
}
