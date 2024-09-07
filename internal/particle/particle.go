package particle

import "image/color"

type Particle interface {
  Update()  
  Color() color.Color
}
