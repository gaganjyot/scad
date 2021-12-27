package meta

type LineWidth struct {
	width float64
}

func NewLineWidth(width float64) *LineWidth {
	return &LineWidth{width: width}
}
