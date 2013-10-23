package main

import (
	"fmt"
	"os"
)

// START OMIT
func zen() (text string, err error) {
	file, err := os.Open("code/zen.txt")
	if err != nil {
		return
	}
	buf := make([]byte, 1000)
	n, err := file.Read(buf)
	if err != nil {
		return
	}
	return string(buf[:n]), err
}

func main() {
	zen, err := zen()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(zen)
}

// END OMIT
