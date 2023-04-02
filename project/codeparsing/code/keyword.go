package code

// KeywordType - тип ключевого слова.
type KeywordType int

const (
	NaturalWord       KeywordType = iota + 1 // блок на естественном языке
	GenericCode                              // блок неспецифической части шифра
	StandardCode                             // блок обозначения стандарта, например "ГОСТ 1234"
	TypeCode                                 // блок обозначения типа изделия (например, "тип 1")
	VersionCode                              // блок обозначения исполнения изделия (например, "исп. 1", "исполнение 2")
	AccuracyClassCode                        // блок обозначения класса точности (например, "класс точности A")
)

// StandardType - тип обозначения стандарта.
type StandardType int

const (
	GOST     StandardType = iota + 1 // ГОСТ xxx, ГОСТ Р ххх
	GOST_ISO                         // ГОСТ ИСО ххх, ГОСТ ISO ххх, ГОСТ Р ИСО ххх, ГОСТ Р ISO ххх
	DIN                              // DIN xxx, DIN EN xxx
	TU                               // ТУ ххх, ТУ У ххх
	STO                              // СТО ххх
	OST                              // ОСТ ххх
	ST_CKBA                          // СТ ЦКБА ххх
)

type Keyword struct {
	Value        string
	Type         KeywordType
	StandardType StandardType
}
