package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Gopher!")
	fmt.Println(os.ExpandEnv("${GOPATH}"))
	fmt.Print("Hello, Marc. The command 'fmt.Print' adds not newline.")
	fmt.Println("A", "B", "C", "D")
	fmt.Print("A", "B", "C")
}
