package v1_base

import "github.com/strider2038/otus-algo/project/codeparsing/code"

// Конфигурация конечного автомата для обнаружения названий стандартов.
//
// Примеры:
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
var standardCodePattern = pattern{
	keywordType: code.StandardCode,
	nodes: map[string][]patternTransition{
		initialState: {
			{condition: exact('г'), target: "гост_г", isCharIgnored: true},
			{condition: exact('d'), target: "din_d", isCharIgnored: true},
			{condition: exact('т'), target: "ту_т", isCharIgnored: true},
			{condition: exact('с'), target: "сто_с", isCharIgnored: true},
			{condition: exact('о'), target: "ост_о", isCharIgnored: true},
		},
		"гост_г": {{condition: exact('о'), target: "гост_о", isCharIgnored: true}},
		"гост_о": {{condition: exact('с'), target: "гост_с", isCharIgnored: true}},
		"гост_с": {{condition: exact('т'), target: "гост_т", isCharIgnored: true}},
		"гост_т": {{condition: space{}, target: "гост_пробел", replacement: ' ', modifyResult: setStandardType(code.GOST), isCharIgnored: true}},
		"din_d":  {{condition: exact('i'), target: "din_i", isCharIgnored: true}},
		"din_i":  {{condition: exact('n'), target: "din_n", isCharIgnored: true}},
		"din_n":  {{condition: space{}, target: "din_пробел", replacement: ' ', modifyResult: setStandardType(code.DIN), isCharIgnored: true}},
		"ту_т":   {{condition: exact('у'), target: "ту_у", isCharIgnored: true}},
		"ту_у":   {{condition: space{}, target: "ту_пробел", replacement: ' ', modifyResult: setStandardType(code.TU), isCharIgnored: true}},
		"сто_с":  {{condition: exact('т'), target: "сто_т", isCharIgnored: true}},
		"сто_т": {
			{condition: exact('о'), target: "сто_о", isCharIgnored: true},
			{condition: space{}, target: "ст_пробел", replacement: ' ', isCharIgnored: true},
		},
		"сто_о": {{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(code.STO), isCharIgnored: true}},
		"ост_о": {{condition: exact('с'), target: "ост_с", isCharIgnored: true}},
		"ост_с": {{condition: exact('т'), target: "ост_т", isCharIgnored: true}},
		"ост_т": {{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(code.OST), isCharIgnored: true}},
		"гост_пробел": {
			{condition: digit{}, target: "гост_цифра"},
			{condition: space{}, target: "гост_пробел", isCharIgnored: true},
			{condition: exact('р'), target: "гост_р", isCharIgnored: true},
			{condition: exact('и'), target: "гост_исо_и", isCharIgnored: true},
			{condition: exact('i'), target: "гост_iso_i", isCharIgnored: true},
		},
		"din_пробел": {
			{condition: exact('e'), target: "din_en_e", isCharIgnored: true},
			{condition: digit{}, target: "гост_цифра"},
			{condition: space{}, target: "din_пробел", isCharIgnored: true},
		},
		"ту_пробел": {
			{condition: exact('у'), target: "ту_у_у", isCharIgnored: true},
			{condition: digit{}, target: "гост_цифра"},
			{condition: space{}, target: "ту_пробел", isCharIgnored: true},
		},
		"ст_пробел": {
			{condition: exact('ц'), target: "ст_цкба_ц", modifyResult: setStandardType(code.ST_CKBA), isCharIgnored: true},
			{condition: space{}, target: "ст_пробел", isCharIgnored: true},
		},
		"din_en_e": {
			{condition: exact('n'), target: "din_en_n", isCharIgnored: true},
		},
		"din_en_n": {
			{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', isCharIgnored: true},
		},
		"ту_у_у": {
			{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', isCharIgnored: true},
		},
		"гост_р": {
			{condition: space{}, target: "гост_р_пробел", replacement: ' ', isCharIgnored: true},
		},
		"гост_р_пробел": {
			{condition: digit{}, target: "гост_цифра"},
			{condition: space{}, target: "гост_р_пробел", isCharIgnored: true},
			{condition: exact('и'), target: "гост_исо_и", isCharIgnored: true},
			{condition: exact('i'), target: "гост_iso_i", isCharIgnored: true},
		},
		"гост_исо_и": {
			{condition: exact('с'), target: "гост_исо_с", isCharIgnored: true},
		},
		"гост_исо_с": {
			{condition: exact('о'), target: "гост_исо_о", isCharIgnored: true},
		},
		"гост_исо_о": {
			{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(code.GOST_ISO), isCharIgnored: true},
		},
		"гост_iso_i": {
			{condition: exact('s'), target: "гост_iso_s", isCharIgnored: true},
		},
		"гост_iso_s": {
			{condition: exact('o'), target: "гост_iso_o", isCharIgnored: true},
		},
		"гост_iso_o": {
			{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', modifyResult: setStandardType(code.GOST_ISO), isCharIgnored: true},
		},
		"ст_цкба_ц": {{condition: exact('к'), target: "ст_цкба_к", isCharIgnored: true}},
		"ст_цкба_к": {{condition: exact('б'), target: "ст_цкба_б", isCharIgnored: true}},
		"ст_цкба_б": {{condition: exact('а'), target: "ст_цкба_а", isCharIgnored: true}},
		"ст_цкба_а": {{condition: space{}, target: "гост_пробел_разделитель", replacement: ' ', isCharIgnored: true}},
		"гост_пробел_разделитель": {
			{condition: digit{}, target: "гост_цифра"},
			{condition: space{}, target: "гост_пробел_разделитель", isCharIgnored: true},
		},
		"гост_цифра": {
			{condition: digit{}, target: "гост_цифра"},
			{condition: oneOf{'.', '-'}, target: "гост_разделитель"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"гост_разделитель": {
			{condition: digit{}, target: "гост_цифра"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
	},
}

// Конфигурация конечного автомата для обнаружения названий кодов исполнения изделия.
//
// Примеры:
// "исполнение 1"
// "исполнение1"
// "исп 1"
// "исп 1"
// "исп1"
// "исп. 1"
// "исп.1"
var versionCodePattern = pattern{
	keywordType: code.VersionCode,
	nodes: map[string][]patternTransition{
		initialState: {
			{condition: exact('и'), target: "и", isCharIgnored: true},
		},
		"и": {{condition: exact('с'), target: "с", isCharIgnored: true}},
		"с": {{condition: exact('п'), target: "п", isCharIgnored: true}},
		"п": {
			{condition: exact('о'), target: "о", isCharIgnored: true},
			{condition: exact('.'), target: "точка", isCharIgnored: true},
			{condition: digit{}, target: "код"},
			{condition: space{}, target: "пробел", isCharIgnored: true},
		},
		"о":   {{condition: exact('л'), target: "л", isCharIgnored: true}},
		"л":   {{condition: exact('н'), target: "н", isCharIgnored: true}},
		"н":   {{condition: exact('е'), target: "е", isCharIgnored: true}},
		"е":   {{condition: exact('н'), target: "н_2", isCharIgnored: true}},
		"н_2": {{condition: exact('и'), target: "и_2", isCharIgnored: true}},
		"и_2": {
			{condition: exact('е'), target: "е_2", isCharIgnored: true},
			{condition: exact('я'), target: "я", isCharIgnored: true},
		},
		"е_2": {
			{condition: space{}, target: "пробел", isCharIgnored: true},
			{condition: digit{}, target: "код"},
		},
		"я": {{condition: space{}, target: "пробел", isCharIgnored: true}},
		"точка": {
			{condition: space{}, target: "пробел", isCharIgnored: true},
			{condition: variationCode{}, target: "код"},
		},
		"пробел": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_1"},
			{condition: space{}, target: "пробел", isCharIgnored: true},
		},
		"код": {
			{condition: variationCode{}, target: "код"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_1": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_2"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_2": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_3"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_3": {
			{condition: digit{}, target: "код"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
	},
}

// Конфигурация конечного автомата для обнаружения классов точности изделия.
//
// Примеры:
// "класс точности А"
// "класса точности А"
// "классов точности А"
// "классом точности А"
var accuracyClassPattern = pattern{
	keywordType: code.AccuracyClassCode,
	nodes: map[string][]patternTransition{
		initialState: {
			{condition: exact('к'), target: "к", isCharIgnored: true},
		},
		"к":   {{condition: exact('л'), target: "л", isCharIgnored: true}},
		"л":   {{condition: exact('а'), target: "а_1", isCharIgnored: true}},
		"а_1": {{condition: exact('с'), target: "с_1", isCharIgnored: true}},
		"с_1": {{condition: exact('с'), target: "с_2", isCharIgnored: true}},
		"с_2": {
			{condition: exact('а'), target: "а_2", isCharIgnored: true},
			{condition: exact('о'), target: "о_1", isCharIgnored: true},
			{condition: space{}, target: "пробел_1", isCharIgnored: true},
		},
		"а_2": {{condition: space{}, target: "пробел_1", isCharIgnored: true}},
		"о_1": {
			{condition: exact('в'), target: "в", isCharIgnored: true},
			{condition: exact('м'), target: "м", isCharIgnored: true},
		},
		"в": {{condition: space{}, target: "пробел_1", isCharIgnored: true}},
		"м": {{condition: space{}, target: "пробел_1", isCharIgnored: true}},
		"пробел_1": {
			{condition: exact('т'), target: "т_1", isCharIgnored: true},
			{condition: space{}, target: "пробел_1", isCharIgnored: true},
		},
		"т_1": {{condition: exact('о'), target: "о_2", isCharIgnored: true}},
		"о_2": {{condition: exact('ч'), target: "ч", isCharIgnored: true}},
		"ч":   {{condition: exact('н'), target: "н", isCharIgnored: true}},
		"н":   {{condition: exact('о'), target: "о_3", isCharIgnored: true}},
		"о_3": {{condition: exact('с'), target: "с_3", isCharIgnored: true}},
		"с_3": {{condition: exact('т'), target: "т_2", isCharIgnored: true}},
		"т_2": {{condition: exact('и'), target: "и", isCharIgnored: true}},
		"и":   {{condition: space{}, target: "пробел_2", isCharIgnored: true}},
		"пробел_2": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_1"},
			{condition: space{}, target: "пробел_2", isCharIgnored: true},
		},
		"код": {
			{condition: variationCode{}, target: "код"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_1": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_2"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_2": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_3"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_3": {
			{condition: digit{}, target: "код"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
	},
}

// Конфигурация конечного автомата для обнаружения типов изделия.
//
// Примеры:
// "тип А"
// "типа А"
var typeCodePattern = pattern{
	keywordType: code.TypeCode,
	nodes: map[string][]patternTransition{
		initialState: {{condition: exact('т'), target: "т", isCharIgnored: true}},
		"т":          {{condition: exact('и'), target: "и", isCharIgnored: true}},
		"и":          {{condition: exact('п'), target: "п", isCharIgnored: true}},
		"п": {
			{condition: exact('а'), target: "а", isCharIgnored: true},
			{condition: space{}, target: "пробел", isCharIgnored: true},
		},
		"а": {{condition: space{}, target: "пробел", isCharIgnored: true}},
		"пробел": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_1"},
			{condition: space{}, target: "пробел", isCharIgnored: true},
		},
		"код": {
			{condition: variationCode{}, target: "код"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_1": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_2"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_2": {
			{condition: digit{}, target: "код"},
			{condition: variationCode{}, target: "буквенный_код_3"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"буквенный_код_3": {
			{condition: digit{}, target: "код"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
	},
}

var naturalWordPattern = pattern{
	keywordType: code.NaturalWord,
	nodes: map[string][]patternTransition{
		initialState: {
			{condition: letter{}, target: "letter"},
		},
		"letter": {
			{condition: letter{}, target: "letter"},
			{condition: oneOf{'-', '\'', '`', '′'}, target: "delimiter"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
		"delimiter": {
			{condition: letter{}, target: "letter"},
		},
	},
}

var genericCodePattern = pattern{
	keywordType: code.GenericCode,
	nodes: map[string][]patternTransition{
		initialState: {
			{condition: notSpace{}, target: "any"},
		},
		"any": {
			{condition: notSpace{}, target: "any"},
			{condition: space{}, target: finalState, isCharIgnored: true},
			{condition: null{}, target: finalState},
		},
	},
}
