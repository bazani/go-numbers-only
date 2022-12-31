package main

import (
	"fmt"
	"regexp"
	"os"
	"golang.design/x/clipboard"
)

var numbersRegex = regexp.MustCompile(`[\D]+`)

func onlyNumbers(input string) string {
	return numbersRegex.ReplaceAllString(input, "")
}

func stitchArgs(arguments []string) string {
	var answer string

	for _, arg := range arguments {
		answer = answer + arg
	}

	return answer
}

func main() {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("Cannot load the clipboard module.")
		panic(err)
	}

	// get the barcode as arguments, since whitespace is consider a new argument
	input := os.Args[1:]
	// stitch the barcode together
	crudeBarcode := stitchArgs(input)
	// get only the numbers
	barcode := onlyNumbers(crudeBarcode)
	// print the numbers from the barcode
	fmt.Println(barcode)
	// and also copies it to the clipboard
	clipboard.Write(clipboard.FmtText, []byte(barcode))
}

