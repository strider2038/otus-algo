package textsearch

func Parse(text string) []Keyword {
	p := parser{}

	return p.Parse([]rune(text))
}
