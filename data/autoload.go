package data

import (
	"flag"
	"fmt"
)

var DictionaryFilePath, InputFilePath string


func AutoLoadFiles() error {
	flag.StringVar(&DictionaryFilePath, "dictionary", "", "a dictionary file path")
	flag.StringVar(&InputFilePath, "input", "", "a input file path")

	flag.Parse()
	if DictionaryFilePath == "" || InputFilePath == "" {
		return fmt.Errorf("please provide dictionary and input file path")
	}

	return nil
}