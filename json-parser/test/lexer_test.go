package test

import (
	"json-parser/internal"
	"testing"
)

func TestEmptyObject(t *testing.T) {
	inputs := []string{"{}", "{ }", "{\n}", "{     }"}
	expectedValues := []internal.Tokens{
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.RightBrace, Value: "}"},
		},
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.RightBrace, Value: "}"},
		},
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.WhiteSpace, Value: "\n"},
			internal.Token{Type: internal.RightBrace, Value: "}"},
		},
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.WhiteSpace, Value: "     "},
			internal.Token{Type: internal.RightBrace, Value: "}"},
		},
	}

	for i, input := range inputs {
		expected := expectedValues[i]
		output := internal.LexicallyAnalyze(input)
		if !expected.Equals(output) {
			t.Errorf("Could not analyze input: %s \n\tExpected: %s\n\tGot: %s", input, expected.String(), output.String())
		}
	}
}

func TestKeyValueString(t *testing.T) {
	inputs := []string{"{\"key\": \"value\"}"}
	expectedValues := []internal.Tokens{
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.String, Value: "key"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.String, Value: "value"},
			internal.Token{Type: internal.RightBrace, Value: "}"},
		},
	}

	for i, input := range inputs {
		expected := expectedValues[i]
		output := internal.LexicallyAnalyze(input)
		if !expected.Equals(output) {
			t.Errorf("Could not analyze input: %s \n\tExpected: %s\n\tGot: %s", input, expected.String(), output.String())
		}
	}
}
