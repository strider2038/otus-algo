package kmp

type Pattern string

func (p Pattern) Left(x int) Pattern {
	return p[0:x]
}

func (p Pattern) Right(x int) Pattern {
	return p[len(p)-x:]
}

type Parser struct {
	pattern Pattern
	pi      []int
}

// NewSlowParser создает текстовый парсер на основа алгоритма Кнутта-Мориса-Пратта,
// медленная версия со сравнением строк.
func NewSlowParser(pattern string) *Parser {
	parser := &Parser{pattern: Pattern(pattern)}

	parser.pi = make([]int, len(pattern)+1)

	for q := 0; q <= len(pattern); q++ {
		line := parser.pattern.Left(q)
		for n := 0; n < q; n++ {
			if line.Left(n) == line.Right(n) {
				parser.pi[q] = n
			}
		}
	}

	return parser
}

// NewParser создает текстовый парсер на основа алгоритма Кнутта-Мориса-Пратта.
// pattern - шаблон для поиска в строке. Быстрая версия.
func NewParser(pattern string) *Parser {
	parser := &Parser{pattern: Pattern(pattern)}

	parser.pi = make([]int, len(pattern)+1)

	// заполнение массива длин префиксов
	for q := 1; q < len(pattern); q++ {
		// n - количество совпавших символов
		n := parser.pi[q]
		for n > 0 && pattern[n] != pattern[q] {
			// будет возвращать к предыдущему префиксу
			n = parser.pi[n]
		}
		if parser.pattern[n] == parser.pattern[q] {
			n++
		}
		parser.pi[q+1] = n
	}

	return parser
}

// Find - возвращает начало вхождения шаблона подстроки pattern
// в текст text. Если вхождения не найдено, то возвращается -1.
func (parser *Parser) Find(text string) int {
	// n - количество совпавших символов
	n := 0

	for q := 0; q < len(text); q++ {
		for n > 0 && text[q] != parser.pattern[n] {
			// очередной символ не совпал с символом в шаблоне,
			// возвращаемся к предыдущему совпадающему префиксу
			n = parser.pi[n]
		}
		if text[q] == parser.pattern[n] {
			// если есть совпадение очередного символа,
			// то увеличиваем длину совпадающего фрагмента
			n++
		}
		// шаблон найден - возвращаем индекс начала вхождения шаблона
		if n == len(parser.pattern) {
			return q - len(parser.pattern) + 1
		}
	}

	return -1
}
