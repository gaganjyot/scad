package storage

import (
	"scad/foundation/base"
	"scad/foundation/utils"
)

const MAX_ENTITES_PER_NODE = 10

type QuadTree struct {
	area         utils.Area
	currentLevel int
	entities     []base.IEntity
	nodes        []QuadTree
}

func NewQuadTree(area utils.Area) QuadTree {
	return QuadTree{area: area, currentLevel: 0, nodes: nil}
}

func newQuadTreeWithLevel(area utils.Area, level int) QuadTree {
	return QuadTree{area: area, currentLevel: level, nodes: nil}
}

func (store *QuadTree) AddEntity(entity base.IEntity) bool {
	if len(store.entities) >= MAX_ENTITES_PER_NODE {
		if store.nodes == nil {
			store.performSplit()
		}
	} else {
		store.entities = append(store.entities, entity)
	}
	return true
}

func (store *QuadTree) RemoveEntityById(entityID uint64) {
	for i, entity := range store.entities {
		if entity.ID() == entityID {
			store.entities = append(store.entities[:i], store.entities[i+1:]...)
			return
		}
	}
	if store.nodes == nil {
		return
	}
	for _, node := range store.nodes {
		node.RemoveEntityById(entityID)
	}
}

func (store *QuadTree) GetEntityByID(entityID uint64) base.IEntity {
	for _, entity := range store.entities {
		if entity.ID() == entityID {
			return entity
		}
	}
	if store.nodes == nil {
		for _, node := range store.nodes {
			if node.GetEntityByID(entityID) != nil {
				return node.GetEntityByID(entityID)
			}
		}
	}
	return nil
}

func (store *QuadTree) GetAllEntities() []base.IEntity {
	ret := []base.IEntity{}
	ret = append(ret, store.entities...)
	if store.nodes != nil {
		for _, node := range store.nodes {
			ret = append(ret, node.GetAllEntities()...)
		}
	}
	return ret
}

func (store *QuadTree) GetEntitiesInArea(area utils.Area) []base.IEntity {
	// TODO
	// we can optimize this functionality that we don't get bounding box
	// for elements that are completely contained in the area
	ret := []base.IEntity{}

	if !area.ContainsArea(store.area) {
		return ret
	}

	if !area.Intersects(store.area) {
		return ret
	}

	for _, entity := range store.entities {
		if area.Intersects(entity.BoundingBox()) {
			ret = append(ret, entity)
		}
	}
	if store.nodes == nil {
		return ret
	}

	for _, node := range store.nodes {
		ret = append(ret, node.GetEntitiesInArea(area)...)
	}

	return ret
}

func (store *QuadTree) IterateAllEntities(callback func(entity base.IEntity)) {
	for _, entity := range store.entities {
		callback(entity)
	}
	if store.nodes != nil {
		for _, node := range store.nodes {
			node.IterateAllEntities(callback)
		}
	}
}

func (store *QuadTree) IterateEntitiesInArea(area utils.Area, callback func(entity base.IEntity)) {

	// TODO
	// we can optimize this functionality that we don't get bounding box
	// for elements that are completely contained in the area

	if !area.ContainsArea(store.area) {
		return
	}

	if !area.Intersects(store.area) {
		return
	}

	for _, entity := range store.entities {
		if area.Intersects(entity.BoundingBox()) {
			callback(entity)
		}
	}

	if store.nodes == nil {
		return
	}

	for _, node := range store.nodes {
		node.IterateEntitiesInArea(area, callback)
	}
}

func (store *QuadTree) performSplit() {
	store.nodes = make([]QuadTree, 4)

	areas := store.area.GetSplitAreas()

	for i := 0; i < len(areas); i++ {
		store.nodes[i] = newQuadTreeWithLevel(areas[i], store.currentLevel+1)
	}

	store.entities = []base.IEntity{}
}
