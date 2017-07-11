package bagofwords

import (
	"errors"
	"strings"
)

type Bag struct {
	CaseSensitive bool
	IgnoreWords   []string
	WordLength    int
}

func CreateABag(isCaseSensitive bool, ignoreWords []string, wordLength int) Bag {
	return Bag{
		CaseSensitive: isCaseSensitive,
		IgnoreWords:   ignoreWords,
		WordLength:    wordLength,
	}
}

func (bag Bag) Analyze(content string) (Response, error) {
	if content == "" {
		return Response{}, errors.New("Content is empty")
	}

	if !bag.CaseSensitive {
		content = strings.ToLower(content)
	}

	dictionary := make(map[string]int64)
	splitedStrings := strings.Split(content, " ")
	var buffer = make([]string, bag.WordLength)
	for index, word := range splitedStrings {
		word = strings.Trim(word, "~!@#$%^&*()_+-=`{}[]|\\:\";'<>?,./ \n\r\t")
		if word == "" || bag.IsIgnoreWord(word) {
			continue
		}

		if len(buffer) >= bag.WordLength {
			buffer = buffer[1:]
		}
		buffer = append(buffer, word)
		if bag.WordLength > index+1 {
			continue
		}

		savedWord := strings.Join(buffer, " ")

		if counter, ok := dictionary[savedWord]; ok {
			dictionary[savedWord] = counter + 1
		} else {
			dictionary[savedWord] = 1
		}
	}

	wordCounter, vocabularyCounter := Count(dictionary)
	return Response{Dictionary: dictionary, CountVocabulary: vocabularyCounter, CountWord: wordCounter}, nil
}

func (bag Bag) IsIgnoreWord(word string) bool {
	if !bag.CaseSensitive {
		word = strings.ToLower(word)
	}

	for _, ignoreWord := range bag.IgnoreWords {
		if !bag.CaseSensitive {
			ignoreWord = strings.ToLower(ignoreWord)
		}

		if ignoreWord == word {
			return true
		}
	}
	return false
}

func Count(dictionary map[string]int64) (wordCounter, vocabularyCounter int64) {
	for _, count := range dictionary {
		if count <= 0 {
			continue
		}

		wordCounter += count
		vocabularyCounter++
	}
	return
}
