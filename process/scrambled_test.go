package process

import (
	"github.com/sanvidhans/scrambled-strings/data"
	"testing"
)

func TestGetMiddleWordPermutations(t *testing.T) {
	testWord := "abc"
	wrdMaps := make(map[string]bool)

	perWords := GetMiddleWordPermutations(testWord)
	for _, v := range perWords {
		wrdMaps[v] = true
	}

	if wrdMaps["abc"] == false {
		t.Error("GetMiddle world Permutation function test failed")
	}

	testWord2 := "pabc"
	wrdMaps2 := make(map[string]bool)

	perWords2 := GetMiddleWordPermutations(testWord2)
	for _, v := range perWords2 {
		wrdMaps2[v] = true
	}

	if wrdMaps2["pabc"] == false || wrdMaps2["pbac"] == false {
		t.Error("GetMiddle world Permutation function test failed")
	}
}

func TestFindSubStringInScrambledString(t *testing.T) {
	word := "acbd"
	scrambledString := []data.InputString{
		{
			Len: len("jkdhfgsdrtacbditdks"),
			InputStr: "jkdhfgsdrtacbditdks",
		},
	}

	_, found := FindSubStringInScrambledString(word, scrambledString)
	if found == false {
		t.Error("FindSubStringInScrambledString test failed")
	}


	word2 := "acbd"
	scrambledString2 := []data.InputString{
		{
			Len: len("jkdhfgsdrtakiditdks"),
			InputStr: "jkdhfgsdrtakiditdks",
		},
	}
	_, found2 := FindSubStringInScrambledString(word2, scrambledString2)
	if found2 == true {
		t.Error("FindSubStringInScrambledString test failed")
	}

}


func TestCountScrambledWord(t *testing.T) {
	words := []data.InputWords{
		{Len: len("axpaj"), Word: "axpaj"},
		{Len: len("apxaj"), Word: "apxaj"},
		{Len: len("dnrbt"), Word: "dnrbt"},
		{Len: len("pjxdn"), Word: "pjxdn"},
		{Len: len("abd"), Word: "abd"},
	}

	scrambledString := []data.InputString{
		{
			Len: len("aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt"),
			InputStr: "aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt",
		},
	}
	counts := CountScrambledWord(words, scrambledString)
	if counts[1] != 4 {
		t.Error("CountScrambledWord function test failed")
	}
}