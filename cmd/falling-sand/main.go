package main

import (
	game "github.com/GatoImorrivel/falling-sand/internal"
	"github.com/hajimehoshi/ebiten"
)

func main() {
  screenWidth := 848
  screenHeight := 484
  ebiten.SetWindowSize(screenWidth, screenHeight)
  ebiten.SetWindowTitle("Falling Sand")
  fallingSand := &game.FallingSand {
    ScreenWidth: screenWidth,
    ScreenHeight: screenHeight,
    ParticleSize: 2,
  }
  if err := ebiten.RunGame(fallingSand); err != nil {
    panic(err)
  }
}
