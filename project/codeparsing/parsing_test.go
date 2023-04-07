package codeparsing_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/project/codeparsing/code"
	"github.com/strider2038/otus-algo/project/codeparsing/v4_regexp"
)

const boltName = ` 
	Болт с шестигранной уменьшенной головкой 
	и направляющим подголовком
	М10х100.36 
	исп. 1 
	тип А 
	класса точности Б 
	ГОСТ 15590-70 
`

type TestCase struct {
	text         string
	wantKeywords []code.Keyword
}

var basicCases = []TestCase{
	{text: "", wantKeywords: []code.Keyword{}},
	{text: " ", wantKeywords: []code.Keyword{}},
	{text: "\t \n", wantKeywords: []code.Keyword{}},
}

var standardsCases = []TestCase{
	{
		text: "гост",
		wantKeywords: []code.Keyword{
			{Value: "гост", Type: code.NaturalWord},
		},
	},
	{
		text: "ГОСТ 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ 1234 ",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ 1234 ГОСТ 4321",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
			{Value: "4321", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ 12-34-56",
		wantKeywords: []code.Keyword{
			{Value: "12-34-56", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ 12.34.56",
		wantKeywords: []code.Keyword{
			{Value: "12.34.56", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ 1234.",
		wantKeywords: []code.Keyword{
			{Value: "1234.", Type: code.StandardCode, StandardType: code.GOST},
			// {Value: "гост 1234", Type: code.StandardCode}, todo: post processing case
		},
	},
	{
		text: "ГОСТ 1234. ",
		wantKeywords: []code.Keyword{
			{Value: "1234.", Type: code.StandardCode, StandardType: code.GOST},
			// {Value: "гост 1234", Type: code.StandardCode}, todo: post processing case
		},
	},
	{
		text: "ГОСТ Р 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ ИСО 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ ISO 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ Р ИСО 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ Р ISO 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ Р ISO 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "DIN EN 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.DIN},
		},
	},
	{
		text: "DIN 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.DIN},
		},
	},
	{
		text: "ТУ У 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.TU},
		},
	},
	{
		text: "ТУ 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.TU},
		},
	},
	{
		text: "СТО 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.STO},
		},
	},
	{
		text: "ОСТ 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.OST},
		},
	},
	{
		text: "СТ ЦКБА 1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.ST_CKBA},
		},
	},
}

var versionCodeCases = []TestCase{
	{
		text: "Исполнение 1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "Исполнения 1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "Исполнение1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "Исп 1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "Исп. 1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "ИСП 1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "исп1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "исп.1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "исп z",
		wantKeywords: []code.Keyword{
			{Value: "z", Type: code.VersionCode},
		},
	},
	{
		text: " исп аааа ",
		wantKeywords: []code.Keyword{
			{Value: "исп", Type: code.NaturalWord},
			{Value: "аааа", Type: code.NaturalWord},
		},
	},
	{
		text: "исп н",
		wantKeywords: []code.Keyword{
			{Value: "н", Type: code.VersionCode},
		},
	},
	{
		text: "исп.Б",
		wantKeywords: []code.Keyword{
			{Value: "б", Type: code.VersionCode},
		},
	},
	{
		text: "исп. В",
		wantKeywords: []code.Keyword{
			{Value: "в", Type: code.VersionCode},
		},
	},
	{
		text: "исп.1 кольцо",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
			{Value: "кольцо", Type: code.NaturalWord},
		},
	},
	{
		text: "исп гайки",
		wantKeywords: []code.Keyword{
			{Value: "исп", Type: code.NaturalWord},
			{Value: "гайки", Type: code.NaturalWord},
		},
	},
	{
		text: "Исполнение",
		wantKeywords: []code.Keyword{
			{Value: "исполнение", Type: code.NaturalWord},
		},
	},
}

var accuracyClassCodeCases = []TestCase{
	{
		text: "Класс Точности A",
		wantKeywords: []code.Keyword{
			{Value: "a", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "Класс Точности Б",
		wantKeywords: []code.Keyword{
			{Value: "б", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "классов точности B", // latin
		wantKeywords: []code.Keyword{
			{Value: "b", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "класса точности B", // latin
		wantKeywords: []code.Keyword{
			{Value: "b", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "классом точности B", // latin
		wantKeywords: []code.Keyword{
			{Value: "b", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "классом точности 0,5",
		wantKeywords: []code.Keyword{
			{Value: "0,5", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "классом точности 0.5",
		wantKeywords: []code.Keyword{
			{Value: "0.5", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "классом точности 0.5",
		wantKeywords: []code.Keyword{
			{Value: "0.5", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "класс точности",
		wantKeywords: []code.Keyword{
			{Value: "класс", Type: code.NaturalWord},
			{Value: "точности", Type: code.NaturalWord},
		},
	},
	{
		text: "класс точности аааа",
		wantKeywords: []code.Keyword{
			{Value: "класс", Type: code.NaturalWord},
			{Value: "точности", Type: code.NaturalWord},
			{Value: "аааа", Type: code.NaturalWord},
		},
	},
	{
		text: "класс",
		wantKeywords: []code.Keyword{
			{Value: "класс", Type: code.NaturalWord},
		},
	},
}

var typeCodeCases = []TestCase{
	{
		text: "Тип 2",
		wantKeywords: []code.Keyword{
			{Value: "2", Type: code.TypeCode},
		},
	},
	{
		text: "ТИПА B",
		wantKeywords: []code.Keyword{
			{Value: "b", Type: code.TypeCode},
		},
	},
	{
		text: "тип u",
		wantKeywords: []code.Keyword{
			{Value: "u", Type: code.TypeCode},
		},
	},
	{
		text: "тип Abc",
		wantKeywords: []code.Keyword{
			{Value: "abc", Type: code.TypeCode},
		},
	},
	{
		text: "тип ваб",
		wantKeywords: []code.Keyword{
			{Value: "ваб", Type: code.TypeCode},
		},
	},
	{
		text: "тип аааа",
		wantKeywords: []code.Keyword{
			{Value: "тип", Type: code.NaturalWord},
			{Value: "аааа", Type: code.NaturalWord},
		},
	},
	{
		text: "винт тип Н", // cyrillic
		wantKeywords: []code.Keyword{
			{Value: "винт", Type: code.NaturalWord},
			{Value: "н", Type: code.TypeCode},
		},
	},
	{
		text: "винт тип H", // latin
		wantKeywords: []code.Keyword{
			{Value: "винт", Type: code.NaturalWord},
			{Value: "h", Type: code.TypeCode},
		},
	},
	{
		text: "тип подшипника",
		wantKeywords: []code.Keyword{
			{Value: "тип", Type: code.NaturalWord},
			{Value: "подшипника", Type: code.NaturalWord},
		},
	},
}

var spaceCases = []TestCase{
	{
		text: "ГОСТ\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: " ГОСТ \t 1234 ",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ\t\t\tР\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ\t\t\tИСО\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ\t\t\tISO\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ\t\t\tР\t\t\tИСО\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ\t\t\tР\t\t\tISO\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "ГОСТ\t\t\tР\t\t\tISO\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.GOST_ISO},
		},
	},
	{
		text: "DIN\t\t\tEN\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.DIN},
		},
	},
	{
		text: "DIN\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.DIN},
		},
	},
	{
		text: "ТУ\t\t\tУ\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.TU},
		},
	},
	{
		text: "ТУ\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.TU},
		},
	},
	{
		text: "СТО\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.STO},
		},
	},
	{
		text: "ОСТ\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.OST},
		},
	},
	{
		text: "СТ\t\t\tЦКБА\t\t\t1234",
		wantKeywords: []code.Keyword{
			{Value: "1234", Type: code.StandardCode, StandardType: code.ST_CKBA},
		},
	},
	{
		text: "Исполнение\t\t\t1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "Исп\t\t\t1",
		wantKeywords: []code.Keyword{
			{Value: "1", Type: code.VersionCode},
		},
	},
	{
		text: "Класс\t\t\tТочности\t\t\tA",
		wantKeywords: []code.Keyword{
			{Value: "a", Type: code.AccuracyClassCode},
		},
	},
	{
		text: "Тип\t\t\t2",
		wantKeywords: []code.Keyword{
			{Value: "2", Type: code.TypeCode},
		},
	},
	{
		text: "\t\t\tслово\t\t\t",
		wantKeywords: []code.Keyword{
			{Value: "слово", Type: code.NaturalWord},
		},
	},
	{
		text: "\t\t\t123\t\t\t",
		wantKeywords: []code.Keyword{
			{Value: "123", Type: code.GenericCode},
		},
	},
}

var variationSuffixCases = createVariationSuffixCases()

func createVariationSuffixCases() []TestCase {
	suffixes := []string{
		"1",
		"12",
		"123",
		"1234",
		"aбв",
		"1aбв",
		"aбв1",
		"а1абв",
		"аб1абв",
		"абв1абв",
	}

	variants := []struct {
		prefix      string
		keywordType code.KeywordType
	}{
		{prefix: "исполнение", keywordType: code.VersionCode},
		{prefix: "тип", keywordType: code.TypeCode},
		{prefix: "класс точности", keywordType: code.AccuracyClassCode},
	}

	out := make([]TestCase, 0, len(suffixes)*3)

	for _, variant := range variants {
		for _, suffix := range suffixes {
			out = append(out, TestCase{
				text: " " + variant.prefix + " " + suffix + " ",
				wantKeywords: []code.Keyword{
					{Value: suffix, Type: variant.keywordType},
				},
			})
		}
	}

	return out
}

var naturalWordCases = []TestCase{
	{
		text: "слово",
		wantKeywords: []code.Keyword{
			{Value: "слово", Type: code.NaturalWord},
		},
	},
	{
		text: "какое-то",
		wantKeywords: []code.Keyword{
			{Value: "какое-то", Type: code.NaturalWord},
		},
	},
	{
		text: "какое'то",
		wantKeywords: []code.Keyword{
			{Value: "какое'то", Type: code.NaturalWord},
		},
	},
	{
		text: "какое`то",
		wantKeywords: []code.Keyword{
			{Value: "какое`то", Type: code.NaturalWord},
		},
	},
	{
		text: "какое′то",
		wantKeywords: []code.Keyword{
			{Value: "какое′то", Type: code.NaturalWord},
		},
	},
	{
		text: "-слово",
		wantKeywords: []code.Keyword{
			{Value: "-слово", Type: code.GenericCode},
			// todo: additional word
		},
	},
	{
		text: "какое′то",
		wantKeywords: []code.Keyword{
			{Value: "какое′то", Type: code.NaturalWord},
		},
	},
}

var combinationCases = []TestCase{
	{
		text: "Болт ГОСТ 1234-56",
		wantKeywords: []code.Keyword{
			{Value: "болт", Type: code.NaturalWord},
			{Value: "1234-56", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: "ГОСТ 1234-56 БОЛТ",
		wantKeywords: []code.Keyword{
			{Value: "1234-56", Type: code.StandardCode, StandardType: code.GOST},
			{Value: "болт", Type: code.NaturalWord},
		},
	},
	{
		text: "ГОСТ 123болт",
		wantKeywords: []code.Keyword{
			{Value: "гост", Type: code.NaturalWord},
			{Value: "123болт", Type: code.GenericCode},
		},
	},
	{
		text: "Электровоз ЭД4М",
		wantKeywords: []code.Keyword{
			{Value: "электровоз", Type: code.NaturalWord},
			{Value: "эд4м", Type: code.GenericCode},
		},
	},
}

var realCases = []TestCase{
	{
		text: "Подшипник роликовый тип 102000 исп.1 ГОСТ 8328-75",
		wantKeywords: []code.Keyword{
			{Value: "подшипник", Type: code.NaturalWord},
			{Value: "роликовый", Type: code.NaturalWord},
			{Value: "102000", Type: code.TypeCode},
			{Value: "1", Type: code.VersionCode},
			{Value: "8328-75", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
	{
		text: boltName,
		wantKeywords: []code.Keyword{
			{Value: "болт", Type: code.NaturalWord},
			{Value: "с", Type: code.NaturalWord},
			{Value: "шестигранной", Type: code.NaturalWord},
			{Value: "уменьшенной", Type: code.NaturalWord},
			{Value: "головкой", Type: code.NaturalWord},
			{Value: "и", Type: code.NaturalWord},
			{Value: "направляющим", Type: code.NaturalWord},
			{Value: "подголовком", Type: code.NaturalWord},
			{Value: "м10х100.36", Type: code.GenericCode},
			{Value: "1", Type: code.VersionCode},
			{Value: "а", Type: code.TypeCode},
			{Value: "б", Type: code.AccuracyClassCode},
			{Value: "15590-70", Type: code.StandardCode, StandardType: code.GOST},
		},
	},
}

type ParsingFunctionCase struct {
	name  string
	parse func(text string) []code.Keyword
}

var allCases = mergeCases(
	basicCases,
	standardsCases,
	versionCodeCases,
	accuracyClassCodeCases,
	typeCodeCases,
	variationSuffixCases,
	spaceCases,
	naturalWordCases,
	combinationCases,
	realCases,
)

var parsingCases = []ParsingFunctionCase{
	// {name: "base", parse: v1_base.Parse},
	// {name: "preprocessing", parse: v2_preprocessing.Parse},
	// {name: "complex fsm", parse: v3_complex_fsm.Parse},
	{name: "regexp", parse: v4_regexp.Parse},
}

func TestParse(t *testing.T) {
	for _, test := range allCases {
		for _, parsingCase := range parsingCases {
			t.Run(parsingCase.name+": "+test.text, func(t *testing.T) {
				got := parsingCase.parse(test.text)

				datatesting.AssertEqualArrays(t, test.wantKeywords, got)
			})
		}
	}
}

func mergeCases(cases ...[]TestCase) []TestCase {
	all := make([]TestCase, 0, len(cases))

	for _, testCases := range cases {
		all = append(all, testCases...)
	}

	return all
}
