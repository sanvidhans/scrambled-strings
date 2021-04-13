package main

import (
	"flag"
	"fmt"
	"github.com/sanvidhans/scrambled-strings/data"
	"github.com/sanvidhans/scrambled-strings/process"
	"sort"
)

var (
	dictionaryFilePath string
	inputFilePath string
)

func init()  {

	flag.StringVar(&dictionaryFilePath, "dictionary", "", "a dictionary file path")
	flag.StringVar(&inputFilePath, "input", "", "a input file path")

	flag.Parse()
}

func main(){

	if dictionaryFilePath == "" || inputFilePath == "" {
		fmt.Println("Error: please provide dictionary and input file path")
		return
	}

	dictionaryWords, err := data.ReadDictionaryWords(dictionaryFilePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	scrambledStrings, err := data.ReadScrambledStrings(inputFilePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	//fmt.Println(scrambledStrings)
	counts := process.CountScrambledWord(dictionaryWords, scrambledStrings)

	//sorting map
	var sortCounts []int
	for k, _ := range counts {
		sortCounts = append(sortCounts, k)
	}
	sort.Ints(sortCounts)

	for _, v := range sortCounts {
		fmt.Printf("Case#%d: %d\n", v, counts[v])
	}
}
