package data

import "testing"

func TestReadDictionary(t *testing.T) {
	DictionaryFilePath = "sampleFiles/dictionary.txt"

	_, err := ReadDictionary()
	if err != nil {
		t.Errorf("error while loading dictionary file")
	}
}
