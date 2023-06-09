package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Tick      int
	Map       GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	world, tags := InitWorld(g.Map.CurrentLevel)
	g.WorldTags = tags
	g.World = world
	return g
}

func (g *Game) Update() error {
	TryMovePlayer(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.Draw(screen)

	ProcessRenderables(g, level, screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return TileWidth * ScreenWidth, TileHeight * ScreenHeight
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Grawl")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
