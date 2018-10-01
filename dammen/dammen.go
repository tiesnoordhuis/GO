package main

import (
	"fmt"
	//"bufio"
	"image/color"
	//"os"

		"github.com/oakmound/oak"
	  "github.com/oakmound/oak/entities"
		"github.com/oakmound/oak/render"
		"github.com/oakmound/oak/scene"
		"github.com/oakmound/oak/dlog"
		"github.com/oakmound/oak/key"
		"github.com/oakmound/oak/event"
		"github.com/oakmound/oak/physics"
)

type Player struct {
	name string
	games_played int
	wins int
	losses int
}

type Game struct {
  turn int
  player1 Player
  player2 Player
  stukken_white Stukken
	stukken_black Stukken
  finished bool
  winner Player
}

type Stuk struct {
	X, Y int

	color string
	owner Player
}

type Stukken struct {
	stukken []Stuk
}

func main() {
	oak.SetupConfig.Screen.Width = 800
  oak.SetupConfig.Screen.Height = 800
		oak.Add("dammen", func(string, interface{}) {
			fmt.Print("name: ")
		  var player_1 Player = Player{name: "ties"}
		  var player_2 Player = Player{name: "name2"}
			var new_game Game = play_game(player_1, player_2)
			for i := 0; i < 5; i++ {
				for j := 0; j < 4; j++ {
					new_white_stuk := Stuk{((i * 2) + (j % 2)), j, "white", new_game.player1}
					new_game.stukken_white.stukken = append(new_game.stukken_white.stukken, new_white_stuk)

					new_black_stuk := Stuk{((i * 2) + (1 - (j % 2))), 10 - j, "black", new_game.player2}
					new_game.stukken_black.stukken = append(new_game.stukken_black.stukken, new_black_stuk)
				}
			}
		  fmt.Println(new_game)
			board_img, err := render.LoadSprite("", "board.png")
			dlog.ErrorCheck(err)
			render.Draw(board_img)
		  display_stukken(new_game.stukken_white, "white")
			display_stukken(new_game.stukken_black, "black")

			selector_white := entities.NewMoving(100, 100, 40, 40,
					render.NewColorBox(10, 10, color.RGBA{255, 255, 255, 100}),
					nil, 0, 0)
			selector_white.Speed = physics.NewVector(5, 5)

			render.Draw(selector_white.R)

			selector_white.Bind(func(id int, _ interface{}) int {
            selector_white := event.GetEntity(id).(*entities.Moving)
            selector_white.Delta.Zero()
            if oak.IsDown(key.W) {
                selector_white.Delta.ShiftY(-selector_white.Speed.Y())
            }
            if oak.IsDown(key.A) {
                selector_white.Delta.ShiftX(-selector_white.Speed.X())
            }
            if oak.IsDown(key.S) {
                selector_white.Delta.ShiftY(selector_white.Speed.Y())
            }
            if oak.IsDown(key.D) {
                selector_white.Delta.ShiftX(selector_white.Speed.X())
            }
            selector_white.ShiftPos(selector_white.Delta.X(), selector_white.Delta.Y())
						fmt.Println(selector_white)

            return 0
        }, event.Enter)

		}, func() bool {

				return true
		}, func() (string, *scene.Result) {
				return "dammen", nil
		})
		oak.Init("dammen")
}

func play_game(player_1 Player, player_2 Player) Game {
	stukken_white := Stukken{}
	stukken_black := Stukken{}
  var new_game Game = Game{turn: 0, player1: player_1, player2: player_2, finished: false, stukken_white: stukken_white, stukken_black: stukken_black}
  return new_game
}

func display_stukken(stukken Stukken, stuk_color string) {
	height := float64(oak.SetupConfig.Screen.Height)
	if stuk_color == "black" {
		for index, value := range stukken.stukken {
			stuk1 := entities.NewMoving(207 + (float64(value.X) * 40), -73 + ((float64(value.Y) * 40)), 40, 40,
					render.NewColorBox(24, 24, color.RGBA{255, 255, 0, 255}),
					nil, 0, 0)
			fmt.Println(index)
			render.Draw(stuk1.R)
		}
	} else if stuk_color == "white" {
		for index, value := range stukken.stukken {
			stuk1 := entities.NewMoving(207 + (float64(value.X) * 40), -233 + (height - (float64(value.Y) * 40)), 40, 40,
					render.NewColorBox(24, 24, color.RGBA{0, 255, 0, 255}),
					nil, 0, 0)
			fmt.Println(index)
			render.Draw(stuk1.R)
		}
	}
}
