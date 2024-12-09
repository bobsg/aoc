package main

import (
	"reflect"
	"testing"
)

func TestGenerateOperations(t *testing.T) {
	inputs := []string{"+", "*"}
	got := GenerateOperations(3, inputs)
	want := [][][]string{
		{{}},
		{{"+"}, {"*"}},
		{{"+", "+"}, {"+", "*"}, {"*", "+"}, {"*", "*"}},
		{{"+", "+", "+"}, {"+", "+", "*"}, {"+", "*", "+"}, {"+", "*", "*"}, {"*", "+", "+"}, {"*", "+", "*"}, {"*", "*", "+"}, {"*", "*", "*"}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}
