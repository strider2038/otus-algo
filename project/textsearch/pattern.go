package textsearch

import "unicode"

type matcher interface {
	Matches(c rune) bool
}

type pattern struct {
	nodes       map[string]patternNode
	keywordType KeywordType
}

type patternNode struct {
	transitions []patternTransition
	isFinal     bool
}

type patternTransition struct {
	condition     matcher
	target        string
	isCharIgnored bool // символ не добавляется в результат при этом переходе
	replacement   rune // символ для замены
}

type exact rune

func (e exact) Matches(c rune) bool {
	return rune(e) == c
}

type digit struct{}

func (d digit) Matches(c rune) bool {
	return unicode.IsDigit(c)
}

type space struct{}

func (s space) Matches(c rune) bool {
	return unicode.IsSpace(c)
}

type oneOf []rune

func (oneOf oneOf) Matches(c rune) bool {
	for _, char := range oneOf {
		if char == c {
			return true
		}
	}

	return false
}
