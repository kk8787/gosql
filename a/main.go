package main

import (
	"fmt"
	"strings"
)

type Token string

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

func lex(str string) []Token {
	tokens := []Token{}

	tmpStr := str
	for len(tmpStr) > 0 {
		parsedString, _tmpStr := lex_string(tmpStr)
		tmpStr = _tmpStr

		char := rune(tmpStr[0])
		fmt.Println("parsedString", parsedString)

		if parsedString != "" {
			tokens = append(tokens, parsedString)
			continue
		}

		// ignore whitespace
		if strings.Contains(string(WHITESPACE), string(char)) {
			tmpStr = tmpStr[1:]
			continue
		}

		if strings.Contains(string(JSON_SYNTAX), string(char)) {
			tokens = append(tokens, Token(string(char)))
			tmpStr = tmpStr[1:]
			continue
		}

		panic(fmt.Sprintf("unknown character: '%s', '%s'", string(char), tmpStr))
	}

	return tokens
}

func parseJson(json string) any {
	// tokens := lex(json)

	return map[string]any{}
}

func main() {
	fmt.Println("Hello, World!")
}
