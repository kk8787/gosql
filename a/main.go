package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Token any

const (
	JSON_QUOTE rune = '"'
)

var WHITESPACE = []rune{' ', '\n', '\t', '\\'}
var JSON_SYNTAX = []rune{':', ',', '{', '}'}

func lex_string(jsonString string) (Token, string) {
	if rune(jsonString[0]) != JSON_QUOTE {
		return "", jsonString
	}

	newString := jsonString[1:]
	tokenString := ""

	for _, c := range newString {
		if c != JSON_QUOTE {
			tokenString = tokenString + string(c)
			continue
		}

		return Token(tokenString), newString[len(tokenString)+1:]
	}

	panic("unclosed string")
}

var NumberString = []rune{'-', 'e', '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func lex_number(jsonString string) (Token, string) {
	tokenString := ""

	for _, c := range jsonString {
		if !strings.Contains(string(NumberString), string(c)) {
			break
		}

		tokenString = tokenString + string(c)
	}

	rest := jsonString[len(tokenString):]

	fmt.Println("tokenString", tokenString)

	if tokenString == "" {
		return "", jsonString
	}

	if strings.Contains(tokenString, ".") {
		parsed, err := strconv.ParseFloat(tokenString, 64)
		if err != nil {
			panic(fmt.Sprintf("failed to parse float: '%s', '%s'", tokenString, jsonString))
		}

		return parsed, rest
	}

	parsed, err := strconv.ParseInt(tokenString, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int: '%s', '%s'", tokenString, jsonString))
	}

	return parsed, rest
}

func lex(str string) []Token {
	tokens := []Token{}

	tmpStr := str
	for len(tmpStr) > 0 {
		parsedString, _tmpStr := lex_string(tmpStr)
		tmpStr = _tmpStr

		if parsedString != "" {
			tokens = append(tokens, parsedString)
			continue
		}

		parsedNumber, _tmpStr := lex_number(tmpStr)
		tmpStr = _tmpStr

		if parsedNumber != "" {
			tokens = append(tokens, parsedNumber)
			continue
		}

		char := rune(tmpStr[0])
		fmt.Println("parsedString", parsedString)

		// ignore whitespace
		if strings.Contains(string(WHITESPACE), string(char)) {
			tmpStr = tmpStr[1:]
			continue
		}

		if strings.Contains(string(JSON_SYNTAX), string(char)) {
			tokens = append(tokens, string(char))
			tmpStr = tmpStr[1:]
			continue
		}

		panic(fmt.Sprintf("unknown character: '%s', '%s'", string(char), tmpStr))
	}

	return tokens
}

func parseJson(json string) map[string]any {
	// tokens := lex(json)

	return map[string]any{}
}

func main() {
	fmt.Println("Hello, World!")
}
