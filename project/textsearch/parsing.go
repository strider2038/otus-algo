package textsearch

import (
	"unicode"

	"github.com/strider2038/otus-algo/project/textsearch/code"
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
	return parseKeywords([]rune(text))
}

// blockParsers - предварительно инициализированный набор парсеров различных типов
// блоков ключевых слов.
var blockParsers = []*stateMachine{
	newStateMachine(standardCodePattern),
	newStateMachine(versionCodePattern),
	newStateMachine(accuracyClassPattern),
	newStateMachine(typeCodePattern),
	newStateMachine(naturalWordPattern),
	newStateMachine(genericCodePattern),
}

func parseKeywords(text []rune) []code.Keyword {
	keywords := make([]code.Keyword, 0)

	for offset := 0; offset < len(text); {
		// пробелы сразу игнорируются
		for ; offset < len(text) && unicode.IsSpace(text[offset]); offset++ {
		}

		// поочередно пытаемся применить каждый конечный автомат к блоку начиная с offset
		for _, blockParser := range blockParsers {
			parser := blockParser.Start()
			parsedCount := 0
			i := offset
			for ; i < len(text); i++ {
				if parser.Handle(unicode.ToLower(text[i])) {
					parsedCount++
				} else {
					break
				}
			}
			// если дошли до конца текста, то вызываем принудительное завершение
			if i == len(text) {
				parser.Finish()
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
