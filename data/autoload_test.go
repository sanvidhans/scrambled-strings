package data

import "testing"

func TestAutoLoadFiles(t *testing.T) {
	DictionaryFilePath = ""
	InputFilePath = ""

	if AutoLoadFiles() != nil {
		if AutoLoadFiles().Error() != "please provide dictionary and input file path" {
			t.Errorf("Enter disctory and inpt")
		}
	}
}

func TestAutoLoadFilesWithFilePath(t *testing.T) {
	DictionaryFilePath = "sampleFiles/dictionary.txt"
	InputFilePath = "sampleFiles/inputs.txt"

	if AutoLoadFiles() != nil {
		if AutoLoadFiles().Error() == "please provide dictionary and input file path" {
			t.Errorf("Enter dictionary file path")
		}
	}
}