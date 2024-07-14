package ntt

import (
	"bufio"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TileWidth  = 69
	TileHeight = 69
)

type Tile struct {
	Shape Rect
}

func NewTile(x, y float32, c rl.Color) Tile {
	tile := Tile{
		Shape: NewRect(
			rl.Vector2{
				X: x,
				Y: y,
			},
			TileWidth, TileHeight, 0,
			c,
		),
	}
	tile.Shape.Filled = true
	return tile
}

type Tiles []Tile

type SpawnData struct {
	PlayerPos   rl.Vector2
	SpawnPoints []rl.Vector2
}

func (t *Tiles) LoadMap(filepath string) SpawnData {
	spawn := SpawnData{}

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var x, y float32
	for scanner.Scan() {
		x = 0
		for _, c := range scanner.Text() {
			switch c {
			case 'x':
				*t = append(*t, NewTile(x, y, rl.Blue))
			case 'p':
				spawn.PlayerPos = rl.Vector2{X: x, Y: y}
			case 's':
				spawn.SpawnPoints = append(spawn.SpawnPoints, rl.Vector2{X: x, Y: y})
			}
			x += TileWidth + 1
		}
		y += TileHeight + 1
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return spawn
}

func (t *Tiles) Render() {
	for _, tile := range *t {
		tile.Shape.Render()
	}
}
