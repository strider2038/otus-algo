package textsearch_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/project/textsearch"
)

func TestParse(t *testing.T) {
	tests := []struct {
		text         string
		wantKeywords []textsearch.Keyword
	}{
		{
			text:         "",
			wantKeywords: []textsearch.Keyword{},
		},
		{
			text: "гост",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "ГОСТ 1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ 1234 ",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: " ГОСТ \t 1234 ",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ 1234 ГОСТ 4321",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
				{Value: "4321", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ 12-34-56",
			wantKeywords: []textsearch.Keyword{
				{Value: "12-34-56", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ 12.34.56",
			wantKeywords: []textsearch.Keyword{
				{Value: "12.34.56", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ 1234.",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234.", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
				// {Value: "гост 1234", Type: textsearch.StandardCode}, todo: post processing case
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ\t\t\tИСО\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST_ISO},
			},
		},
		{
			text: "ГОСТ\t\t\tISO\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST_ISO},
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\tИСО\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST_ISO},
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\tISO\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST_ISO},
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\tISO\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.GOST_ISO},
			},
		},
		{
			text: "DIN\t\t\tEN\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.DIN},
			},
		},
		{
			text: "DIN\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.DIN},
			},
		},
		{
			text: "ТУ\t\t\tУ\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.TU},
			},
		},
		{
			text: "ТУ\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.TU},
			},
		},
		{
			text: "СТО\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.STO},
			},
		},
		{
			text: "ОСТ\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.OST},
			},
		},
		{
			text: "СТ\t\t\tЦКБА\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234", Type: textsearch.StandardCode, StandardType: textsearch.ST_CKBA},
			},
		},
		{
			text: "Болт ГОСТ 1234-56",
			wantKeywords: []textsearch.Keyword{
				{Value: "болт", Type: textsearch.NaturalWord},
				{Value: "1234-56", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
			},
		},
		{
			text: "ГОСТ 1234-56 БОЛТ",
			wantKeywords: []textsearch.Keyword{
				{Value: "1234-56", Type: textsearch.StandardCode, StandardType: textsearch.GOST},
				{Value: "болт", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "ГОСТ 123болт",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост", Type: textsearch.NaturalWord},
				{Value: "123болт", Type: textsearch.GenericCode},
			},
		},
		{
			text: "Электровоз ЭД4М",
			wantKeywords: []textsearch.Keyword{
				{Value: "электровоз", Type: textsearch.NaturalWord},
				{Value: "эд4м", Type: textsearch.GenericCode},
			},
		},
		{
			text: "слово",
			wantKeywords: []textsearch.Keyword{
				{Value: "слово", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "какое-то",
			wantKeywords: []textsearch.Keyword{
				{Value: "какое-то", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "какое'то",
			wantKeywords: []textsearch.Keyword{
				{Value: "какое'то", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "какое`то",
			wantKeywords: []textsearch.Keyword{
				{Value: "какое`то", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "какое′то",
			wantKeywords: []textsearch.Keyword{
				{Value: "какое′то", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "-слово",
			wantKeywords: []textsearch.Keyword{
				{Value: "-слово", Type: textsearch.GenericCode},
				// todo: additional word
			},
		},
		{
			text: "какое′то",
			wantKeywords: []textsearch.Keyword{
				{Value: "какое′то", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "Исполнение 1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "Исполнения 1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "Исполнение1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "Исп 1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "Исп. 1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "ИСП 1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп.1",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп z",
			wantKeywords: []textsearch.Keyword{
				{Value: "z", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп 12 ",
			wantKeywords: []textsearch.Keyword{
				{Value: "12", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп 123 ",
			wantKeywords: []textsearch.Keyword{
				{Value: "123", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп 1234 ",
			wantKeywords: []textsearch.Keyword{
				{Value: "исп", Type: textsearch.NaturalWord},
				{Value: "1234", Type: textsearch.GenericCode},
			},
		},
		{
			text: "исп абв",
			wantKeywords: []textsearch.Keyword{
				{Value: "абв", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп н",
			wantKeywords: []textsearch.Keyword{
				{Value: "н", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп.Б",
			wantKeywords: []textsearch.Keyword{
				{Value: "б", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп. В",
			wantKeywords: []textsearch.Keyword{
				{Value: "в", Type: textsearch.VersionCode},
			},
		},
		{
			text: "исп.1 кольцо",
			wantKeywords: []textsearch.Keyword{
				{Value: "1", Type: textsearch.VersionCode},
				{Value: "кольцо", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "исп гайки",
			wantKeywords: []textsearch.Keyword{
				{Value: "исп", Type: textsearch.NaturalWord},
				{Value: "гайки", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "Исполнение",
			wantKeywords: []textsearch.Keyword{
				{Value: "исполнение", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "Класс Точности A",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс точности a", Type: textsearch.AccuracyClassCode},
			},
		},
		{
			text: "Класс Точности Б",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс точности b", Type: textsearch.AccuracyClassCode},
			},
		},
		{
			text: "классов точности B",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс точности b", Type: textsearch.AccuracyClassCode},
			},
		},
		{
			text: "класса точности B",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс точности b", Type: textsearch.AccuracyClassCode},
			},
		},
		{
			text: "классом точности B",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс точности b", Type: textsearch.AccuracyClassCode},
			},
		},
		{
			text: "классом точности 0,5",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс точности 0,5", Type: textsearch.AccuracyClassCode},
			},
		},
		{
			text: "класс точности",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс", Type: textsearch.NaturalWord},
				{Value: "точности", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "Тип 2",
			wantKeywords: []textsearch.Keyword{
				{Value: "тип 2", Type: textsearch.TypeCode},
			},
		},
		{
			text: "ТИПА B",
			wantKeywords: []textsearch.Keyword{
				{Value: "тип b", Type: textsearch.TypeCode},
			},
		},
		{
			text: "тип u",
			wantKeywords: []textsearch.Keyword{
				{Value: "тип u", Type: textsearch.TypeCode},
			},
		},
		{
			text: "тип Abc",
			wantKeywords: []textsearch.Keyword{
				{Value: "тип abc", Type: textsearch.TypeCode},
			},
		},
		{
			text: "тип ваб",
			wantKeywords: []textsearch.Keyword{
				{Value: "тип vab", Type: textsearch.TypeCode},
			},
		},
		{
			text: "винт тип Н", // cyrillic
			wantKeywords: []textsearch.Keyword{
				{Value: "винт", Type: textsearch.NaturalWord},
				{Value: "тип h", Type: textsearch.TypeCode},
			},
		},
		{
			text: "винт тип H", // latin
			wantKeywords: []textsearch.Keyword{
				{Value: "винт", Type: textsearch.NaturalWord},
				{Value: "тип h", Type: textsearch.TypeCode},
			},
		},
		{
			text: "тип подшипника",
			wantKeywords: []textsearch.Keyword{
				{Value: "тип", Type: textsearch.NaturalWord},
				{Value: "подшипника", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "класс",
			wantKeywords: []textsearch.Keyword{
				{Value: "класс", Type: textsearch.NaturalWord},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			got := textsearch.Parse(test.text)

			datatesting.AssertEqualArrays(t, test.wantKeywords, got)
		})
	}
}
