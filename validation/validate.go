package validation

import (
	"fmt"
	"github.com/sanvidhans/scrambled-strings/data"
)

func ValidateFiles()  {
	if data.DictionaryFilePath == "" || data.InputFilePath == "" {
		fmt.Println("please enter dictionary and input file path")
		return
	}
}