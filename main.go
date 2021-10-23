package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

//Game holds all data the entire game will need.
type Game struct{
	World *ecs.Manager
	WorldTags map[string]ecs.Tag
}
func NewGame() *Game{
	g := &Game{}
	world, tags := InitializeWorld()
	g.WorldTags = tags
	g.World = world
	return g
}

//Update is called each tic.
func (g *Game) Update() error{
	return nil
}

//Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image){
	ProcessRenderers(g, screen)
}

//Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int){
	gd := NewGameData()
	return gd.ScreenWidth, gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Title")
	if err := ebiten.RunGame(g); err != nil{
		log.Fatal(err)
	}
}
