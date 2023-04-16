package v4_regexp

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/strider2038/otus-algo/project/codeparsing/code"
)

// Parse - разбирает название изделия на смысловые блоки, подходящие для
// организации его поиска в системе.
//
// Например, из названия "Подшипник роликовый тип 102000 исп.1 ГОСТ 8328-75" будут выделены
// блоки:
//   - слово NaturalWord "подшипник"
//   - слово NaturalWord "роликовый"
//   - код типа TypeCode "102000"
//   - код исполнения VersionCode "1"
//   - код стандарта StandardCode с типом GOST "8328-75"
func Parse(text string) []code.Keyword {
	// оптимизация с предобработкой строки:
	// 1) переводим в нижний регистр;
	// 2) удаляем лишние пробелы;
	// 3) добавляем пробел в конец вместо терминального символа.

	cleaned := strings.Builder{}
	cleaned.Grow(utf8.RuneCountInString(text) + 1)
	isSpace := false
	ignoreSpace := true
	for _, c := range text {
		isSpace = unicode.IsSpace(c)
		if isSpace {
			if ignoreSpace {
				continue
			}

			cleaned.WriteRune(' ')
			ignoreSpace = true
		} else {
			cleaned.WriteRune(unicode.ToLower(c))
			ignoreSpace = false
		}
	}

	// добавление пробела в конец строки вместо терминального символа
	if !isSpace {
		cleaned.WriteRune(' ')
	}

	return parseKeywords(cleaned.String())
}

func parseKeywords(text string) []code.Keyword {
	keywords := make([]code.Keyword, 0)

	for offset := 0; offset < len(text); {
		// пробелы сразу игнорируются
		for ; offset < len(text) && text[offset] == ' '; offset++ {
		}

		// поочередно пытаемся применить каждый парсер к блоку начиная с offset
		for _, parser := range blockParsers {
			if kw, parsedCount := parser.Parse(text[offset:]); parsedCount > 0 {
				keywords = append(keywords, kw...)
				offset += parsedCount

				break
			}
		}
	}

	return keywords
}

var blockParsers = []*blockParser{
	{
		keywordType: code.StandardCode,
		subType:     code.GOST,
		pattern:     regexp.MustCompile(`^гост (р )?(\d+[-.\d+]*) `),
		index:       2,
	},
	{
		keywordType: code.StandardCode,
		subType:     code.GOST_ISO,
		pattern:     regexp.MustCompile(`^гост (р )?(iso|исо) (\d+[-.\d+]*) `),
		index:       3,
	},
	{
		keywordType: code.StandardCode,
		subType:     code.DIN,
		pattern:     regexp.MustCompile(`^din (en )?(\d+[-.\d+]*) `),
		index:       2,
	},
	{
		keywordType: code.StandardCode,
		subType:     code.TU,
		pattern:     regexp.MustCompile(`^ту (у )?(\d+[-.\d+]*) `),
		index:       2,
	},
	{
		keywordType: code.StandardCode,
		subType:     code.STO,
		pattern:     regexp.MustCompile(`^сто (\d+[-.\d+]*) `),
		index:       1,
	},
	{
		keywordType: code.StandardCode,
		subType:     code.OST,
		pattern:     regexp.MustCompile(`^ост (\d+[-.\d+]*) `),
		index:       1,
	},
	{
		keywordType: code.StandardCode,
		subType:     code.ST_CKBA,
		pattern:     regexp.MustCompile(`^ст цкба (\d+[-.\d+]*) `),
		index:       1,
	},
	{
		keywordType: code.VersionCode,
		pattern:     regexp.MustCompile(`^(исп( |. ?)|исполнен(ия|ие) )(([а-внa-z]{1,3})|([а-внa-z]{0,3}\d+\S*)) `),
		index:       4,
	},
	{
		keywordType: code.VersionCode,
		pattern:     regexp.MustCompile(`^(исп|исполнен(ия|ие))(\d+\S*) `),
		index:       3,
	},
	{
		keywordType: code.AccuracyClassCode,
		pattern:     regexp.MustCompile(`^класс(а|ов|ом)? точности (([а-внa-z]{1,3})|([а-внa-z]{0,3}\d+\S*)) `),
		index:       2,
	},
	{
		keywordType: code.TypeCode,
		pattern:     regexp.MustCompile(`^типа? (([а-внa-z]{1,3})|([а-внa-z]{0,3}\d+\S*)) `),
		index:       1,
	},
	{
		keywordType: code.NaturalWord,
		pattern:     regexp.MustCompile("^([a-zа-я]+[-'`′a-zа-я]*) "),
		index:       1,
	},
	{
		keywordType: code.GenericCode,
		pattern:     regexp.MustCompile(`(\S+) `),
		index:       1,
	},
}
