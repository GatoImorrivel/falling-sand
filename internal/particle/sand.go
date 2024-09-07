package particle

import "image/color"

type SandParticle struct {

}

func (p *SandParticle) Update() {

}

func (p *SandParticle) Color() color.Color {
  return color.RGBA {
    R: 248,
    G: 213,
    B: 113,
    A: 255,
  }
}
