package main

import (
	"reflect"
	"testing"
)

func TestGenerateOperations(t *testing.T) {

	got := GenerateOperations("+", "*", 3)
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
