package base

import "scad/foundation/meta"

type IDocument interface {
	AddEntity(entity IEntity)
	RemoveEntityById(entityID uint64)
	GetAllEntities() []IEntity

	AddLayer(layer meta.Layer)
	RemoveLayer(layer meta.Layer)
	GetLayer(layerID uint64) meta.Layer
	GetAllLayers() []meta.Layer

	AddLineWidth(lineWidth meta.LineWidth)
	RemoveLineWidth(lineWidth meta.LineWidth)
	GetAllLineWidths() []meta.LineWidth

	AddLineType(lineType meta.LineType)
	RemoveLineType(lineType meta.LineType)
	GetAllLineTypes() []meta.LineType
}
