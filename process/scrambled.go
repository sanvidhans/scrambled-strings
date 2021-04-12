package process

import (
	"github.com/sanvidhans/scrambled-strings/data"
	"strconv"
	"strings"
)


func getPermutations(originalWord string) []string {
	var permutations []string
	if len(originalWord) <= 3 {
		permutations = append(permutations, originalWord)
		return permutations
	}

	permutations = append(permutations, originalWord)

	firstChar := originalWord[:1]
	lastChar := originalWord[len(originalWord)-1:]

	middleWord := originalWord[1:len(originalWord)-1]

	Perm([]rune(middleWord), func(a []rune){
		wrd := firstChar+string(a)+lastChar
		permutations = append(permutations, wrd)
	})

	return permutations
}

func CountScrambledWord(words []data.InputWords, inputString []data.InputString) map[string]int {
	counts := make(map[string]int)

	wordCount := 0
	for _, w := range words {
		allPossiblePermutations := getPermutations(w.Word)

		for _, v := range allPossiblePermutations {
			lineNumber, found := findSubStringInScrambledString(v, inputString)

			if found {
				wordCount += 1
				caseNumber := strconv.Itoa(lineNumber)

				if counts[caseNumber] == 0 {
					counts[caseNumber] = 1
					wordCount = 1
				}else{
					counts[caseNumber] = wordCount
				}

				break
			}
		}

	}
	return counts
}


func findSubStringInScrambledString(sub string, scrambledStrings []data.InputString) (int, bool) {
	var caseCount int
	var found bool

	if len(scrambledStrings) > 0 {
		var i int
		for i = 0; i<len(scrambledStrings); i++{

			 if strings.Contains(scrambledStrings[i].InputStr, sub){
			 	found = true
			 	break
			 }
		}

		return i+1, found
	}

	return caseCount, false
}