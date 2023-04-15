package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Environemnt int

const (
	_ Environemnt = iota

	Cavern
	Forest
	Island
	Hell
)

type Difficulty int

const (
	_ Difficulty = iota

	easiest
	easy
	medium
	hard
	hardest
	boss
)

type Level struct {
	Tiles []MapTile
	Rooms []Rect

	Difficulty
	Environemnt
}

type MapTile struct {
	PixelX, PixelY int
	Blocked        bool
	Image          *ebiten.Image
}

// NewLevel creates a new game level with its tilemap
func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	return l
}

func (l *Level) CreateTiles() []MapTile {
	tiles := make([]MapTile, ScreenWidth*ScreenHeight)
	index := 0

	for x := 0; x < ScreenWidth; x++ {
		for y := 0; y < ScreenHeight; y++ {
			index = GetIndexFromXY(x, y)
			pX, pY := x*TileWidth, y*TileHeight
			wall, _, _ := ebitenutil.NewImageFromFile("assets/wall.png")
			tiles[index] = MapTile{
				PixelX:  pX,
				PixelY:  pY,
				Blocked: true,
				Image:   wall,
			}
		}
	}

	return tiles
}

func (l *Level) Draw(screen *ebiten.Image) {
	for x := 0; x < ScreenWidth; x++ {
		for y := 0; y < ScreenHeight; y++ {
			tile := l.Tiles[GetIndexFromXY(x, y)]
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, opts)
		}
	}
}

func (l *Level) GenerateLevelTiles() {
	minSize := 6
	maxSize := 12
	maxRooms := 40

	l.Tiles = l.CreateTiles()
	for i := 0; i < maxRooms; i++ {
		w := GetRandomBetween(minSize, maxSize)
		h := GetRandomBetween(minSize, maxSize)
		x := GetRandomInt(ScreenWidth - w - 1)
		y := GetRandomInt(ScreenHeight - h - 1)

		newRoom := NewRect(x, y, w, h)
		ok := true
		for _, otherRoom := range l.Rooms {
			if newRoom.Intersect(otherRoom) {
				ok = false
				break
			}
		}
		if ok {
			l.createRoom(newRoom)
			if len(l.Rooms) != 0 {
				newX, newY := newRoom.Center()
				prevX, prevY := l.Rooms[len(l.Rooms)-1].Center()

				coinflip := RollDice(2)
				if coinflip == 1 {
					l.createHorizontalTunnel(prevX, newX, prevY)
					l.createVerticalTunnel(prevY, newY, newX)
				} else {
					l.createHorizontalTunnel(prevX, newX, newY)
					l.createVerticalTunnel(prevY, newY, prevX)
				}
			}

			l.Rooms = append(l.Rooms, newRoom)
		}
	}
}

func (l *Level) createRoom(room Rect) {
	for y := room.Y1; y < room.Y2; y++ {
		for x := room.X1; x < room.X2; x++ {
			index := GetIndexFromXY(x, y)
			l.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			l.Tiles[index].Image = floor
		}
	}
}

func (l *Level) createHorizontalTunnel(x1, x2, y int) {
	for x := min(x1, x2); x <= max(x1, x2); x++ {
		index := GetIndexFromXY(x, y)
		if index > 0 && index < ScreenWidth*ScreenHeight {
			l.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			l.Tiles[index].Image = floor
		}
	}
}

func (l *Level) createVerticalTunnel(y1, y2, x int) {
	for y := min(y1, y2); y <= max(y1, y2); y++ {
		index := GetIndexFromXY(x, y)
		if index > 0 && index < ScreenWidth*ScreenHeight {
			l.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			l.Tiles[index].Image = floor
		}
	}
}

// GetIndexFromXY gets the index of the tile array for the given x,y coordinates
// Coordinates are logical tiles, not pixel!
func GetIndexFromXY(x, y int) int {
	return (y * ScreenWidth) + x
}

func pick(s []string) string {
	return s[rand.Intn(len(s))]
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
