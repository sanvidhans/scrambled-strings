package data

import "testing"

func TestReadDictionaryWordsWithWrongPath(t *testing.T) {
	dictionaryFilePath := "dictionary.txt"

	_, err := ReadDictionaryWords(dictionaryFilePath)
	if err != nil && err.Error() != "dictionary file not found" {
		t.Errorf("Dictionary with wrong path test failed")
	}
}

func TestReadDictionaryWordsWithCorrectPath(t *testing.T) {
	dictionaryFilePath := "test_dictionary_file.txt"

	_, err := ReadDictionaryWords(dictionaryFilePath)
	if err != nil {
		t.Errorf("Dictionary file with valid path test failed")
	}
}

func TestReadDictionaryWordsWithMinimumLengthOfWord(t *testing.T) {
	dictionaryFilePath := "test_dictionary_minlen_file.txt"

	_, err := ReadDictionaryWords(dictionaryFilePath)
	if err != nil && err.Error() != "dictionary words should be minimum 2 and maximum 105 character long" {
		t.Errorf("Minimum length of dictionary word test failed")
	}
}

func TestReadDictionaryWordsWithMaximumLengthOfWord(t *testing.T) {
	dictionaryFilePath := "test_dictionary_maxlen_file.txt"

	_, err := ReadDictionaryWords(dictionaryFilePath)
	if err != nil && err.Error() != "dictionary words should be minimum 2 and maximum 105 character long" {
		t.Errorf("Maximum length of dictionary word test failed")
	}
}

//func TestReadDictionaryWordsWithEmptyFile(t *testing.T) {
//	DictionaryFilePath = "test_dictionary_empty_file.txt"
//
//	_, err := ReadDictionaryWords()
//	if err != nil && err.Error() != "" {
//		t.Errorf("Empty dictionary file test failed")
//	}
//}

func TestReadDictionaryWordsWithTotalWordMaxLengthFile(t *testing.T) {
	dictionaryFilePath := "test_dictionary_total_max_file.txt"

	_, err := ReadDictionaryWords(dictionaryFilePath)
	if err != nil && err.Error() != "dictionary total words length should be less than or equal to 105 characters" {
		t.Errorf("Total word length of dictionary words test failed")
	}
}