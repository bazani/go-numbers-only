package numbers_only

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

func main() {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("Cannot load the clipboard module.")
		panic(err)
	}

	barcode := os.Args[1]
	answer := onlyNumbers(barcode)
	fmt.Println(answer)
	clipboard.Write(clipboard.FmtText, []byte(answer))
}

