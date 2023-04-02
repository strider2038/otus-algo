package v3_complex_fsm

import (
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

	cleaned := make([]rune, 0, utf8.RuneCountInString(text)+1)
	ignoreSpace := true
	for _, c := range text {
		if unicode.IsSpace(c) {
			if ignoreSpace {
				continue
			}

			cleaned = append(cleaned, ' ')
			ignoreSpace = true
		} else {
			cleaned = append(cleaned, unicode.ToLower(c))
			ignoreSpace = false
		}
	}

	// добавление пробела в конец строки вместо терминального символа
	// за счет этого не нужно принудительно вызывать Finish()
	if len(cleaned) > 0 && cleaned[len(cleaned)-1] != ' ' {
		cleaned = append(cleaned, ' ')
	}

	return parseKeywords(cleaned)
}

// blockParsers - предварительно инициализированный набор парсеров различных типов
// блоков ключевых слов.
var blockParsers = []*stateMachine{
	newStateMachine(codePattern),
	newStateMachine(naturalWordPattern),
	newStateMachine(genericCodePattern),
}

func parseKeywords(text []rune) []code.Keyword {
	keywords := make([]code.Keyword, 0)

	for offset := 0; offset < len(text); {
		// пробелы сразу игнорируются
		for ; offset < len(text) && text[offset] == ' '; offset++ {
		}

		// поочередно пытаемся применить каждый конечный автомат к блоку начиная с offset
		for _, blockParser := range blockParsers {
			parser := blockParser.Start()
			parsedCount := 0
			i := offset
			for ; i < len(text); i++ {
				if parser.Handle(text[i]) {
					parsedCount++
				} else {
					break
				}
			}
			// если конечный автомат успешно отработал, то добавляем распознанные блоки
			// и смещаемся к следующему блоку
			if parser.IsFinished() {
				keywords = append(keywords, parser.Get()...)
				offset += parsedCount

				break
			}
		}
	}

	return keywords
}
