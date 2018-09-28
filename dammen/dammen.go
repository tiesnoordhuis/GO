package main

import (
	"fmt"
)

type Square struct {
	taken bool
	color string
	player string
  X int
  Y int
}

type Player struct {
	name string
	games_played int
	wins int
	losses int
}

type Board struct {
  squares [10][10]Square
}

type Game struct {
  turn int
  player1 Player
  player2 Player
  boards []Board
  finished bool
  winner Player
}

func main() {
  var player_1 Player = Player{name: "ties"}
  var player_2 Player = Player{name: "name2"}
	new_game := play_game(player_1, player_2)
  fmt.Println(new_game)
  display_board(new_game.boards[0])
}

func play_game(player_1 Player, player_2 Player) Game {
  var new_game Game = Game{turn: 0, player1: player_1, player2: player_2, finished: false}
  new_game.boards = make([]Board, 1)
  var new_board Board = build_empty_board()
  new_game.boards[0] = new_board
  return new_game
}

func build_empty_board() Board {
  var new_board Board
  for index_row := range new_board.squares {
    for index_colum := range new_board.squares[index_row] {
			if ((index_row + index_colum) % 2) == 0 {
        new_board.squares[index_row][index_colum] = Square{taken: false, color: "black", X: index_colum, Y: index_row}
      } else {
        new_board.squares[index_row][index_colum] = Square{taken: false, color: "white", X: index_colum, Y: index_row}
      }
		}
	}
  return new_board
}

func display_board(board Board) {
  for index_row := range board.squares {
    for index_colum := range board.squares[index_row] {
			if board.squares[index_row][index_colum].color == "black" {
        fmt.Printf("#")
      } else {
        fmt.Printf("O")
      }
		}
    fmt.Printf("\n")
	}
}
