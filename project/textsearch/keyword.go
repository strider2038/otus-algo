package textsearch

type KeywordType int

const (
	NaturalWord       KeywordType = iota + 1 // блок на естественном языке
	GenericCode                              // блок неспецифической части шифра
	StandardCode                             // блок обозначения стандарта, например "ГОСТ 1234"
	TypeCode                                 // блок обозначения типа изделия (например, "тип 1")
	VersionCode                              // блок обозначения исполнения изделия (например, "исп. 1", "исполнение 2")
	AccuracyClassCode                        // блок обозначения класса точности (например, "класс точности A")
)

type Keyword struct {
	Value string
	Type  KeywordType
}
