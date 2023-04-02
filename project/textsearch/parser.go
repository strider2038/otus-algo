package textsearch

import "unicode"

type parser struct{}

func (p *parser) Parse(text []rune) []Keyword {
	keywords := make([]Keyword, 0)

	sms := []*stateMachine{
		newStateMachine(standardCodePattern),
		newStateMachine(versionCodePattern),
		newStateMachine(accuracyClassPattern),
		newStateMachine(typeCodePattern),
		newStateMachine(naturalWordPattern),
		newStateMachine(genericCodePattern),
	}

	for offset := 0; offset < len(text); {
		for ; unicode.IsSpace(text[offset]); offset++ {
		}

		for _, sm := range sms {
			sm.Reset()
			n := 0
			i := offset
			for ; i < len(text); i++ {
				if sm.Handle(unicode.ToLower(text[i])) {
					n++
				} else {
					break
				}
			}
			if i == len(text) {
				sm.Finish()
			}
			if sm.IsFinished() {
				keywords = append(keywords, sm.Get()...)
				offset += n

				break
			}
		}
	}

	return keywords
}
