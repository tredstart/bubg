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
	Shape    Rect
	Material rl.Color
}

func NewTile(x, y float32, c rl.Color) Tile {
	return Tile{
		Shape: Rect{
			Rectangle: rl.NewRectangle(x, y, TileWidth, TileHeight),
		},
		Material: c,
	}
}

type Tiles []Tile

func (t *Tiles) LoadMap(filepath string) {
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
			if c == 'x' {
				*t = append(*t, NewTile(x, y, rl.Blue))
			}
			x += TileWidth + 1
		}
		y += TileHeight + 1
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (t *Tiles) Render() {
	for _, tile := range *t {
		rl.DrawRectangleRec(tile.Shape.Rectangle, tile.Material)
	}
}
