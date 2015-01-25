package zumbies

import (
	"github.com/highway900/GarageEngine/engine"
	//"log"
	//"github.com/highway900/GarageEngine/engine/input"
)

type Player struct {
	engine.BaseComponent
	Map *Map
}

func NewPlayer() *Player {
	return &Player{BaseComponent: engine.NewComponent()}
}

func (p *Player) Start() {
	if p.Map == nil {
		p.Map = Layers[0]
	}
}
