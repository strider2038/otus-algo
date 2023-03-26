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
			text: "гост",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост", Type: textsearch.NaturalWord},
			},
		},
		{
			text: "ГОСТ 1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: " ГОСТ \t 1234 ",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ 1234 ГОСТ 4321",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост 1234", Type: textsearch.StandardCode},
				{Value: "гост 4321", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ 12-34-56",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост 12-34-56", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ 12.34.56",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост 12.34.56", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ 1234.",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост 1234.", Type: textsearch.StandardCode},
				// {Value: "гост 1234", Type: textsearch.StandardCode}, todo: processing case
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост р 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ\t\t\tИСО\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост исо 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ\t\t\tISO\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост iso 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\tИСО\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост р исо 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\tISO\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост р iso 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ГОСТ\t\t\tР\t\t\tISO\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "гост р iso 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "DIN\t\t\tEN\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "din en 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "DIN\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "din 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ТУ\t\t\tУ\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "ту у 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ТУ\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "ту 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "СТО\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "сто 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "ОСТ\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "ост 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "СТ\t\t\tЦКБА\t\t\t1234",
			wantKeywords: []textsearch.Keyword{
				{Value: "ст цкба 1234", Type: textsearch.StandardCode},
			},
		},
		{
			text: "Болт ГОСТ 1234-56",
			wantKeywords: []textsearch.Keyword{
				{Value: "болт", Type: textsearch.NaturalWord},
				{Value: "гост 1234-56", Type: textsearch.StandardCode},
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
