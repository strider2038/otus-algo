package v2_preprocessing

import (
	"unicode"

	"github.com/strider2038/otus-algo/project/codeparsing/code"
)

// matcher - интерфейс для различных способов проверки символа.
type matcher interface {
	Matches(c rune) bool
}

// pattern - конфигурация для конечного автомата с картой переходов.
// Используется только для удобного конфигурирования конечного автомата.
type pattern struct {
	nodes       map[string][]patternTransition
	keywordType code.KeywordType
}

// patternTransition - параметры перехода в следующее состояние.
type patternTransition struct {
	condition    matcher
	target       string
	modifyResult func(result *result)

	isCharIgnored bool // символ не добавляется в результат при этом переходе
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
	// оптимизация: простое сравнение с пробелом за счет предобработки
	return c == ' '
}

type notSpace struct{}

func (d notSpace) Matches(c rune) bool {
	return c > 0 && c != ' '
}

type letter struct{}

func (a letter) Matches(c rune) bool {
	return unicode.IsLetter(c)
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

type variationCode struct{}

func (v variationCode) Matches(c rune) bool {
	if unicode.IsDigit(c) || c >= 'a' && c <= 'z' {
		return true
	}
	switch c {
	case 'а', 'б', 'в', 'н', ',', '.':
		return true
	}
	return false
}

func setStandardType(t code.StandardType) func(r *result) {
	return func(r *result) {
		r.subType = t
	}
}
