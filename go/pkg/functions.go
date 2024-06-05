package main

import (
	"fmt"
	"reflect"
)

func FromString(str string) map[string]interface{} {
	tokens := Lex(&str)
	res, _ := Parse(*tokens, 0)
	result, _ := res.(map[string]interface{})
	return result
}

func ToString(json any) string {
	if reflect.ValueOf(json).Kind() == reflect.Map {
		s := "{"
		mapLen := len(json.(map[string]interface{}))

		i := 0
		for key, value := range json.(map[string]interface{}) {
			s += fmt.Sprintf("\"%v\": %v", key, ToString(value))
			if i < mapLen-1 {
				s += ", "
			} else {
				s += "}"
			}
			i += 1
		}
		return s
	} else if reflect.ValueOf(json).Kind() == reflect.Array {
		s := "["
		arrayLen := len(json.([]interface{}))

		for i, value := range json.([]interface{}) {
			s += ToString(value)
			if i < arrayLen-1 {
				s += ", "
			} else {
				s += "]"
			}
		}
		return s
	} else if reflect.ValueOf(json).Kind() == reflect.String {
		return fmt.Sprintf("\"%v\"", json)
	} else if reflect.ValueOf(json).Kind() == reflect.Bool {
		if json.(bool) {
			return "true"
		} else {
			return "false"
		}
	} else if json == nil {
		return "null"
	}
	return fmt.Sprintf("%v", json)
}
