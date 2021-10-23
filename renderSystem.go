package main

import "github.com/hajimehoshi/ebiten/v2"


func ProcessRenderers(g *Game, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["renderers"]) {
		pos := result.Components[position].(*Position)
		img := result.Components[renderer].(*Renderer).Image

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos.X), float64(pos.Y))
		screen.DrawImage(img, op)
	}
}

