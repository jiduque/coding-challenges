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
	inputs := []string{"{\"key\": \"value\"}", "{\n\"key\": \"value\"\n}"}
	expectedValues := []internal.Tokens{
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.String, Value: "key"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.String, Value: "value"},
			internal.Token{Type: internal.RightBrace, Value: "}"},
		},
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.WhiteSpace, Value: "\n"},
			internal.Token{Type: internal.String, Value: "key"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.String, Value: "value"},
			internal.Token{Type: internal.WhiteSpace, Value: "\n"},
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

func TestSeveralValues(t *testing.T) {
	inputs := []string{
		"{\n  \"key1\": true,\n  \"key2\": false,\n  \"key3\": null,\n  \"key4\": \"value\",\n  \"key5\": 101\n}",
	}

	expectedValues := []internal.Tokens{
		{
			internal.Token{Type: internal.LeftBrace, Value: "{"},
			internal.Token{Type: internal.WhiteSpace, Value: "\n  "},
			internal.Token{Type: internal.String, Value: "key1"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.Boolean, Value: "true"},
			internal.Token{Type: internal.Comma, Value: ","},
			internal.Token{Type: internal.WhiteSpace, Value: "\n  "},
			internal.Token{Type: internal.String, Value: "key2"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.Boolean, Value: "false"},
			internal.Token{Type: internal.Comma, Value: ","},
			internal.Token{Type: internal.WhiteSpace, Value: "\n  "},
			internal.Token{Type: internal.String, Value: "key3"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.NULL, Value: "null"},
			internal.Token{Type: internal.Comma, Value: ","},
			internal.Token{Type: internal.WhiteSpace, Value: "\n  "},
			internal.Token{Type: internal.String, Value: "key4"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.String, Value: "value"},
			internal.Token{Type: internal.Comma, Value: ","},
			internal.Token{Type: internal.WhiteSpace, Value: "\n  "},
			internal.Token{Type: internal.String, Value: "key5"},
			internal.Token{Type: internal.Colon, Value: ":"},
			internal.Token{Type: internal.WhiteSpace, Value: " "},
			internal.Token{Type: internal.Numeric, Value: "101"},
			internal.Token{Type: internal.WhiteSpace, Value: "\n"},
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
