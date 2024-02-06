package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

var ratio = 4.0

var noError = false

func main() {
	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	println("Type 'exit' Or 'CTRL + C' To Quit.")

	for noError == false {
		print("Enter Ratio Value: ")
		var _, err = fmt.Scan(&ratio)

		if err != nil {
			println("Ratio Must Be A Number.")
			ratio = 4.0
		} else {
			noError = true
		}
	}

	println(fmt.Sprintf("Ratio Set: %.0f:1", ratio))
	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	var input string
	for {
		print("Enter Material Thickness: ")
		var _, err = fmt.Scan(&input)
		if err != nil {
			println("Application Panicked.")
			input = "0"
		}

		if input == "exit" {
			println("Exiting...")
			time.Sleep(1 * time.Second)
			break
		}

		var thickness, err2 = strconv.ParseFloat(input, 16)

		if err2 != nil {
			println("Thickness Must Be A Number.")
			thickness = 0
		}

		area := ratio * math.Pi * math.Pow(thickness, 2)

		println(fmt.Sprintf("Calculated Output Area: %.4f", area))
		println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
}
