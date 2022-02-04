package main

import (
	"debug/pe"
	"fmt"
	"os"
)

func main() {
	// Validate command line parameters
	if len(os.Args) != 3 {
		Usage()
	}
}

// Dig searches PE section for a code cave that is at least n bytes in size
func Dig(s *pe.Section, n int) {

}

// Usage prints helper message and exit
func Usage() {
	fmt.Println("Usage:\n\tcaveminer.exe <cave size>")
	fmt.Println("Example:\n\tcaveminer.exe calc.exe 80")
	os.Exit(0)
}
