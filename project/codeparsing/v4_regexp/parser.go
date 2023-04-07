package v4_regexp

import (
	"github.com/strider2038/otus-algo/project/codeparsing/code"
	"regexp"
)

type blockParser struct {
	keywordType code.KeywordType
	subType     code.StandardType
	index       int
	pattern     *regexp.Regexp
}

func (p *blockParser) Parse(text string) ([]code.Keyword, int) {
	indices := p.pattern.FindStringSubmatchIndex(text)
	offset := p.index * 2
	if len(indices) < offset+2 || indices[offset] < 0 {
		return nil, 0
	}

	keyword := code.Keyword{
		Value:        text[indices[offset]:indices[offset+1]],
		Type:         p.keywordType,
		StandardType: p.subType,
	}

	return []code.Keyword{keyword}, indices[offset+1]
}
