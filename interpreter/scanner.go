package interpreter

import (
	"glox"
	"strconv"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"func":   FUNC,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	errorReporter glox.ErrorReporter
	source        string
	tokens        []Token
	start         int
	current       int
	line          int
}

func NewScanner(source string, er glox.ErrorReporter) *Scanner {
	return &Scanner{
		errorReporter: er,
		source:        source,
		tokens:        make([]Token, 0),
		start:         0,
		current:       0,
		line:          1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, *NewToken(EOF, "", nil, s.line))
	return s.tokens
}

func (s *Scanner) scanToken() {
	char := s.advance()
	switch char {
	case '(':
		s.addNonLiteralToken(LEFT_PAREN)
		break
	case ')':
		s.addNonLiteralToken(RIGHT_PAREN)
		break
	case '{':
		s.addNonLiteralToken(LEFT_BRACE)
		break
	case '}':
		s.addNonLiteralToken(RIGHT_BRACE)
		break
	case ',':
		s.addNonLiteralToken(COMMA)
		break
	case '.':
		s.addNonLiteralToken(DOT)
		break
	case '-':
		s.addNonLiteralToken(MINUS)
		break
	case '+':
		s.addNonLiteralToken(PLUS)
		break
	case ';':
		s.addNonLiteralToken(SEMICOLON)
		break
	case '*':
		s.addNonLiteralToken(STAR)
		break
	case '!':
		if s.match('=') {
			s.addNonLiteralToken(BANG_EQUAL)
		} else {
			s.addNonLiteralToken(BANG)
		}
		break
	case '=':
		if s.match('=') {
			s.addNonLiteralToken(EQUAL_EQUAL)
		} else {
			s.addNonLiteralToken(EQUAL)
		}
		break
	case '<':
		if s.match('=') {
			s.addNonLiteralToken(LESS_EQUAL)
		} else {
			s.addNonLiteralToken(LESS)
		}
		break
	case '>':
		if s.match('=') {
			s.addNonLiteralToken(GREATER_EQUAL)
		} else {
			s.addNonLiteralToken(GREATER)
		}
		break
	case '/':
		if s.match('/') {
			for s.peek() != nil && *s.peek() != '\n' {
				s.advance()
			}
		} else {
			s.addNonLiteralToken(SLASH)
		}
		break
	case ' ', '\r', '\t':
		break
	case '\n':
		s.line++
		break
	case '"':
		s.str()
		break
	default:
		if isDigit(char) {
			s.number()
		} else if isAlpha(char) {
			s.identifier()
		} else {
			s.errorReporter.Process(s.line, "Unexpected character.")
		}
	}
}

func (s *Scanner) advance() byte {
	char := s.source[s.current]
	s.current++
	return char
}

func (s *Scanner) addNonLiteralToken(tokenType TokenType) {
	s.addLiteralToken(tokenType, nil)
}

func (s *Scanner) addLiteralToken(tokenType TokenType, literal interface{}) {
	lexeme := s.source[s.start:s.current]
	s.tokens = append(s.tokens, *NewToken(tokenType, lexeme, literal, s.line))
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() || s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() *byte {
	if s.isAtEnd() {
		return nil
	}
	char := s.source[s.current]
	return &char
}

func (s *Scanner) peekNext() *byte {
	if s.current+1 >= len(s.source) {
		return nil
	}
	char := s.source[s.current+1]
	return &char
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) str() {
	for s.peek() != nil && *s.peek() != '"' {
		if *s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.errorReporter.Process(s.line, "Unterminated string.")
		return
	}

	s.advance()
	s.addLiteralToken(STRING, s.source[s.start+1:s.current-1])
}

func (s *Scanner) number() {
	for s.peek() != nil && isDigit(*s.peek()) {
		s.advance()
	}

	if s.peek() != nil && *s.peek() == '.' && s.peekNext() != nil && isDigit(*s.peekNext()) {
		s.advance()
		for s.peek() != nil && isDigit(*s.peek()) {
			s.advance()
		}
	}

	extractedNumber, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		panic(err)
	}
	s.addLiteralToken(NUMBER, extractedNumber)
}

func (s *Scanner) identifier() {
	for s.peek() != nil && isAlphaNumeric(*s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType, isKeyword := keywords[text]
	if !isKeyword {
		tokenType = IDENTIFIER
	}
	s.addNonLiteralToken(tokenType)
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		char == '_'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlphaNumeric(char byte) bool {
	return isAlpha(char) || isDigit(char)
}
