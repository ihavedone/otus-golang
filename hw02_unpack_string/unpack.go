package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrUnexpectedSlash = errors.New("unexpected slash")
	ErrUnexpectedDigit = errors.New("unexpected digit")
)

func isDigit(current string) bool {
	_, err := strconv.ParseInt(current, 10, 64)
	return err == nil
}

func extractSymbolAndGetTail(input string) (string, string, error) {
	if len(input) == 0 {
		return "", "", nil
	}

	runes, offset := []rune(input), 1
	currentSymbol := string(runes[:offset])
	if isDigit(currentSymbol) {
		return "", "", ErrUnexpectedDigit
	}

	if currentSymbol == "\\" {
		offset = 2
		currentSymbol = string(runes[offset-1 : offset])
		if !isDigit(currentSymbol) && currentSymbol != "\\" {
			return "", "", ErrUnexpectedSlash
		}
	}
	if offset < len(runes) {
		nextRune := string(runes[offset : offset+1])
		if isDigit(nextRune) {
			parsedInt, _ := strconv.ParseInt(nextRune, 10, 64)
			currentSymbol = strings.Repeat(currentSymbol, int(parsedInt))
			offset++
		}
	}
	return currentSymbol, string(runes[offset:]), nil
}

func Unpack(input string) (string, error) {
	inputTail := input
	var result strings.Builder
	var err error
	for {
		var currentSymbol string
		currentSymbol, inputTail, err = extractSymbolAndGetTail(inputTail)
		if err != nil {
			return "", err
		}
		result.WriteString(currentSymbol)
		if len(inputTail) == 0 {
			break
		}
	}
	return result.String(), nil
}
