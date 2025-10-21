package main

import (
	"fmt"
	"log"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	sWidth = 1000
	sHeight = 800
)

type Particle struct{
	posX float64
	posY float64
	vel Velocity
	rad float64
	color color.Color	
}

type Game struct{
	particles []Particle
}

type Velocity struct {
	dirX float64
	dirY float64
}

func makeParticle(posX float64, posY float64, vel Velocity, rad float64, color color.Color) Particle {
	particle := Particle{
		posX: posX,
		posY: posY,
		vel: vel,
		rad: rad,
		color: color,
	}
	return particle
}

func (g *Game) Init(){
	particle := makeParticle(sWidth/2, sHeight/2, Velocity{rand.Float64() * 10, rand.Float64() * 10}, 20, color.RGBA{255, 0, 0, 255})
	g.particles = append(g.particles, particle)
}

func preventWallCollision(p *Particle) {
	if p.posX - p.rad < 0 || p.posX + p.rad > sWidth {
		p.vel.dirX = -p.vel.dirX
	}
	if p.posY - p.rad < 0 || p.posY + p.rad > sHeight {
		p.vel.dirY = -p.vel.dirY
	}
}

func (g *Game) Update() error {

	for i := range g.particles {
		p := &g.particles[i]
		p.posX += p.vel.dirX
		p.posY += p.vel.dirY
		preventWallCollision(p)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		randRad := rand.Float64() * 50
		randPosX := rand.Float64() * sWidth - randRad
		randPosY := rand.Float64() * sHeight - randRad
		randVel := Velocity{
			dirX: -10 + rand.Float64() * 20,
			dirY: -10 + rand.Float64() * 20,
		}

		if randRad<10 { randRad = 10 }

		randColor := color.RGBA{uint8(rand.Int() * 255), uint8(rand.Int() * 255), uint8(rand.Int() * 255), uint8(rand.Int() * 255)}

		particle := makeParticle(randPosX, randPosY, randVel, randRad, randColor)
		g.particles = append(g.particles, particle)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

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
	ebiten.SetWindowTitle("moving balls")

	game := Game{}
	game.Init();

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(game.particles))

}
