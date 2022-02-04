package main

import (
	"fmt"
	"os"
)

func main() {
	// Validate command line parameters
	if len(os.Args) != 3 {
		Usage()
	}
}

// Usage prints helper message and exit
func Usage() {
	fmt.Println("Usage:\n\tcaveminer.exe <cave size>")
	fmt.Println("Example:\n\tcaveminer.exe calc.exe 80")
	os.Exit(0)
}
