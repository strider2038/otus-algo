package textsearch

var standardCodePattern = pattern{
	// конечный автомат для обнаружения названий стандартов
	// примеры:
	// "ГОСТ 19532-74"
	// "ГОСТ Р 50792-95"
	// "ГОСТ Р ИСО 1580-2013"
	// "ГОСТ Р ISO 1580-2013"
	// "ГОСТ ИСО 1580-2013"
	// "ГОСТ ISO 1580-2013"
	// "DIN EN 934"
	// "DIN 934"
	// "ТУ У 1580"
	// "ТУ 1580"
	// "СТО 1580"
	// "ОСТ 1580"
	// "СТ ЦКБА 1580"
	keywordType: StandardCode,
	nodes: map[string]patternNode{
		"": {
			transitions: []patternTransition{
				{condition: exact('г'), target: "гост_г"},
				{condition: exact('d'), target: "din_d"},
				{condition: exact('т'), target: "ту_т"},
				{condition: exact('с'), target: "сто_с"},
				{condition: exact('о'), target: "ост_о"},
			},
		},
		"гост_г": {transitions: []patternTransition{{condition: exact('о'), target: "гост_о"}}},
		"гост_о": {transitions: []patternTransition{{condition: exact('с'), target: "гост_с"}}},
		"гост_с": {transitions: []patternTransition{{condition: exact('т'), target: "гост_т"}}},
		"гост_т": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел", replacement: ' '}}},
		"din_d":  {transitions: []patternTransition{{condition: exact('i'), target: "din_i"}}},
		"din_i":  {transitions: []patternTransition{{condition: exact('n'), target: "din_n"}}},
		"din_n":  {transitions: []patternTransition{{condition: space{}, target: "din_пробел", replacement: ' '}}},
		"ту_т":   {transitions: []patternTransition{{condition: exact('у'), target: "ту_у"}}},
		"ту_у":   {transitions: []patternTransition{{condition: space{}, target: "ту_пробел", replacement: ' '}}},
		"сто_с":  {transitions: []patternTransition{{condition: exact('т'), target: "сто_т"}}},
		"сто_т": {
			transitions: []patternTransition{
				{condition: exact('о'), target: "сто_о"},
				{condition: space{}, target: "ст_пробел", replacement: ' '},
			},
		},
		"сто_о": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел_разделитель", replacement: ' '}}},
		"ост_о": {transitions: []patternTransition{{condition: exact('с'), target: "ост_с"}}},
		"ост_с": {transitions: []patternTransition{{condition: exact('т'), target: "ост_т"}}},
		"ост_т": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел_разделитель", replacement: ' '}}},
		"гост_пробел": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "гост_пробел", isCharIgnored: true},
				{condition: exact('р'), target: "гост_р"},
				{condition: exact('и'), target: "гост_исо_и"},
				{condition: exact('i'), target: "гост_iso_i"},
			},
		},
		"din_пробел": {
			transitions: []patternTransition{
				{condition: exact('e'), target: "din_en_e"},
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "din_пробел", isCharIgnored: true},
			},
		},
		"ту_пробел": {
			transitions: []patternTransition{
				{condition: exact('у'), target: "ту_у_у"},
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "ту_пробел", isCharIgnored: true},
			},
		},
		"ст_пробел": {
			transitions: []patternTransition{
				{condition: exact('ц'), target: "ст_цкба_ц"},
				{condition: space{}, target: "ст_пробел", isCharIgnored: true},
			},
		},
		"din_en_e": {
			transitions: []patternTransition{
				{condition: exact('n'), target: "din_en_n"},
			},
		},
		"din_en_n": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' '},
			},
		},
		"ту_у_у": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' '},
			},
		},
		"гост_р": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_р_пробел", replacement: ' '},
			},
		},
		"гост_р_пробел": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "гост_р_пробел", isCharIgnored: true},
				{condition: exact('и'), target: "гост_исо_и"},
				{condition: exact('i'), target: "гост_iso_i"},
			},
		},
		"гост_исо_и": {
			transitions: []patternTransition{
				{condition: exact('с'), target: "гост_исо_с"},
			},
		},
		"гост_исо_с": {
			transitions: []patternTransition{
				{condition: exact('о'), target: "гост_исо_о"},
			},
		},
		"гост_исо_о": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' '},
			},
		},
		"гост_iso_i": {
			transitions: []patternTransition{
				{condition: exact('s'), target: "гост_iso_s"},
			},
		},
		"гост_iso_s": {
			transitions: []patternTransition{
				{condition: exact('o'), target: "гост_iso_o"},
			},
		},
		"гост_iso_o": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' '},
			},
		},
		"ст_цкба_ц": {transitions: []patternTransition{{condition: exact('к'), target: "ст_цкба_к"}}},
		"ст_цкба_к": {transitions: []patternTransition{{condition: exact('б'), target: "ст_цкба_б"}}},
		"ст_цкба_б": {transitions: []patternTransition{{condition: exact('а'), target: "ст_цкба_а"}}},
		"ст_цкба_а": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел_разделитель", replacement: ' '}}},
		"гост_пробел_разделитель": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "гост_пробел_разделитель", isCharIgnored: true},
			},
		},
		"гост_цифра": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
				{condition: oneOf{'.', '-'}, target: "гост_разделитель"},
			},
			isFinal: true,
		},
		"гост_разделитель": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
			},
			isFinal: true,
		},
	},
}
