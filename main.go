package main

import (
	"fmt"
	"github.com/sanvidhans/scrmabled-strings/data"
	"github.com/sanvidhans/scrmabled-strings/scrambled"
	"github.com/sanvidhans/scrmabled-strings/validation"
)

func main(){
	//validateFiles()
	//dictionary := flag.String("dictionary", "", "a string")
	//input := flag.String("input", "", "a string")
	validation.ValidateFiles()

	dictionaryWords, err := data.ReadDictionary()
	if err != nil {
		fmt.Printf("parsing disctionary file failed, Error: %s\n", err.Error())
	}

	scrambledStrings, err := data.ReadScrambledStrings()
	if err != nil {
		fmt.Printf("parsing input file failed, Error: %s\n", err.Error())
	}

	counts := scrambled.CountScrambledWord(dictionaryWords, scrambledStrings)

	for k, va := range counts {
		fmt.Println(k, ":", va)
	}
}
