package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

var position *ecs.Component
var renderer *ecs.Component


func InitializeWorld()(*ecs.Manager, map[string]ecs.Tag)  {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()


	//register components
	player := manager.NewComponent()
	position = manager.NewComponent()
	renderer = manager.NewComponent()


	playerImg, _, err := ebitenutil.NewImageFromFile("assets/bunny1_stand.png")
	if err != nil {
		log.Fatal(err)
	}



	manager.NewEntity().
		AddComponent(player, &Player{}).
		AddComponent(renderer, &Renderer{
			Image: playerImg,
		}).
		AddComponent(position, &Position{
			X:0,
			Y:0,
		})

	players:= ecs.BuildTag(player, position)
	tags["players"] = players
	renderers := ecs.BuildTag(renderer, position)
	tags["renderers"] = renderers
	// The manager is actually the world
	return manager, tags
}
