package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	JSON_QUOTE      = '"'
	JSON_WHITESPACE = []rune{' ', '\t', '\b', '\n', '\r'}
	JSON_SYNTAX     = []rune{JSON_COMMA, JSON_COLON, JSON_LEFTBRACKET, JSON_RIGHTBRACKET, JSON_LEFTBRACE, JSON_RIGHTBRACE}
	JSON_NUMBER     = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '.', 'e', 'E'}
)

var (
	FALSE_LEN = len(`false`)
	TRUE_LEN  = len(`true`)
	NULL_LEN  = len(`null`)
)

func lexString(str string, idx int) (string, int, error) {
	jsonValue := ""

	if rune(str[idx]) == JSON_QUOTE {
		str = str[idx+1:]
	} else {
		return "", 0, errors.New("No string")
	}

	for _, c := range str {
		if c == JSON_QUOTE {
			return jsonValue, len(jsonValue) + 2, nil
		}
		jsonValue += string(c)
	}

	panic(fmt.Sprintf("Expected end of string quote: %v", str))
}

func lexNumber(str string, idx int) (float64, int, error) {
	jsonNumber := ""

	for _, c := range str[idx:] {
		if containsRune(JSON_NUMBER, c) {
			jsonNumber += string(c)
		} else {
			break
		}
	}

	if len(jsonNumber) == 0 {
		return 0, 0, errors.New("No number")
	}

	if strings.Contains(jsonNumber, ".") {
		value, err := strconv.ParseFloat(jsonNumber, 64)
		if err != nil {
			panic(err)
		}
		return value, len(jsonNumber), nil
	}

	value, err := strconv.ParseInt(jsonNumber, 10, 64)
	if err != nil {
		panic(err)
	}
	return float64(value), len(jsonNumber), nil
}

func lexBool(str string, idx int) (bool, int, error) {
	stringLen := len(str[idx:])

	if stringLen >= TRUE_LEN && str[idx:idx+TRUE_LEN] == "true" {
		return true, TRUE_LEN, nil
	} else if stringLen >= FALSE_LEN && str[idx:idx+FALSE_LEN] == "false" {
		return false, FALSE_LEN, nil
	}

	return false, 0, errors.New("No boolean")
}

func lexNull(str string, idx int) (int, error) {
	stringLen := len(str[idx:])
	if stringLen >= NULL_LEN && str[idx:idx+NULL_LEN] == "null" {
		return NULL_LEN, nil
	}
	return 0, errors.New("No null")
}

func Lex(str *string) *[]any {
	tokens := make([]any, 0)

	idx := 0
	for true {
		if idx >= len(*str) {
			break
		}

		jsonString, cnt, err := lexString(*str, idx)
		if err == nil {
			tokens = append(tokens, jsonString)
			idx += cnt
			continue
		}

		jsonNumber, cnt, err := lexNumber(*str, idx)
		if err == nil {
			tokens = append(tokens, jsonNumber)
			idx += cnt
			continue
		}

		jsonBool, cnt, err := lexBool(*str, idx)
		if err == nil {
			tokens = append(tokens, jsonBool)
			idx += cnt
			continue
		}

		cnt, err = lexNull(*str, idx)
		if err == nil {
			tokens = append(tokens, nil)
			idx += cnt
			continue
		}

		if containsRune(JSON_WHITESPACE, rune((*str)[idx])) {
			idx += 1
		} else if containsRune(JSON_SYNTAX, rune((*str)[idx])) {
			tokens = append(tokens, string((*str)[idx]))
			idx += 1
		} else {
			panic(fmt.Sprintf("Unexpected character: %v", (*str)[idx]))
		}
	}

	return &tokens
}

func containsRune(runes []rune, target rune) bool {
	for _, r := range runes {
		if r == target {
			return true
		}
	}
	return false
}
