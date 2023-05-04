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
	layers [100][100]int
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if g.layers[i][j] == 1 {
				c := rand.Intn(4)
				if c == 0 {
					if i > 0 {
						g.layers[i][j] = 0
						g.layers[i-1][j] = 1
					}
				} else if c == 1 {
					if i < 99 {
						g.layers[i][j] = 0
						g.layers[i+1][j] = 1
					}
				} else if c == 2 {
					if j > 0 {
						g.layers[i][j] = 0
						g.layers[i][j-1] = 1
					}
				} else if c == 3 {
					if j < 99 {
						g.layers[i][j] = 0
						g.layers[i][j+1] = 1
					}
				}
			}
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	tileXcount := len(g.layers[0])
	tileYcount := len(g.layers)

	x, _ := screen.Size()
	tileSize := x / tileXcount

	for i := 0; i < tileXcount; i++ {
		for j := 0; j < tileYcount; j++ {
			if g.layers[j][i] == 0 {
				ebitenutil.DrawRect(screen, float64(i*tileSize), float64(j*tileSize), float64(tileSize), float64(tileSize), color.White)
			} else {
				if g.layers[j][i] == 1 {
					ebitenutil.DrawRect(screen, float64(i*tileSize), float64(j*tileSize), float64(tileSize), float64(tileSize), color.Black)
				}
			}

		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 500
}

func main() {
	game := &Game{
		layers: generateMap()}
	// Specify the window size as you like. Here, a doubled size is specified.
	getSize()
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
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

func generateMap() [100][100]int {
	antMap := [100][100]int{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if rand.Intn(100) < 5 {
				antMap[i][j] = rand.Intn(2)
			}
		}
	}
	return antMap
}
