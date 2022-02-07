package main

import (
	"debug/pe"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Validate command line parameters
	if len(os.Args) != 3 {
		Usage()
	}

	// Check if we got the correct number for cave size
	caveSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	} else if caveSize <= 0 {
		fmt.Println("Error: <cave size> must be bigger than zero")
		Usage()
	}

	// Try to open input file or exit program otherwise
	f, err := pe.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Loop through all PE sections
	sep := strings.Repeat("-", 35)
	fmt.Println(sep)
	for _, s := range f.Sections {
		fmt.Printf("Section %s (%d bytes)\n", s.Name, s.Size)
		Dig(s, caveSize)
		fmt.Println(sep)
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
