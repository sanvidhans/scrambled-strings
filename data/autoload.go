package data

import (
	"flag"
	"log"
)

var DictionaryFilePath, InputFilePath string

func init(){
	flag.StringVar(&DictionaryFilePath, "dictionary", "", "a dictionary file path")
	flag.StringVar(&InputFilePath, "input", "", "a input file path")

	flag.Parse()
	if DictionaryFilePath == "" || InputFilePath == "" {
		log.Println("Please enter dictionary and input file path.")
		return
	}
}