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
		initialState: {
			transitions: []patternTransition{
				{condition: exact('г'), target: "гост_г", isCharIgnored: true},
				{condition: exact('d'), target: "din_d", isCharIgnored: true},
				{condition: exact('т'), target: "ту_т", isCharIgnored: true},
				{condition: exact('с'), target: "сто_с", isCharIgnored: true},
				{condition: exact('о'), target: "ост_о", isCharIgnored: true},
			},
		},
		"гост_г": {transitions: []patternTransition{{condition: exact('о'), target: "гост_о", isCharIgnored: true}}},
		"гост_о": {transitions: []patternTransition{{condition: exact('с'), target: "гост_с", isCharIgnored: true}}},
		"гост_с": {transitions: []patternTransition{{condition: exact('т'), target: "гост_т", isCharIgnored: true}}},
		"гост_т": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел", replacement: ' ', modifyResult: setStandardType(GOST), isCharIgnored: true}}},
		"din_d":  {transitions: []patternTransition{{condition: exact('i'), target: "din_i", isCharIgnored: true}}},
		"din_i":  {transitions: []patternTransition{{condition: exact('n'), target: "din_n", isCharIgnored: true}}},
		"din_n":  {transitions: []patternTransition{{condition: space{}, target: "din_пробел", replacement: ' ', modifyResult: setStandardType(DIN), isCharIgnored: true}}},
		"ту_т":   {transitions: []patternTransition{{condition: exact('у'), target: "ту_у", isCharIgnored: true}}},
		"ту_у":   {transitions: []patternTransition{{condition: space{}, target: "ту_пробел", replacement: ' ', modifyResult: setStandardType(TU), isCharIgnored: true}}},
		"сто_с":  {transitions: []patternTransition{{condition: exact('т'), target: "сто_т", isCharIgnored: true}}},
		"сто_т": {
			transitions: []patternTransition{
				{condition: exact('о'), target: "сто_о", isCharIgnored: true},
				{condition: space{}, target: "ст_пробел", replacement: ' ', isCharIgnored: true},
			},
		},
		"сто_о": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(STO), isCharIgnored: true}}},
		"ост_о": {transitions: []patternTransition{{condition: exact('с'), target: "ост_с", isCharIgnored: true}}},
		"ост_с": {transitions: []patternTransition{{condition: exact('т'), target: "ост_т", isCharIgnored: true}}},
		"ост_т": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(OST), isCharIgnored: true}}},
		"гост_пробел": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "гост_пробел", isCharIgnored: true},
				{condition: exact('р'), target: "гост_р", isCharIgnored: true},
				{condition: exact('и'), target: "гост_исо_и", isCharIgnored: true},
				{condition: exact('i'), target: "гост_iso_i", isCharIgnored: true},
			},
		},
		"din_пробел": {
			transitions: []patternTransition{
				{condition: exact('e'), target: "din_en_e", isCharIgnored: true},
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "din_пробел", isCharIgnored: true},
			},
		},
		"ту_пробел": {
			transitions: []patternTransition{
				{condition: exact('у'), target: "ту_у_у", isCharIgnored: true},
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "ту_пробел", isCharIgnored: true},
			},
		},
		"ст_пробел": {
			transitions: []patternTransition{
				{condition: exact('ц'), target: "ст_цкба_ц", modifyResult: setStandardType(ST_CKBA), isCharIgnored: true},
				{condition: space{}, target: "ст_пробел", isCharIgnored: true},
			},
		},
		"din_en_e": {
			transitions: []patternTransition{
				{condition: exact('n'), target: "din_en_n", isCharIgnored: true},
			},
		},
		"din_en_n": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', isCharIgnored: true},
			},
		},
		"ту_у_у": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', isCharIgnored: true},
			},
		},
		"гост_р": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_р_пробел", replacement: ' ', isCharIgnored: true},
			},
		},
		"гост_р_пробел": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: "гост_р_пробел", isCharIgnored: true},
				{condition: exact('и'), target: "гост_исо_и", isCharIgnored: true},
				{condition: exact('i'), target: "гост_iso_i", isCharIgnored: true},
			},
		},
		"гост_исо_и": {
			transitions: []patternTransition{
				{condition: exact('с'), target: "гост_исо_с", isCharIgnored: true},
			},
		},
		"гост_исо_с": {
			transitions: []patternTransition{
				{condition: exact('о'), target: "гост_исо_о", isCharIgnored: true},
			},
		},
		"гост_исо_о": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(GOST_ISO), isCharIgnored: true},
			},
		},
		"гост_iso_i": {
			transitions: []patternTransition{
				{condition: exact('s'), target: "гост_iso_s", isCharIgnored: true},
			},
		},
		"гост_iso_s": {
			transitions: []patternTransition{
				{condition: exact('o'), target: "гост_iso_o", isCharIgnored: true},
			},
		},
		"гост_iso_o": {
			transitions: []patternTransition{
				{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(GOST_ISO), isCharIgnored: true},
			},
		},
		"ст_цкба_ц": {transitions: []patternTransition{{condition: exact('к'), target: "ст_цкба_к", isCharIgnored: true}}},
		"ст_цкба_к": {transitions: []patternTransition{{condition: exact('б'), target: "ст_цкба_б", isCharIgnored: true}}},
		"ст_цкба_б": {transitions: []patternTransition{{condition: exact('а'), target: "ст_цкба_а", isCharIgnored: true}}},
		"ст_цкба_а": {transitions: []patternTransition{{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', isCharIgnored: true}}},
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
				{condition: space{}, target: finalState, isCharIgnored: true},
				{condition: null{}, target: finalState},
			},
		},
		"гост_разделитель": {
			transitions: []patternTransition{
				{condition: digit{}, target: "гост_цифра"},
				{condition: space{}, target: finalState, isCharIgnored: true}, // todo: test char
				{condition: null{}, target: finalState},
			},
		},
	},
}

var versionCodePattern = pattern{
	keywordType: VersionCode,
	nodes: map[string]patternNode{
		initialState: {
			transitions: []patternTransition{
				{condition: exact('и'), target: "и", isCharIgnored: true},
			},
		},
		"и": {transitions: []patternTransition{{condition: exact('с'), target: "с", isCharIgnored: true}}},
		"с": {transitions: []patternTransition{{condition: exact('п'), target: "п", isCharIgnored: true}}},
		"п": {transitions: []patternTransition{
			{condition: exact('о'), target: "о", isCharIgnored: true},
			{condition: exact('.'), target: "точка", isCharIgnored: true},
			{condition: digit{}, target: "код_1"},
			{condition: space{}, target: "пробел", isCharIgnored: true},
		}},
		"о":   {transitions: []patternTransition{{condition: exact('л'), target: "л", isCharIgnored: true}}},
		"л":   {transitions: []patternTransition{{condition: exact('н'), target: "н", isCharIgnored: true}}},
		"н":   {transitions: []patternTransition{{condition: exact('е'), target: "е", isCharIgnored: true}}},
		"е":   {transitions: []patternTransition{{condition: exact('н'), target: "н_2", isCharIgnored: true}}},
		"н_2": {transitions: []patternTransition{{condition: exact('и'), target: "и_2", isCharIgnored: true}}},
		"и_2": {transitions: []patternTransition{
			{condition: exact('е'), target: "е_2", isCharIgnored: true},
			{condition: exact('я'), target: "я", isCharIgnored: true},
		}},
		"е_2": {transitions: []patternTransition{
			{condition: space{}, target: "пробел", isCharIgnored: true},
			{condition: digit{}, target: "код_1"},
		}},
		"я": {transitions: []patternTransition{{condition: space{}, target: "пробел", isCharIgnored: true}}},
		"точка": {transitions: []patternTransition{
			{condition: space{}, target: "пробел", isCharIgnored: true},
			{condition: variationCode{}, target: "код_1"},
		}},
		// todo: пропуски пробелов
		"пробел": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_1"},
		}},
		"код_1": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_2"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
		"код_2": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_3"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
		"код_3": {transitions: []patternTransition{
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
	},
}

var accuracyClassPattern = pattern{
	keywordType: AccuracyClassCode,
	nodes: map[string]patternNode{
		initialState: {transitions: []patternTransition{
			{condition: exact('к'), target: "к", isCharIgnored: true},
		}},
		"к":   {transitions: []patternTransition{{condition: exact('л'), target: "л", isCharIgnored: true}}},
		"л":   {transitions: []patternTransition{{condition: exact('а'), target: "а_1", isCharIgnored: true}}},
		"а_1": {transitions: []patternTransition{{condition: exact('с'), target: "с_1", isCharIgnored: true}}},
		"с_1": {transitions: []patternTransition{{condition: exact('с'), target: "с_2", isCharIgnored: true}}},
		"с_2": {transitions: []patternTransition{
			{condition: exact('а'), target: "а_2", isCharIgnored: true},
			{condition: exact('о'), target: "о_1", isCharIgnored: true},
			{condition: space{}, target: "пробел_1", isCharIgnored: true},
		}},
		"а_2": {transitions: []patternTransition{{condition: space{}, target: "пробел_1", isCharIgnored: true}}},
		"о_1": {transitions: []patternTransition{
			{condition: exact('в'), target: "в", isCharIgnored: true},
			{condition: exact('м'), target: "м", isCharIgnored: true},
		}},
		"в":        {transitions: []patternTransition{{condition: space{}, target: "пробел_1", isCharIgnored: true}}},
		"м":        {transitions: []patternTransition{{condition: space{}, target: "пробел_1", isCharIgnored: true}}},
		"пробел_1": {transitions: []patternTransition{{condition: exact('т'), target: "т_1", isCharIgnored: true}}},
		"т_1":      {transitions: []patternTransition{{condition: exact('о'), target: "о_2", isCharIgnored: true}}},
		"о_2":      {transitions: []patternTransition{{condition: exact('ч'), target: "ч", isCharIgnored: true}}},
		"ч":        {transitions: []patternTransition{{condition: exact('н'), target: "н", isCharIgnored: true}}},
		"н":        {transitions: []patternTransition{{condition: exact('о'), target: "о_3", isCharIgnored: true}}},
		"о_3":      {transitions: []patternTransition{{condition: exact('с'), target: "с_3", isCharIgnored: true}}},
		"с_3":      {transitions: []patternTransition{{condition: exact('т'), target: "т_2", isCharIgnored: true}}},
		"т_2":      {transitions: []patternTransition{{condition: exact('и'), target: "и", isCharIgnored: true}}},
		"и":        {transitions: []patternTransition{{condition: space{}, target: "пробел_2", isCharIgnored: true}}},
		// todo: пропуски пробелов
		"пробел_2": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_1"},
		}},
		"код_1": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_2"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
		"код_2": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_3"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
		"код_3": {transitions: []patternTransition{
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
	},
}

var typeCodePattern = pattern{
	keywordType: TypeCode,
	nodes: map[string]patternNode{
		initialState: {transitions: []patternTransition{{condition: exact('т'), target: "т", isCharIgnored: true}}},
		"т":          {transitions: []patternTransition{{condition: exact('и'), target: "и", isCharIgnored: true}}},
		"и":          {transitions: []patternTransition{{condition: exact('п'), target: "п", isCharIgnored: true}}},
		"п": {transitions: []patternTransition{
			{condition: exact('а'), target: "а", isCharIgnored: true},
			{condition: space{}, target: "пробел", isCharIgnored: true},
		}},
		"а": {transitions: []patternTransition{{condition: space{}, target: "пробел", isCharIgnored: true}}},
		// todo: пропуски пробелов
		"пробел": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_1"},
		}},
		"код_1": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_2"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
		"код_2": {transitions: []patternTransition{
			{condition: variationCode{}, target: "код_3"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
		"код_3": {transitions: []patternTransition{
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		}},
	},
}

var naturalWordPattern = pattern{
	keywordType: NaturalWord,
	nodes: map[string]patternNode{
		initialState: {
			transitions: []patternTransition{
				{condition: letter{}, target: "letter"},
			},
		},
		"letter": {
			transitions: []patternTransition{
				{condition: letter{}, target: "letter"},
				{condition: oneOf{'-', '\'', '`', '′'}, target: "delimiter"},
				{condition: space{}, target: finalState, isCharIgnored: true}, // todo: test char
				{condition: null{}, target: finalState},
			},
		},
		"delimiter": {
			transitions: []patternTransition{
				{condition: letter{}, target: "letter"},
			},
		},
	},
}

var genericCodePattern = pattern{
	keywordType: GenericCode,
	nodes: map[string]patternNode{
		initialState: {
			transitions: []patternTransition{
				{condition: notSpace{}, target: "any"},
			},
		},
		"any": {
			transitions: []patternTransition{
				{condition: notSpace{}, target: "any"},
				{condition: space{}, target: finalState, isCharIgnored: true}, // todo: test char
				{condition: null{}, target: finalState},
			},
		},
	},
}
