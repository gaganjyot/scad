package utils

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func FromRGBA(r, g, b, a uint8) Color {
	return Color{r: r, g: g, b: b, a: a}
}

func FromRGB(r, g, b, a uint8) Color {
	return Color{r: r, g: g, b: b, a: 0}
}

func (color *Color) R() uint8 {
	return color.r
}

func (color *Color) G() uint8 {
	return color.g
}

func (color *Color) B() uint8 {
	return color.b
}

func (color *Color) A() uint8 {
	return color.a
}
