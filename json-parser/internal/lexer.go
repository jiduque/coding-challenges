package internal

import (
	"bytes"
	"strconv"
	"strings"
)

type TokenType int

const (
	LeftBrace TokenType = iota
	RightBrace
	LeftBracket
	RightBracket
	Colon
	Comma

	WhiteSpace
	Numeric
	String
	Boolean
	NULL
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) Equals(other Token) bool {
	if t.Type != other.Type {
		return false
	}
	if t.Value != other.Value {
		return false
	}
	return true
}

func (t Token) String() string {
	return "{ Type: " + strconv.Itoa(int(t.Type)) + ", Value: \"" + t.Value + "\" }"
}

type Tokens []Token

func (ts Tokens) Equals(other Tokens) bool {
	if len(ts) != len(other) {
		return false
	}
	for i, token := range ts {
		if !token.Equals(other[i]) {
			return false
		}
	}
	return true
}

func (ts Tokens) String() string {
	output := "[ "
	for _, token := range ts {
		output += token.String() + ", "
	}
	output += "]"
	return output
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

var EOF = rune(0)

type Scanner struct {
	r *strings.Reader
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return EOF
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) scanWhitespace() Token {
	// Create a buffer and read the current character into it.
	var output bytes.Buffer
	output.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		ch := s.read()
		if ch == EOF {
			break
		}

		if !isWhitespace(ch) {
			s.unread()
			break
		}
		output.WriteRune(ch)
	}

	return Token{WhiteSpace, output.String()}
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == EOF {
			break
		} else if !isLetter(ch) && !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

func (s *Scanner) Scan() *Token {
	ch := s.read()

	switch ch {
	case EOF:
		return nil
	case '{':
		return &Token{LeftBrace, "{"}
	case '}':
		return &Token{RightBrace, "}"}
	case '[':
		return &Token{LeftBracket, "["}
	case ']':
		return &Token{RightBracket, "]"}
	case ':':
		return &Token{Colon, ":"}
	case ',':
		return &Token{Comma, ","}
	}

	s.unread()
	if isWhitespace(ch) {
		value := s.scanWhitespace()
		return &value
	}
	else if isLetter(ch) {
		return s.scanIdent()
	}

	return nil
}

func (s *Scanner) Analyze() Tokens {
	var output Tokens
	token := s.Scan()
	for token != nil {
		output = append(output, *token)
		token = s.Scan()
	}
	return output
}

func LexicallyAnalyze(input string) Tokens {
	reader := strings.NewReader(input)
	scanner := Scanner{r: reader}
	return scanner.Analyze()
}
