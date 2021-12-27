package base

import "scad/foundation/utils"

type IEntityStore interface {
	AddEntity(entity IEntity)
	RemoveEntityById(entityID uint64)
	GetEntityByID(entityID uint64) IEntity
	GetAllEntities() []IEntity
	GetEntitiesInArea(area utils.Area) []IEntity
	IterateAllEntities(callback func(entity IEntity))
	IterateEntitiesInArea(area utils.Area, callback func(entity IEntity))
}
