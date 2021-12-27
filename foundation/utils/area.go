package utils

type Area struct {
	minPt Coordinate
	maxPt Coordinate
}

func AreaFromValues(x1, y1, x2, y2 float64) Area {
	return Area{NewCoord(x1, y1), NewCoord(x2, y2)}
}

func AreaFromCoords(pt1, pt2 Coordinate) Area {
	return Area{pt1, pt2}
}

func (a *Area) Width() float64 {
	return a.maxPt.x - a.minPt.x
}

func (a *Area) Height() float64 {
	return a.maxPt.y - a.minPt.y
}

func (a *Area) ContainsCoordinate(c Coordinate) bool {
	return (c.x >= a.minPt.x && c.x <= a.maxPt.x) && (c.y >= a.minPt.y && c.y <= a.maxPt.y)
}

func (a *Area) ContainsArea(c Area) bool {
	return a.ContainsCoordinate(c.minPt) && a.ContainsCoordinate(c.maxPt)
}

func (a *Area) Intersects(b Area) bool {
	return a.ContainsCoordinate(b.minPt) || a.ContainsCoordinate(b.maxPt) || b.ContainsCoordinate(a.minPt) || b.ContainsCoordinate(a.maxPt)
}

func (a *Area) GetSplitAreas() []Area {
	center := a.Center()
	return []Area{
		AreaFromCoords(a.minPt, center),
		AreaFromValues(center.x, a.minPt.y, a.maxPt.x, center.y),
		AreaFromValues(a.minPt.x, center.y, center.x, a.maxPt.y),
		AreaFromCoords(center, a.maxPt),
	}
}

func (a *Area) Center() Coordinate {
	pt := a.minPt
	return NewCoord(pt.x+a.Width()/2, pt.y+a.Height()/2)
}

func (a *Area) MinPoint() Coordinate {
	return a.minPt
}

func (a *Area) MaxPoint() Coordinate {
	return a.maxPt
}
