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

func ReadDictionary() ([]InputWords, error) {
	var (
		inputsWords []InputWords
		totalWords int
	)

	dFile, err := os.Open(DictionaryFilePath)
	if err != nil {
		return inputsWords, err
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
			Word: strings.TrimSpace(dS.Text()),
		}

		if v.Len > WORD_MAX_LENGTH || v.Len < WORD_MIN_LENGTH {
			return inputsWords, errors.New("dictionary words are with low length")
		}

		inputsWords = append(inputsWords, v)
		totalWords += v.Len

		if totalWords > WORDS_MAX_LENGHT {
			return inputsWords, errors.New("dictionary words are with out off length")
		}
	}

	err = dS.Err()
	if err != nil {
		return inputsWords, err
	}

	return inputsWords, nil
}

func ReadScrambledStrings() ([]InputString, error) {
	var (
		inputStrings []InputString
	)

	iFile, err := os.Open(InputFilePath)
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
			InputStr: strings.TrimSpace(scmData.Text()),
		}

		if v.Len > WORD_MAX_LENGTH || v.Len < 2 {
			return inputStrings, errors.New("scrambled string is should be less than 105 character")
		}

		inputStrings = append(inputStrings, v)
	}

	err = scmData.Err()
	if err != nil {
		return inputStrings, err
	}

	return inputStrings, nil
}
