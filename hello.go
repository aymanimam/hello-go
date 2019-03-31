package main

import (
	"fmt"
	"github.com/aymanimam/hello-go/gopherdojo/slide331/server"
)

func main() {
	server.StartServer()
}

// Old tests
func ConcatStr(s1, s2 string) string {
	return s1 + s2
}

func PrintOddEven() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 1:
			fmt.Printf("%d- Odd\n", i)
		case i%2 == 0:
			fmt.Printf("%d- Even\n", i)
		default:
			fmt.Printf("Error!")
		}
	}
}
