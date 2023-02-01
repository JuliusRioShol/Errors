// Demonstrate errors
package main

import (
	"fmt"
)

func area(length float64, width float64) float64 {
	result := length * width

	return result
}

func main() {
	length := 5.5
	width := -10

	// call the area function
	result := area(length, float64(width))

	fmt.Println(result)
}
