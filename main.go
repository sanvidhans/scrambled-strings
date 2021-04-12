package main

import (
	"fmt"
	"github.com/sanvidhans/scrambled-strings/data"
	"github.com/sanvidhans/scrambled-strings/process"
	"github.com/sanvidhans/scrambled-strings/validation"
	"log"
	"runtime"
)

func init(){
	err := data.AutoLoadFiles()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func main(){
	//validateFiles()
	//dictionary := flag.String("dictionary", "", "a string")
	//input := flag.String("input", "", "a string")
	validation.ValidateFiles()

	dictionaryWords, err := data.ReadDictionary()
	if err != nil {
		fmt.Printf("parsing disctionary file failed, Error: %s\n", err.Error())
		runtime.Breakpoint()
	}

	scrambledStrings, err := data.ReadScrambledStrings()
	if err != nil {
		fmt.Printf("parsing input file failed, Error: %s\n", err.Error())
		runtime.Breakpoint()
	}

	counts := process.CountScrambledWord(dictionaryWords, scrambledStrings)

	for k, va := range counts {
		fmt.Println(k, ":", va)
	}
}
