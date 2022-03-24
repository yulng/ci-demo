package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("HOME-->" + os.Getenv("HOME"))
}
