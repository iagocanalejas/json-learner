package main

import (
	"reflect"
	"testing"
)

func TestEmptyObject(t *testing.T) {
	result := make(map[string]interface{}, 0)
	if !reflect.DeepEqual(FromString("{}"), result) {
		t.Fail()
	}
}

func TestBasicObject(t *testing.T) {
	if FromString(`{"foo":"bar"}`)["foo"] != "bar" {
		t.Fail()
	}
}

func TestBasicNumber(t *testing.T) {
	if FromString(`{"foo":1}`)["foo"] != float64(1) {
		t.Fail()
	}
}

func TestEmptyArray(t *testing.T) {
	if reflect.DeepEqual(FromString(`{"foo":[]}`)["foo"], [0]any{}) {
		t.Fail()
	}
}

func TestBasicArray(t *testing.T) {
	if reflect.DeepEqual(FromString(`{"foo":[1,2,"three"]}`)["foo"], [3]any{1, 2, "three"}) {
		t.Fail()
	}
}

func TestNestedObject(t *testing.T) {
	result := make(map[string]interface{}, 1)
	result["bar"] = 2
	if reflect.DeepEqual(FromString(`{"foo":{"bar":2}}`)["foo"], result) {
		t.Fail()
	}
}

func TestTrue(t *testing.T) {
	if FromString(`{"foo":true}`)["foo"] != true {
		t.Fail()
	}
}

func TestFalse(t *testing.T) {
	if FromString(`{"foo":false}`)["foo"] != false {
		t.Fail()
	}
}

func TestNull(t *testing.T) {
	if FromString(`{"foo":null}`)["foo"] != nil {
		t.Fail()
	}
}

func TestBasicWhitespace(t *testing.T) {
	if FromString(`{ "foo" : [1, 2, "three"] }`)["foo"] == nil {
		t.Fail()
	}
}

func TestToString(t *testing.T) {
	if ToString(FromString(`{"foo":1}`)) != `{"foo": 1}` {
		t.Fail()
	}
}
