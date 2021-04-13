package data

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const (
	WORD_MIN_LENGTH = 2
	WORD_MAX_LENGTH = 105
	WORDS_MAX_LENGHT = 105
)

type InputWords struct {
	Len int
	Word string
}

type InputString struct {
	Len int
	InputStr string
}

func ReadDictionaryWords(dictionaryFilePath string) ([]InputWords, error) {
	var (
		inputsWords []InputWords
		totalWords int
		duplicateWords = make(map[string]bool)
	)

	dFile, err := os.Open(dictionaryFilePath)
	if err != nil {
		return inputsWords, errors.New("dictionary file not found")
	}

	defer func() ([]InputWords, error) {
		if err = dFile.Close(); err != nil {
			return inputsWords, err
		}
		return inputsWords, nil
	}()

	dS := bufio.NewScanner(dFile)

	for dS.Scan() {
		v := InputWords{
			Len: len(strings.TrimSpace(dS.Text())),
			Word: strings.TrimSpace(strings.ToLower(dS.Text())),
		}

		if v.Len > WORD_MAX_LENGTH || v.Len < WORD_MIN_LENGTH {
			return inputsWords, errors.New("dictionary words should be minimum 2 and maximum 105 character long")
		}

		if duplicateWords[v.Word] == false {
			inputsWords = append(inputsWords, v)
			totalWords += v.Len
			duplicateWords[v.Word] = true
		}

		if totalWords > WORDS_MAX_LENGHT {
			return inputsWords, errors.New("dictionary total words length should be less than or equal to 105 characters")
		}
	}

	err = dS.Err()
	if err != nil {
		return inputsWords, err
	}

	return inputsWords, nil
}

func ReadScrambledStrings(inputFilePath string) ([]InputString, error) {
	var (
		inputStrings []InputString
	)

	iFile, err := os.Open(inputFilePath)
	if err != nil {
		return inputStrings, err
	}

	defer func() ([]InputString, error) {
		if err = iFile.Close(); err != nil {
			return inputStrings, err
		}
		return inputStrings, nil
	}()

	scmData := bufio.NewScanner(iFile)

	for scmData.Scan() {
		v := InputString{
			Len: len(strings.TrimSpace(scmData.Text())),
			InputStr: strings.TrimSpace(strings.ToLower(scmData.Text())),
		}

		if v.Len > WORD_MAX_LENGTH || v.Len < 2 {
			return inputStrings, errors.New("scrambled string should be less than 105 character")
		}

		inputStrings = append(inputStrings, v)
	}

	err = scmData.Err()
	if err != nil {
		return inputStrings, err
	}

	return inputStrings, nil
}
