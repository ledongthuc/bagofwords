package bagofwords

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	b := Bag{
		IgnoreWords: []string{
			"Out",
		},
		CaseSensitive: false,
	}

	input := `
	Out believe has request not how comfort evident. Up delight cousins we feeling minutes. Genius has looked end piqued spring. Down has rose feel find man. Learning day desirous informed expenses material returned six the. She enabled invited exposed him another. Reasonably conviction solicitude me mr at discretion reasonable. Age out full gate bed day lose.
	`
	output, err := b.Analyze(input)
	if err != nil {
		t.Errorf("FAIL: input %s, output %+v", input, err)
	}

	if output.CountWord != 54 && output.CountVocabulary != 51 {
		t.Errorf("FAIL: input %s, Word count: %+v, Vocabulary count: %+v", input, output.CountWord, output.CountVocabulary)
	}

}

func TestIsIgnoreWord(t *testing.T) {
	ignoreWords := []string{
		"",
		"wing",
		"button",
		"colour",
		"Excellent",
	}
	testCases := map[string]bool{
		"wing":      true,
		"ing":       false,
		"wwing":     false,
		"Wing":      false,
		"Excellent": true,
	}
	b := Bag{
		IgnoreWords:   ignoreWords,
		CaseSensitive: true,
	}

	for input, expectedOutput := range testCases {
		if output := b.IsIgnoreWord(input); output != expectedOutput {
			t.Errorf("FAIL: input %s, output: %v, expected output %v", input, output, expectedOutput)
		}
	}
}

func TestCount(t *testing.T) {
	testCases := map[*map[string]int64][]int64{
		&map[string]int64{
			"third":   5,
			"part":    3,
			"hapless": 1,
		}: []int64{9, 3},
		&map[string]int64{
			"third":   0,
			"part":    0,
			"hapless": 0,
		}: []int64{0, 0},
	}

	for input, expectedOutputs := range testCases {
		if output1, output2 := Count(*input); output1 != expectedOutputs[0] || output2 != expectedOutputs[1] {
			t.Errorf("FAIL: input %+v, output1: %+v, output2: %+v,expected output %+v", input, output1, output2, expectedOutputs)
		}
	}

}
