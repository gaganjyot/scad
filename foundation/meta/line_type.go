package meta

type LineType struct {
	pattern string
}

func NewLineType(pattern string) *LineType {
	return &LineType{}
}
