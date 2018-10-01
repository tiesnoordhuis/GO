package main

import (
  //"image"
	"image/color"
	//"math/rand"
	//"path/filepath"
	//"time"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/alg/floatgeom"
	//"github.com/oakmound/oak/collision"
	//"github.com/oakmound/oak/collision/ray"
	//"github.com/oakmound/oak/dlog"
	//"github.com/oakmound/oak/entities"
	//"github.com/oakmound/oak/event"
	//"github.com/oakmound/oak/key"
	//"github.com/oakmound/oak/mouse"
	//"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
  //"github.com/oakmound/oak/render/mod"
	"github.com/oakmound/oak/scene"
)

var game_ongoing bool
var turn int

func main() {
	oak.Add("game", func(string, interface{}) {
		// Initialization
    game_ongoing = true
		turn = 0
		sprite1 := render.NewColorBox(50, 50, color.RGBA{0, 128, 0, 128})
    //sprite2 := render.NewColorBox(50, 50, color.RGBA{128, 0, 0, 128})
		render.NewCompositeM()
		var composite_both render.CompositeM
		counter := 0
    for i := 0; i < 10; i++ {
      for j := 0; j < 10; j++ {
				if ((i + j) % 2) == 0 {
					composite_both.Append(render.NewColorBox(50, 50, color.RGBA{128, 0, 0, 128}))
					composite_both.AddOffset(counter, floatgeom.Point2{float64(i) * 100, float64(j) * 100})
					counter++
				} else {
					composite_both.Append(render.NewColorBox(50, 50, color.RGBA{128, 128, 0, 128}))
					composite_both.AddOffset(counter, floatgeom.Point2{float64(i) * 100 + 25, float64(j) * 100 + 25})
					counter++
				}

      }
    }
		render.Draw(composite_both)
	}, func() bool {
    turn++
		return game_ongoing
	}, func() (string, *scene.Result) {
		return "game", nil
	})

	// This indicates to oak to automatically open and load image and audio
	// files local to the project before starting any scene.
	//oak.SetupConfig.BatchLoad = true

	oak.Init("game")
}
