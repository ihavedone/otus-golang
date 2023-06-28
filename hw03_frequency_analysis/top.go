package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const expectedSliceSize = 10

var (
	blackListWords  = []string{"", "-"}
	separatorRegexp = regexp.MustCompile(`[-]?[\s,.!?'":;@#$%^&*(){}<>\[\]_]+[-]?`)
)

type Word struct {
	value   string
	counter int
}
type Words []Word

func (words *Words) add(text string) {
	for _, v := range blackListWords {
		if text == v {
			return
		}
	}
	*words = append(*words, Word{value: text, counter: 1})
}

func (words *Words) Calc(text string) {
	for indx := range *words {
		if (*words)[indx].value == text {
			(*words)[indx].counter++
			return
		}
	}
	words.add(text)
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
	output := make(Words, 0)
	for _, word := range words {
		output = append(output, word)
		if len(output) >= n {
			break
		}
	}
	return output
}

func (words Words) ToStrings() []string {
	output := make([]string, 0)
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
	return words.Sort().Head(expectedSliceSize).ToStrings()
}
