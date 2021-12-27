package meta

import "scad/foundation/utils"

type LayerBuilder struct {
	layer Layer
}

func NewLayerBuilder(name string) *LayerBuilder {
	return &LayerBuilder{layer: Layer{name: name}}
}

func (lb *LayerBuilder) WithColor(color utils.Color) *LayerBuilder {
	lb.layer.color = color
	return lb
}

func (lb *LayerBuilder) WithLineType(lt LineType) *LayerBuilder {
	lb.layer.lineType = lt
	return lb
}

func (lb *LayerBuilder) WithLineWidth(lw LineWidth) *LayerBuilder {
	lb.layer.lineWidth = lw
	return lb
}

func (lb *LayerBuilder) Build() Layer {
	return lb.layer
}
