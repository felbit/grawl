package main

import "github.com/hajimehoshi/ebiten/v2"

func ProcessRenderables(g *Game, level Level, screen *ebiten.Image) {
	for _, r := range g.World.Query(g.WorldTags["renderables"]) {
		pos := r.Components[position].(*Position)
		img := r.Components[renderable].(*Renderable).Image

		index := GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(img, opts)
	}
}
