package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const EXPECTED_SLICE_SIZE = 10

var blackListWords = []string{"", "-"}
var separatorRegexp = regexp.MustCompile(`[-]?[\s,.!?'":;@#$%^&*(){}<>\[\]_]+[-]?`)

type Word struct {
	value   string
	counter int
}
type Words []Word

func (wordsPtr *Words) add(text string) {
	for _, v := range blackListWords {
		if text == v {
			return
		}
	}
	*wordsPtr = append(*wordsPtr, Word{value: text, counter: 1})
}

func (wordsPtr *Words) Calc(text string) {
	for indx := range *wordsPtr {
		if (*wordsPtr)[indx].value == text {
			(*wordsPtr)[indx].counter++
			return
		}
	}
	wordsPtr.add(text)
}

func (words Words) Sort() Words {
	sort.Slice(words, func(i, j int) bool {
		if (words)[i].counter == (words)[j].counter {
			return (words)[i].value < (words)[j].value
		}
		return (words)[i].counter > (words)[j].counter
	})
	return words
}

func (words Words) Head(n int) Words {
	var output Words
	for _, word := range words {
		output = append(output, word)
		if len(output) >= n {
			break
		}
	}
	return output
}

func (words Words) ToStrings() []string {
	var output []string
	for _, word := range words {
		output = append(output, word.value)
	}
	return output
}

func Top10(text string) []string {
	tokens := separatorRegexp.Split(text, -1)
	var words Words
	for _, token := range tokens {
		words.Calc(strings.ToLower(token))
	}
	return words.Sort().Head(EXPECTED_SLICE_SIZE).ToStrings()
}
