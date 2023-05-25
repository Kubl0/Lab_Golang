package main

import (
	"flag"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2" //in v2.3.3
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct {
	layers [100][100]Organism
}

type Organism struct {
	orgType          string
	timeToMove       int
	currentMove      int
	timeToTurn       int
	currentTurn      int
	direction        int
	isCarrying       bool
	carryTime        int
	currentCarryTime int
	energy           int
	neededEnergy     int
}

func (game *Game) freePositionsAround(x, y int) [][]int {
	var freePositions [][]int
	if x != 0 && game.layers[x-1][y].orgType == "none" {
		freePositions = append(freePositions, []int{x - 1, y})
	}
	if x != 99 && game.layers[x+1][y].orgType == "none" {
		freePositions = append(freePositions, []int{x + 1, y})
	}
	if y != 0 && game.layers[x][y-1].orgType == "none" {
		freePositions = append(freePositions, []int{x, y - 1})
	}
	if y != 99 && game.layers[x][y+1].orgType == "none" {
		freePositions = append(freePositions, []int{x, y + 1})
	}
	return freePositions
}

func (game *Game) moveStraight(x, y, direction int) {
	if direction == 0 {
		if x != 0 && game.layers[x-1][y].orgType == "none" {
			game.layers[x-1][y] = game.layers[x][y]
			game.layers[x-1][y].currentMove = 0
			game.layers[x][y] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
		} else if x == 0 || game.layers[x-1][y].orgType == "ant" {
			game.layers[x][y].direction = 1
		}
	} else if direction == 1 {
		if x != 99 && game.layers[x+1][y].orgType == "none" {
			game.layers[x+1][y] = game.layers[x][y]
			game.layers[x+1][y].currentMove = 0
			game.layers[x][y] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
		} else if x == 99 || game.layers[x+1][y].orgType == "ant" {
			game.layers[x][y].direction = 0
		}
	} else if direction == 2 {
		if y != 0 && game.layers[x][y-1].orgType == "none" {
			game.layers[x][y-1] = game.layers[x][y]
			game.layers[x][y-1].currentMove = 0
			game.layers[x][y] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
		} else if y == 0 || game.layers[x][y-1].orgType == "ant" {
			game.layers[x][y].direction = 3
		}
	} else if direction == 3 {
		if y != 99 && game.layers[x][y+1].orgType == "none" {
			game.layers[x][y+1] = game.layers[x][y]
			game.layers[x][y+1].currentMove = 0
			game.layers[x][y] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
		} else if y == 99 || game.layers[x][y+1].orgType == "ant" {
			game.layers[x][y].direction = 2
		}
	}
}

func (game *Game) isLeafNear(x, y int, ant *Organism) {
	if x != 0 && game.layers[x-1][y].orgType == "leaf" {
		game.layers[x][y].isCarrying = true
		game.layers[x-1][y] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
	} else if x != 99 && game.layers[x+1][y].orgType == "leaf" {
		game.layers[x][y].isCarrying = true
		game.layers[x+1][y] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
	} else if y != 0 && game.layers[x][y-1].orgType == "leaf" {
		game.layers[x][y].isCarrying = true
		game.layers[x][y-1] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
	} else if y != 99 && game.layers[x][y+1].orgType == "leaf" {
		game.layers[x][y].isCarrying = true
		game.layers[x][y+1] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
	}
}

func (game *Game) carryTime(x, y int, ant *Organism) {
	if ant.isCarrying {
		if ant.currentCarryTime < ant.carryTime {
			ant.currentCarryTime++
		} else {
			pos := game.freePositionsAround(x, y)
			if len(pos) != 0 {
				pos := pos[rand.Intn(len(pos))]
				game.layers[pos[0]][pos[1]] = Organism{"leaf", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
				game.layers[x][y].isCarrying = false
				game.layers[x][y].currentCarryTime = 0
				game.layers[x][y].energy = 0
			}
		}
	}
}

func (game *Game) move(x, y int, ant *Organism) {
	if ant.currentMove < ant.timeToMove {
		game.layers[x][y].currentMove++
	}
	if ant.currentTurn < ant.timeToTurn {
		game.layers[x][y].currentTurn++
	}

	if ant.energy < ant.neededEnergy {
		game.layers[x][y].energy++
	}

	ant1 := game.layers[x][y]

	if ant1.currentMove == ant1.timeToMove {
		if ant1.currentTurn == ant1.timeToTurn {
			freePositions := game.freePositionsAround(x, y)
			if len(freePositions) != 0 {
				pos := freePositions[rand.Intn(len(freePositions))]
				game.layers[pos[0]][pos[1]] = game.layers[x][y]
				game.layers[pos[0]][pos[1]].currentMove = 0
				game.layers[pos[0]][pos[1]].currentTurn = 0
				game.layers[x][y] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
			}
		} else {
			game.moveStraight(x, y, ant.direction)
		}
	}

	if ant.orgType == "ant" {
		if ant.energy == ant.neededEnergy {
			game.isLeafNear(x, y, ant)
		}
		game.carryTime(x, y, ant)
	}
}

func (g *Game) Update() error {
	for tileX := range g.layers {
		for tileY := range g.layers[tileX] {
			if g.layers[tileX][tileY].orgType == "ant" {
				g.move(tileX, tileY, &g.layers[tileX][tileY])
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
			if g.layers[tileX][tileY].orgType == "ant" && g.layers[tileX][tileY].isCarrying == false {
				ebitenutil.DrawRect(screen, float64(tileX*tileSize), float64(tileY*tileSize), float64(tileSize), float64(tileSize), color.White)
			} else if g.layers[tileX][tileY].orgType == "leaf" {
				ebitenutil.DrawRect(screen, float64(tileX*tileSize), float64(tileY*tileSize), float64(tileSize), float64(tileSize), color.RGBA{0, 255, 0, 255})
			} else if g.layers[tileX][tileY].orgType == "ant" && g.layers[tileX][tileY].isCarrying == true {
				ebitenutil.DrawRect(screen, float64(tileX*tileSize), float64(tileY*tileSize), float64(tileSize), float64(tileSize), color.RGBA{0, 0, 255, 255})
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
			if rand.Intn(1000) < 50 {
				ant := &Organism{"ant", 10, 0, 5, 0, rand.Intn(4), false, 50, 0, 0, 20}
				game.layers[i][j] = *ant
			} else if rand.Intn(1000) > 990 {
				leaf := &Organism{"leaf", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
				game.layers[i][j] = *leaf
			} else {
				game.layers[i][j] = Organism{"none", 0, 0, 0, 0, 0, false, 0, 0, 0, 0}
			}
		}
	}
	return game.layers
}
