package parsing

type Pattern string

func (p Pattern) Left(x int) Pattern {
	return p[0:x]
}

func (p Pattern) Right(x int) Pattern {
	return p[len(p)-x:]
}

type Alphabet []int8

func NewAlphabet(alphabet string) []int8 {
	if len(alphabet) == 0 {
		panic("empty alphabet")
	}
	if len(alphabet) > 128 {
		panic("too big alphabet")
	}

	// размер индексной таблицы - наибольший символ из алфавита
	maxIndex := 0
	for _, char := range alphabet {
		if int(char) > maxIndex {
			maxIndex = int(char)
		}
	}

	indices := make([]int8, maxIndex+1)

	// заполняем таблицу специальными значениями
	for i := range indices {
		indices[i] = -1
	}
	// заполняем таблицу порядковыми номерами
	for i, char := range alphabet {
		indices[char] = int8(i)
	}

	return indices
}

func (alphabet Alphabet) Index(c rune) int8 {
	index := alphabet[c]
	if index < 0 {
		panic("index out of alphabet")
	}

	return index
}

type Parser struct {
	alphabet Alphabet
	pattern  Pattern
	delta    [][]int
}

func NewParser(alphabet, pattern string) *Parser {
	parser := &Parser{alphabet: NewAlphabet(alphabet), pattern: Pattern(pattern)}

	parser.delta = make([][]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		parser.delta[i] = make([]int, len(alphabet))
	}

	for q := range pattern {
		for _, c := range alphabet {
			line := parser.pattern.Left(q) + Pattern(c)
			k := q + 1
			for parser.pattern.Left(k) != line.Right(k) {
				k--
			}
			parser.delta[q][parser.alphabet.Index(c)] = k
		}
	}

	return parser
}

func (parser *Parser) Find(text string) int {
	q := 0

	for i, c := range text {
		q = parser.delta[q][parser.alphabet.Index(c)]
		if q == len(parser.pattern) {
			return i - len(parser.pattern) + 1
		}
	}

	return -1
}
