package storage

import (
	"scad/foundation/utils"
	"testing"
)

func TestSum(t *testing.T) {
	tree := NewQuadTree(utils.AreaFromCoordinates(0.0, 0.0, 5000.0, 5000.0))
	if tree.entities != nil {
		t.Error("Entities should be nil for a new quadtree")
	}
}
