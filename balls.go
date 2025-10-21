package main

import (
	"fmt"
	"log"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	sWidth = 640
	sHeight = 480
)

type Particle struct{
	posX float64
	posY float64
	rad float64
	color color.Color	
}

type Game struct{
	particles []Particle
}

func makeParticle(posX float64, posY float64, rad float64, color color.Color) Particle {
	particle := Particle{
		posX: posX,
		posY: posY,
		rad: rad,
		color: color,
	}
	return particle
}

func (g *Game) Init(){
	particle := makeParticle(sWidth/2, sHeight/2, 20, color.RGBA{255, 0, 0, 255})
	g.particles = append(g.particles, particle)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.particles {
		ebitenutil.DrawCircle(screen, v.posX, v.posY, v.rad, v.color)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}

func main() {
	ebiten.SetWindowSize(sWidth, sHeight)
	ebiten.SetWindowTitle("Hello, World!")

	game := Game{}
	game.Init();

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
	for _, v := range game.particles {
		fmt.Println(v)
	}
}

