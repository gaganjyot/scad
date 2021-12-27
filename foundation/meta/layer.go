package meta

import "scad/foundation/utils"

type Layer struct {
	name      string
	color     utils.Color
	lineWidth LineWidth
	lineType  LineType
}

func (l *Layer) Name() string {
	return l.name
}

func (l *Layer) Color() utils.Color {
	return l.color
}

func (l *Layer) LineWidth() LineWidth {
	return l.lineWidth
}

func (l *Layer) LineType() LineType {
	return l.lineType
}
