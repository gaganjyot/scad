package base

import "scad/foundation/utils"

type IEntity interface {
	Move(point utils.Coordinate) IEntity
	Rotate(angle float64) IEntity
	Scale(factor float64) IEntity

	ID() uint64
	NearestPointOnPath(ref utils.Coordinate) utils.Coordinate
	NearestPointOnEntity(ref utils.Coordinate) utils.Coordinate

	LayerName() string
	BoundingBox() utils.Area
	Color() utils.Color
}
