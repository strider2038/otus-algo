package codeparsing_test

import (
	"encoding/json"
	"os"
	"testing"
)

func BenchmarkParse_FullCase(b *testing.B) {
	for _, parsingCase := range parsingCases {
		b.Run(parsingCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = parsingCase.parse(boltName)
			}
		})
	}
}

func BenchmarkParse_TestCases(b *testing.B) {
	for _, parsingCase := range parsingCases {
		b.Run(parsingCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, test := range allCases {
					_ = parsingCase.parse(test.text)
				}
			}
		})
	}
}

func BenchmarkParse_RealCases(b *testing.B) {
	names := loadNames(b)

	for _, parsingCase := range parsingCases {
		b.Run(parsingCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, name := range names {
					_ = parsingCase.parse(name)
				}
			}
		})
	}
}

func loadNames(tb testing.TB) []string {
	data, err := os.ReadFile("./testdata/product_names.json")
	if err != nil {
		tb.Fatal(`read "product_names.json":`, err)
	}

	var items []struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &items); err != nil {
		tb.Fatal(`unmarshal "product_names.json":`, err)
	}

	names := make([]string, 0, len(items))
	for _, item := range items {
		names = append(names, item.Name)
	}

	tb.Log("read", len(names), "product names")

	return names
}
