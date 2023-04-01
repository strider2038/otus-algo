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
	}
	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			got := textsearch.Parse(test.text)

			datatesting.AssertEqualArrays(t, test.wantKeywords, got)
		})
	}
}
