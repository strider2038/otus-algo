package textsearch

import "unicode"

type parser struct{}

func (p *parser) Parse(text []rune) []Keyword {
	keywords := make([]Keyword, 0)

	sm := newStateMachine(standardCodePattern, func(keyword Keyword) {
		keywords = append(keywords, keyword)
	})

	for _, char := range text {
		sm.Handle(unicode.ToLower(char))
	}

	sm.Finish()

	return keywords
}
