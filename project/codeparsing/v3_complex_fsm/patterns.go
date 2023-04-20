package v3_complex_fsm

import "github.com/strider2038/otus-algo/project/codeparsing/code"

var codePattern = pattern{
	initialState: {
		{condition: exact('г'), target: "гост_г", isCharIgnored: true},
		{condition: exact('d'), target: "din_d", isCharIgnored: true},
		{condition: exact('т'), target: "ту_т", isCharIgnored: true},
		{condition: exact('с'), target: "сто_с", isCharIgnored: true},
		{condition: exact('о'), target: "ост_о", isCharIgnored: true},
		{condition: exact('и'), target: "исполнение_и", isCharIgnored: true},
		{condition: exact('к'), target: "класс_точности_к", isCharIgnored: true},
	},

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
	"гост_г": {{condition: exact('о'), target: "гост_о", isCharIgnored: true}},
	"гост_о": {{condition: exact('с'), target: "гост_с", isCharIgnored: true}},
	"гост_с": {{condition: exact('т'), target: "гост_т", isCharIgnored: true}},
	"гост_т": {{condition: space{}, target: "гост_пробел", modifyResult: setStandardType(code.GOST), isCharIgnored: true}},
	"din_d":  {{condition: exact('i'), target: "din_i", isCharIgnored: true}},
	"din_i":  {{condition: exact('n'), target: "din_n", isCharIgnored: true}},
	"din_n":  {{condition: space{}, target: "din_пробел", modifyResult: setStandardType(code.DIN), isCharIgnored: true}},
	"ту_т": {
		{condition: exact('у'), target: "ту_у", isCharIgnored: true},
		{condition: exact('и'), target: "тип_и", isCharIgnored: true},
	},
	"ту_у":  {{condition: space{}, target: "ту_пробел", modifyResult: setStandardType(code.TU), isCharIgnored: true}},
	"сто_с": {{condition: exact('т'), target: "сто_т", isCharIgnored: true}},
	"сто_т": {
		{condition: exact('о'), target: "сто_о", isCharIgnored: true},
		{condition: space{}, target: "ст_пробел", isCharIgnored: true},
	},
	"сто_о": {{condition: space{}, target: "гост_пробел_разделитель", modifyResult: setStandardType(code.STO), isCharIgnored: true}},
	"ост_о": {{condition: exact('с'), target: "ост_с", isCharIgnored: true}},
	"ост_с": {{condition: exact('т'), target: "ост_т", isCharIgnored: true}},
	"ост_т": {{condition: space{}, target: "гост_пробел_разделитель", modifyResult: setStandardType(code.OST), isCharIgnored: true}},
	"гост_пробел": {
		{condition: digit{}, target: "гост_цифра", start: true},
		{condition: space{}, target: "гост_пробел", isCharIgnored: true},
		{condition: exact('р'), target: "гост_р", isCharIgnored: true},
		{condition: exact('и'), target: "гост_исо_и", isCharIgnored: true},
		{condition: exact('i'), target: "гост_iso_i", isCharIgnored: true},
	},
	"din_пробел": {
		{condition: exact('e'), target: "din_en_e", isCharIgnored: true},
		{condition: digit{}, target: "гост_цифра", start: true},
		{condition: space{}, target: "din_пробел", isCharIgnored: true},
	},
	"ту_пробел": {
		{condition: exact('у'), target: "ту_у_у", isCharIgnored: true},
		{condition: digit{}, target: "гост_цифра", start: true},
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
		{condition: space{}, target: "гост_пробел_разделитель", isCharIgnored: true},
	},
	"ту_у_у": {
		{condition: space{}, target: "гост_пробел_разделитель", isCharIgnored: true},
	},
	"гост_р": {
		{condition: space{}, target: "гост_р_пробел", isCharIgnored: true},
	},
	"гост_р_пробел": {
		{condition: digit{}, target: "гост_цифра", start: true},
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
		{condition: space{}, target: "гост_пробел_разделитель", modifyResult: setStandardType(code.GOST_ISO), isCharIgnored: true},
	},
	"гост_iso_i": {
		{condition: exact('s'), target: "гост_iso_s", isCharIgnored: true},
	},
	"гост_iso_s": {
		{condition: exact('o'), target: "гост_iso_o", isCharIgnored: true},
	},
	"гост_iso_o": {
		{condition: space{}, target: "гост_пробел_разделитель", modifyResult: setStandardType(code.GOST_ISO), isCharIgnored: true},
	},
	"ст_цкба_ц": {{condition: exact('к'), target: "ст_цкба_к", isCharIgnored: true}},
	"ст_цкба_к": {{condition: exact('б'), target: "ст_цкба_б", isCharIgnored: true}},
	"ст_цкба_б": {{condition: exact('а'), target: "ст_цкба_а", isCharIgnored: true}},
	"ст_цкба_а": {{condition: space{}, target: "гост_пробел_разделитель", isCharIgnored: true}},
	"гост_пробел_разделитель": {
		{condition: digit{}, target: "гост_цифра", start: true},
		{condition: space{}, target: "гост_пробел_разделитель", isCharIgnored: true},
	},
	"гост_цифра": {
		{condition: digit{}, target: "гост_цифра"},
		{condition: oneOf{'.', '-'}, target: "гост_разделитель"},
		{condition: space{}, target: finalState, isCharIgnored: true},
	},
	"гост_разделитель": {
		{condition: digit{}, target: "гост_цифра"},
		{condition: space{}, target: finalState, isCharIgnored: true},
	},

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
	"исполнение_и": {{condition: exact('с'), target: "исполнение_с", isCharIgnored: true, modifyResult: setCodeType(code.VersionCode)}},
	"исполнение_с": {{condition: exact('п'), target: "исполнение_п", isCharIgnored: true}},
	"исполнение_п": {
		{condition: exact('о'), target: "исполнение_о", isCharIgnored: true},
		{condition: exact('.'), target: "исполнение_точка", isCharIgnored: true},
		{condition: digit{}, target: "вариация_код", start: true},
		{condition: space{}, target: "исполнение_пробел", isCharIgnored: true},
	},
	"исполнение_о":   {{condition: exact('л'), target: "исполнение_л", isCharIgnored: true}},
	"исполнение_л":   {{condition: exact('н'), target: "исполнение_н", isCharIgnored: true}},
	"исполнение_н":   {{condition: exact('е'), target: "исполнение_е", isCharIgnored: true}},
	"исполнение_е":   {{condition: exact('н'), target: "исполнение_н_2", isCharIgnored: true}},
	"исполнение_н_2": {{condition: exact('и'), target: "исполнение_и_2", isCharIgnored: true}},
	"исполнение_и_2": {
		{condition: exact('е'), target: "исполнение_е_2", isCharIgnored: true},
		{condition: exact('я'), target: "исполнение_я", isCharIgnored: true},
	},
	"исполнение_е_2": {
		{condition: space{}, target: "исполнение_пробел", isCharIgnored: true},
		{condition: digit{}, target: "вариация_код", start: true},
	},
	"исполнение_я": {{condition: space{}, target: "исполнение_пробел", isCharIgnored: true}},
	"исполнение_точка": {
		{condition: space{}, target: "исполнение_пробел", isCharIgnored: true},
		{condition: variationCode{}, target: "вариация_код", start: true},
	},
	"исполнение_пробел": {
		{condition: digit{}, target: "вариация_код", start: true},
		{condition: variationCode{}, target: "вариация_буквенный_код_1", start: true},
		{condition: space{}, target: "исполнение_пробел", isCharIgnored: true},
	},

	// Конфигурация конечного автомата для обнаружения классов точности изделия.
	//
	// Примеры:
	// "класс точности А"
	// "класса точности А"
	// "классов точности А"
	// "классом точности А"
	"класс_точности_к":   {{condition: exact('л'), target: "класс_точности_л", isCharIgnored: true, modifyResult: setCodeType(code.AccuracyClassCode)}},
	"класс_точности_л":   {{condition: exact('а'), target: "класс_точности_а_1", isCharIgnored: true}},
	"класс_точности_а_1": {{condition: exact('с'), target: "класс_точности_с_1", isCharIgnored: true}},
	"класс_точности_с_1": {{condition: exact('с'), target: "класс_точности_с_2", isCharIgnored: true}},
	"класс_точности_с_2": {
		{condition: exact('а'), target: "класс_точности_а_2", isCharIgnored: true},
		{condition: exact('о'), target: "класс_точности_о_1", isCharIgnored: true},
		{condition: space{}, target: "класс_точности_пробел_1", isCharIgnored: true},
	},
	"класс_точности_а_2": {{condition: space{}, target: "класс_точности_пробел_1", isCharIgnored: true}},
	"класс_точности_о_1": {
		{condition: exact('в'), target: "класс_точности_в", isCharIgnored: true},
		{condition: exact('м'), target: "класс_точности_м", isCharIgnored: true},
	},
	"класс_точности_в": {{condition: space{}, target: "класс_точности_пробел_1", isCharIgnored: true}},
	"класс_точности_м": {{condition: space{}, target: "класс_точности_пробел_1", isCharIgnored: true}},
	"класс_точности_пробел_1": {
		{condition: exact('т'), target: "класс_точности_т_1", isCharIgnored: true},
		{condition: space{}, target: "класс_точности_пробел_1", isCharIgnored: true},
	},
	"класс_точности_т_1": {{condition: exact('о'), target: "класс_точности_о_2", isCharIgnored: true}},
	"класс_точности_о_2": {{condition: exact('ч'), target: "класс_точности_ч", isCharIgnored: true}},
	"класс_точности_ч":   {{condition: exact('н'), target: "класс_точности_н", isCharIgnored: true}},
	"класс_точности_н":   {{condition: exact('о'), target: "класс_точности_о_3", isCharIgnored: true}},
	"класс_точности_о_3": {{condition: exact('с'), target: "класс_точности_с_3", isCharIgnored: true}},
	"класс_точности_с_3": {{condition: exact('т'), target: "класс_точности_т_2", isCharIgnored: true}},
	"класс_точности_т_2": {{condition: exact('и'), target: "класс_точности_и", isCharIgnored: true}},
	"класс_точности_и":   {{condition: space{}, target: "класс_точности_пробел_2", isCharIgnored: true}},
	"класс_точности_пробел_2": {
		{condition: digit{}, target: "вариация_код", start: true},
		{condition: variationCode{}, target: "вариация_буквенный_код_1", start: true},
		{condition: space{}, target: "класс_точности_пробел_2", isCharIgnored: true},
	},

	// Конфигурация конечного автомата для обнаружения типов изделия.
	//
	// Примеры:
	// "тип А"
	// "типа А"
	"тип_и": {{condition: exact('п'), target: "тип_п", isCharIgnored: true, modifyResult: setCodeType(code.TypeCode)}},
	"тип_п": {
		{condition: exact('а'), target: "тип_а", isCharIgnored: true},
		{condition: space{}, target: "тип_пробел", isCharIgnored: true},
	},
	"тип_а": {{condition: space{}, target: "тип_пробел", isCharIgnored: true}},
	"тип_пробел": {
		{condition: digit{}, target: "вариация_код", start: true},
		{condition: variationCode{}, target: "вариация_буквенный_код_1", start: true},
		{condition: space{}, target: "тип_пробел", isCharIgnored: true},
	},

	// общая часть для вариаций (исполнение, тип, класс точности)
	"вариация_код": {
		{condition: variationCode{}, target: "вариация_код"},
		{condition: space{}, target: finalState, isCharIgnored: true},
	},
	"вариация_буквенный_код_1": {
		{condition: digit{}, target: "вариация_код"},
		{condition: variationCode{}, target: "вариация_буквенный_код_2"},
		{condition: space{}, target: finalState, isCharIgnored: true},
	},
	"вариация_буквенный_код_2": {
		{condition: digit{}, target: "вариация_код"},
		{condition: variationCode{}, target: "вариация_буквенный_код_3"},
		{condition: space{}, target: finalState, isCharIgnored: true},
	},
	"вариация_буквенный_код_3": {
		{condition: digit{}, target: "вариация_код"},
		{condition: space{}, target: finalState, isCharIgnored: true},
	},
}

var genericPattern = pattern{
	initialState: {
		{condition: letter{}, target: "letter", modifyResult: setCodeType(code.NaturalWord)},
		{condition: always{}, target: "generic_code", modifyResult: setCodeType(code.GenericCode)},
	},
	"letter": {
		{condition: letter{}, target: "letter"},
		{condition: oneOf{'-', '\'', '`', '′'}, target: "delimiter"},
		{condition: space{}, target: finalState, isCharIgnored: true},
		{condition: always{}, target: "generic_code", modifyResult: setCodeType(code.GenericCode)},
	},
	"delimiter": {
		{condition: letter{}, target: "letter"},
		{condition: always{}, target: "generic_code", modifyResult: setCodeType(code.GenericCode)},
	},
	"generic_code": {
		{condition: space{}, target: finalState, isCharIgnored: true},
		{condition: always{}, target: "generic_code", modifyResult: setCodeType(code.GenericCode)},
	},
}
