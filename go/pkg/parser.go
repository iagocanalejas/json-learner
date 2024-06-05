package main

import (
	"fmt"
)

func parseArray(tokens []any, idx int) ([]any, int) {
	jsonArray := make([]any, 0)

	lIdx := 1 // 0 = [
	if tokens[idx+lIdx] == string(JSON_RIGHTBRACKET) {
		return jsonArray, lIdx + 1
	}

	for true {
		json, cnt := Parse(tokens, idx+lIdx)
		jsonArray = append(jsonArray, json)
		lIdx += cnt

		if tokens[idx+lIdx] == string(JSON_RIGHTBRACKET) {
			return jsonArray, lIdx + 1
		} else if tokens[idx+lIdx] != string(JSON_COMMA) {
			panic(fmt.Sprintf("expected comma after pair in object, got %v", tokens[idx+lIdx]))
		}
		lIdx += 1
	}

	panic("this should have returned")
}

func parseObject(tokens []any, idx int) (map[string]interface{}, int) {
	jsonObject := make(map[string]interface{})

	lIdx := 1 // 0 = {
	if tokens[idx+lIdx] == string(JSON_RIGHTBRACE) {
		return jsonObject, lIdx + 1
	}

	for true {
		jsonKey, ok := tokens[idx+lIdx].(string)
		if !ok {
			panic(fmt.Sprintf("expected string key, got: %v", tokens[idx+lIdx]))
		}
		lIdx += 1

		if tokens[idx+lIdx] != string(JSON_COLON) {
			panic(fmt.Sprintf("expected colon after key in object, got: %v", tokens[idx+lIdx]))
		}
		lIdx += 1

		jsonValue, cnt := Parse(tokens, idx+lIdx)

		lIdx += cnt
		jsonObject[jsonKey] = jsonValue

		if tokens[idx+lIdx] == string(JSON_RIGHTBRACE) {
			return jsonObject, lIdx + 1
		} else if tokens[idx+lIdx] != string(JSON_COMMA) {
			panic(fmt.Sprintf("expected comma after pair in object, got %v", tokens[idx+lIdx]))
		}
		lIdx += 1
	}

	panic("this should have returned")
}

func Parse(tokens []any, idx int) (any, int) {
	if idx == 0 && tokens[idx] != string(JSON_LEFTBRACE) {
		panic("root must be an object")
	}

	if tokens[idx] == string(JSON_LEFTBRACKET) {
		return parseArray(tokens, idx)
	} else if tokens[idx] == string(JSON_LEFTBRACE) {
		return parseObject(tokens, idx)
	} else {
		return tokens[idx], 1
	}
}
