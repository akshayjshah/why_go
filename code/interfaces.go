package main

import (
	"fmt"
	"net/http"
	"os"
)
// START OMIT
type Readable interface {
	Read(p []byte) (n int, err error)
}

func head(r Readable) string {
	buf := make([]byte, 50)
	n, _ := r.Read(buf)
	return string(buf[:n])
}

func main() {
	file, _ := os.Open("code/zen.txt")
	fmt.Println(head(file))

	response, _ := http.Get("http://hearsaysocial.com")
	fmt.Println(head(response.Body))
}
// END OMIT
