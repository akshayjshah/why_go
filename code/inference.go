package main

import "fmt"

// START OMIT
func main() {
	greeting := "Hello, 世界"
	fmt.Println(greeting)

	unicode := []rune(greeting)
	for pos, point := range unicode {
		fmt.Printf("Rune %d is at position %d.\n", point, pos)
	}
}
// END OMIT
