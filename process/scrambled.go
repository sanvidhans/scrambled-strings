package process

import (
	"github.com/sanvidhans/scrambled-strings/data"
	"strings"
)


func GetMiddleWordPermutations(originalWord string) []string {
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

func CountScrambledWord(words []data.InputWords, inputString []data.InputString) map[int]int {
	wCounts := make(map[int]int)

	wordCount := 0
	for _, w := range words {
		allPossiblePermutations := GetMiddleWordPermutations(w.Word)

		for _, v := range allPossiblePermutations {
			lineNumber, found := FindSubStringInScrambledString(v, inputString)

			if found {
				wordCount += 1

				if wCounts[lineNumber] == 0 {
					wCounts[lineNumber] = 1
					wordCount = 1
				}else{
					wCounts[lineNumber] = wordCount
				}

				break
			}
		}

	}

	return wCounts
}

func FindSubStringInScrambledString(sub string, scrambledStrings []data.InputString) (int, bool) {
	var caseCount int
	var found bool
	if len(scrambledStrings) > 0 {

		for i := 0; i < len(scrambledStrings); i++ {
			caseCount = i
			 if strings.Contains(scrambledStrings[i].InputStr, sub){
			 	found = true
			 	break
			 }
		}

		return caseCount+1, found
	}

	return caseCount, false
}