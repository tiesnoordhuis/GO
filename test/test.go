package main

import (
  "image/color"
  "time"
  "fmt"
  //"path/filepath"

  "github.com/oakmound/oak"
  "github.com/oakmound/oak/collision"
  "github.com/oakmound/oak/entities"
  "github.com/oakmound/oak/dlog"
  "github.com/oakmound/oak/event"
  "github.com/oakmound/oak/key"
  "github.com/oakmound/oak/mouse"
  "github.com/oakmound/oak/physics"
  "github.com/oakmound/oak/render"
  "github.com/oakmound/oak/scene"
)

const (
    Enemy  collision.Label = 1
    Player collision.Label = 2
    fieldWidth  = 1000
    fieldHeight = 1000
)

var (
    playerAlive = true
)

func main() {
  oak.SetupConfig.Screen.Width = 1000
  oak.SetupConfig.Screen.Height = 1000
  fmt.Println(oak.SetupConfig)

    oak.Add("tds", func(string, interface{}) {
        playerAlive = true
        oak.SetViewportBounds(0, 0, fieldWidth, fieldHeight)
        board_img, err := render.LoadSprite("", "board.png")
        dlog.ErrorCheck(err)
        render.Draw(board_img)
        char := entities.NewMoving(100, 100, 32, 32,
            render.NewColorBox(32, 32, color.RGBA{0, 255, 0, 255}),
            nil, 0, 0)

        char.Speed = physics.NewVector(5, 5)
        render.Draw(char.R)

        char.Bind(func(id int, _ interface{}) int {
            char := event.GetEntity(id).(*entities.Moving)
            char.Delta.Zero()
            if oak.IsDown(key.W) {
                char.Delta.ShiftY(-char.Speed.Y())
            }
            if oak.IsDown(key.A) {
                char.Delta.ShiftX(-char.Speed.X())
            }
            if oak.IsDown(key.S) {
                char.Delta.ShiftY(char.Speed.Y())
            }
            if oak.IsDown(key.D) {
                char.Delta.ShiftX(char.Speed.X())
            }
            char.ShiftPos(char.Delta.X(), char.Delta.Y())
            hit := char.HitLabel(Enemy)
            if hit != nil {
                playerAlive = false
            }

            return 0
        }, event.Enter)

        char.Bind(func(id int, me interface{}) int {
    			char := event.GetEntity(id).(*entities.Moving)
    			mevent := me.(mouse.Event)
    			render.DrawForTime(
    				render.NewLine(char.X()+char.W/2, char.Y()+char.H/2, mevent.X(), mevent.Y(), color.RGBA{0, 128, 0, 128}),
    				time.Millisecond*50,
    				1)
    			return 0
    		}, mouse.Press)

    }, func() bool {
        return playerAlive
    }, func() (string, *scene.Result) {
        return "tds", nil
    })
    oak.Init("tds")
}
