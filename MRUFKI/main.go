package main

import (
	"flag"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct {
	layers [100][100]Organism
}

type Organism struct {
	orgType     string
	timeToMove  int
	currentMove int
}

func (game *Game) freePositionsAround(x, y int) [][]int {
	var freePositions [][]int
	if x != 0 && game.layers[x-1][y].timeToMove == 0 {
		freePositions = append(freePositions, []int{x - 1, y})
	}
	if x != 99 && game.layers[x+1][y].timeToMove == 0 {
		freePositions = append(freePositions, []int{x + 1, y})
	}
	if y != 0 && game.layers[x][y-1].timeToMove == 0 {
		freePositions = append(freePositions, []int{x, y - 1})
	}
	if y != 99 && game.layers[x][y+1].timeToMove == 0 {
		freePositions = append(freePositions, []int{x, y + 1})
	}
	return freePositions
}

func (g *Game) Update() error {
	for tileX := range g.layers {
		for tileY := range g.layers[tileX] {
			if g.layers[tileX][tileY].orgType == "ant" {
				if g.layers[tileX][tileY].timeToMove != 0 {
					if g.layers[tileX][tileY].currentMove < g.layers[tileX][tileY].timeToMove {
						g.layers[tileX][tileY].currentMove++
					} else {
						freePositions := g.freePositionsAround(tileX, tileY)
						if len(freePositions) != 0 {
							pos := freePositions[rand.Intn(len(freePositions))]
							g.layers[pos[0]][pos[1]] = g.layers[tileX][tileY]
							g.layers[pos[0]][pos[1]].currentMove = 0
							g.layers[tileX][tileY] = Organism{"none", 0, 0}
						}
					}

				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tileXcount := len(g.layers[0])

	x, _ := screen.Size()
	tileSize := x / tileXcount

	for tileX := range g.layers {
		for tileY := range g.layers[tileX] {
			if g.layers[tileX][tileY].timeToMove != 0 {
				ebitenutil.DrawRect(screen, float64(tileX*tileSize), float64(tileY*tileSize), float64(tileSize), float64(tileSize), color.White)
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 500
}

func main() {
	game := &Game{
		layers: [100][100]Organism{}}
	game.generateMap()
	getSize()
	ebiten.SetWindowTitle("MRUFKI")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func getSize() {
	w := flag.Int("w", 640, "window width")
	h := flag.Int("h", 480, "window height")
	flag.Parse()

	ebiten.SetWindowSize(*w, *h)
}

func (game *Game) generateMap() [100][100]Organism {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if rand.Intn(100) < 5 {
				ant := &Organism{"ant", 3, 0}
				game.layers[i][j] = *ant
			} else {
				game.layers[i][j] = Organism{"none", 0, 0}
			}
		}
	}
	return game.layers
}
