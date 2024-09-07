package game

import (
	"image/color"

	"github.com/GatoImorrivel/falling-sand/internal/particle"
	"github.com/hajimehoshi/ebiten"
)

type FallingSand struct {
  ScreenWidth int
  ScreenHeight int
  ParticleSize int
  particles []particle.Particle
}

func (g *FallingSand) Update(screen *ebiten.Image) error {
	return nil
}

func (g *FallingSand) Draw(screen *ebiten.Image) {
  screen.Fill(color.RGBA {
    R: 0x00,
    G: 0x00,
    B: 0x00,
    A: 0xff,
  })
}

func (g *FallingSand) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return g.ScreenWidth, g.ScreenHeight
}
