package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	// fmt.Println(colors)

	// for i := 0; i < len(colors); i++ {
	// 	fmt.Println(colors[i])
	// }

	// for i := range colors {
	// 	fmt.Println(colors[i])
	// }

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	sum := 1
	for {
		sum += sum
		fmt.Println("Sum:", sum)

		if sum > 10000 {
			break
		}

	}
}
