package validation

import (
	"github.com/sanvidhans/scrmabled-strings/data"
	"log"
)

func ValidateFiles()  {
	if data.DictionaryFilePath == "" || data.InputFilePath == "" {
		log.Println("Please enter dictionary and input file path.")
		return
	}
}