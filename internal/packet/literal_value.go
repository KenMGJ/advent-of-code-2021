package packet

import "fmt"

type LiteralValue struct {
	idType  int
	version int
	value   int
}

func NewLiteralValue(idType, value, version int) *LiteralValue {
	return &LiteralValue{
		idType:  idType,
		version: version,
		value:   value,
	}
}

func (l *LiteralValue) IdType() int {
	return l.idType
}

func (l *LiteralValue) Value() int {
	return l.value
}

func (l *LiteralValue) Version() int {
	return l.version
}

func (l *LiteralValue) VersionSum() int {
	return l.version
}

func (l *LiteralValue) String() string {
	return fmt.Sprintf("L{%d %d %d}", l.idType, l.value, l.version)
}
