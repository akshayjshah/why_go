package main

import "fmt"
// START OMIT
type Language struct {
	name string
	trans map[int]string
}

func main() {
	dict := make(map[int]string)
	dict[1] = "one"
	dict[2] = "two"

	eng := &Language{name: "English", trans: dict}

	for num, word := range eng.trans {
		fmt.Printf("In English, %d is %s.\n", num, word)
	}
}
// END OMIT
