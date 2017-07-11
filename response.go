package bagofwords

import (
	"bytes"
	"fmt"
)

type Response struct {
	Dictionary      map[string]int64
	CountVocabulary int64
	CountWord       int64
}

func (r Response) UniqueWords() (words []string) {
	words = make([]string, 0, len(r.Dictionary))
	for word := range r.Dictionary {
		words = append(words, word)
	}
	return words
}

func (r Response) Percent(key string) float64 {
	count, ok := r.Dictionary[key]
	if !ok {
		return 0.0
	}

	return float64(count) * 100.0 / float64(r.CountWord)
}

func (r Response) ToCSV() string {
	content := bytes.NewBufferString("")
	content.WriteString(fmt.Sprintf("%s,%d\n", "Words count", r.CountWord))
	content.WriteString(fmt.Sprintf("%s,%d\n\n\n\n", "Unique words count", r.CountVocabulary))

	content.WriteString(fmt.Sprintf("%s,%s\n", "Keyword", "Word counting"))
	for word, count := range r.Dictionary {
		content.WriteString(fmt.Sprintf("%s,%d\n", word, count))
	}
	return content.String()
}
