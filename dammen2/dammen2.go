package main

import (
	"fmt"
	"bufio"
  //"image"
	//"image/color"
	"os"
  "strconv"
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

type Point2 struct {
  X, Y int
}

type Stuk struct {
	location Point2
	color string
  display_string string
	owner Player
  upgraded bool
}

type Stukken struct {
	stukken []Stuk
}

type Board struct {
  display_stukken [10][10]string
}


func main() {
      names := ask_names()
		  var player_1 Player = Player{name: names[0]}
		  var player_2 Player = Player{name: names[1]}
			var new_game Game = play_game(player_1, player_2)
			for i := 0; i < 5; i++ {
				for j := 0; j < 4; j++ {
          location_white := Point2{((i * 2) + (j % 2)), j}
					new_white_stuk := Stuk{location_white, "white", "O", new_game.player1, false}
					new_game.stukken_white.stukken = append(new_game.stukken_white.stukken, new_white_stuk)

          location_black := Point2{((i * 2) + (1 - (j % 2))), 9 - j}
					new_black_stuk := Stuk{location_black, "black", "#", new_game.player2, false}
					new_game.stukken_black.stukken = append(new_game.stukken_black.stukken, new_black_stuk)
				}
			}
      empty_board_grid := [10][10]string{{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
        {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
      }
      empty_board := Board{empty_board_grid}
      display_stukken(new_game.stukken_white, new_game.stukken_black, empty_board)

      //run game
      for new_game.finished == false {
        stuk_moved := false
        stuk_hit_but_continue := false
        //white move
        for stuk_moved == false {
          //make values to check if forced hit
          var pos_forced_hit []Point2
          var pos_forced_hit_to []Point2
          for index := range new_game.stukken_white.stukken {
            hit_slice := new_game.stukken_white.stukken[index].check_hit(new_game)
            if len(hit_slice) > 1 {
              pos_forced_hit = append(pos_forced_hit, hit_slice[0])
              for i := 1; i < len(hit_slice); i++ {
                pos_forced_hit_to = append(pos_forced_hit_to, hit_slice[i])
              }
            }
          }

          // means you have to hit
          if len(pos_forced_hit) > 0 {
            fmt.Println("you have to hit")
            user_want_move_white := ask_move("white")
            //fmt.Println(user_want_move_white)
            //stuk_to_move_white := find_stuk(user_want_move_white[0], new_game)
            //fmt.Println(stuk_to_move_white)
            stuk_moved = false
            stuk_hit_but_continue = false
            var stuk_remove_location Point2
            for _, value := range pos_forced_hit_to {
              if user_want_move_white[1] == value && contains(pos_forced_hit, user_want_move_white[0]) {
                index_moving := find_index(user_want_move_white[0], new_game)
                new_game.stukken_white.stukken[index_moving].location = value
                stuk_moved = true
                // get location of the stuk to remove
                if new_game.stukken_black.stukken[index_moving].upgraded {
                  stuk_remove_location = get_location_removing_stuk_upgraded(user_want_move_white, "white", new_game)
                } else {
                  stuk_remove_location.X = (user_want_move_white[1].X + user_want_move_white[0].X) / 2
                  stuk_remove_location.Y = (user_want_move_white[1].Y + user_want_move_white[0].Y) / 2
                }
                index_of_removing := find_index(stuk_remove_location, new_game)
                new_game.stukken_black.stukken = append(new_game.stukken_black.stukken[:index_of_removing], new_game.stukken_black.stukken[index_of_removing + 1:]...)
                fmt.Println("succes")
                //check if more hits are possible after this
                for index := range new_game.stukken_white.stukken {
                  hit_slice_again := new_game.stukken_white.stukken[index].check_hit(new_game)
                  if len(hit_slice_again) > 1 {
                    stuk_moved = false
                    stuk_hit_but_continue = true
                    break
                  }
                }
              }
            }
            if stuk_moved == false {
              if stuk_hit_but_continue {
                fmt.Println("you have to hit again")
              } else {
                fmt.Println("stuk not moved")
              }
            }
          } else {

            user_want_move_white := ask_move("white")
            //fmt.Println(user_want_move_white)
            stuk_to_move_white := find_stuk(user_want_move_white[0], new_game)
            //fmt.Println(stuk_to_move_white)
            move_pos_white := stuk_to_move_white.make_move_pos()
            //fmt.Println(move_pos_white)
            movable_pos_white := check_move_pos(move_pos_white, stuk_to_move_white.location, stuk_to_move_white.upgraded, new_game)
            //fmt.Println(movable_pos_white)
            stuk_moved = false
            for _, value := range movable_pos_white {
              if user_want_move_white[1] == value {
                index_moving := find_index(user_want_move_white[0], new_game)
                if new_game.stukken_white.stukken[index_moving].upgraded == false && value.Y > stuk_to_move_white.location.Y {
                  new_game.stukken_white.stukken[index_moving].location = value
                  stuk_moved = true
                  fmt.Println("succes")
                } else if new_game.stukken_white.stukken[index_moving].upgraded {
                  new_game.stukken_white.stukken[index_moving].location = value
                  stuk_moved = true
                  fmt.Println("succes")
                }
              }
            }
            if stuk_moved == false {
              fmt.Println("stuk not moved")
            }
          }
          //upgrade stuk to Dam
          upgrade_stuk_pos_white := check_to_upgrade(new_game)
          if len(upgrade_stuk_pos_white) > 0 {
            upgrade_index := find_index(upgrade_stuk_pos_white[0], new_game)
            new_game.stukken_white.stukken[upgrade_index].upgraded = true
            new_game.stukken_white.stukken[upgrade_index].display_string = "@"
          }
          display_stukken(new_game.stukken_white, new_game.stukken_black, empty_board)

          new_game.finished = check_finished(new_game)
        }

        stuk_moved = false

        fmt.Println("switch player")

        //black move
        for stuk_moved == false {
          //make values to check if forced hit
          var pos_forced_hit []Point2
          var pos_forced_hit_to []Point2
          for index := range new_game.stukken_black.stukken {
            hit_slice := new_game.stukken_black.stukken[index].check_hit(new_game)
            if len(hit_slice) > 1 {
              pos_forced_hit = append(pos_forced_hit, hit_slice[0])
              for i := 1; i < len(hit_slice); i++ {
                pos_forced_hit_to = append(pos_forced_hit_to, hit_slice[i])
              }
            }
          }

          // means you have to hit
          if len(pos_forced_hit) > 0 {
            fmt.Println("you have to hit")
            user_want_move_black := ask_move("black")
            //fmt.Println(user_want_move_black)
            //stuk_to_move_black := find_stuk(user_want_move_black[0], new_game)
            //fmt.Println(stuk_to_move_black)
            stuk_moved = false
            stuk_hit_but_continue = false
            var stuk_remove_location Point2
            for _, value := range pos_forced_hit_to {
              if user_want_move_black[1] == value && contains(pos_forced_hit, user_want_move_black[0]) {
                index_moving := find_index(user_want_move_black[0], new_game)
                new_game.stukken_black.stukken[index_moving].location = value
                stuk_moved = true
                // get location of the stuk to remove
                if new_game.stukken_black.stukken[index_moving].upgraded {
                  stuk_remove_location = get_location_removing_stuk_upgraded(user_want_move_black, "black", new_game)
                } else {
                  stuk_remove_location.X = (user_want_move_black[1].X + user_want_move_black[0].X) / 2
                  stuk_remove_location.Y = (user_want_move_black[1].Y + user_want_move_black[0].Y) / 2
                }
                index_of_removing := find_index(stuk_remove_location, new_game)
                new_game.stukken_white.stukken = append(new_game.stukken_white.stukken[:index_of_removing], new_game.stukken_white.stukken[index_of_removing + 1:]...)
                fmt.Println("succes")
                //check if more hits are possible after this
                for index := range new_game.stukken_black.stukken {
                  hit_slice_again := new_game.stukken_black.stukken[index].check_hit(new_game)
                  if len(hit_slice_again) > 1 {
                    stuk_moved = false
                    stuk_hit_but_continue = true
                    break
                  }
                }
              }
            }
            if stuk_moved == false {
              if stuk_hit_but_continue {
                fmt.Println("you have to hit again")
              } else {
                fmt.Println("stuk not moved")
              }
            }
          } else {
            user_want_move_black := ask_move("black")
            //fmt.Println(user_want_move_black)
            stuk_to_move_black := find_stuk(user_want_move_black[0], new_game)
            //fmt.Println(stuk_to_move_black)
            move_pos_black := stuk_to_move_black.make_move_pos()
            //fmt.Println(move_pos_black)
            movable_pos_black := check_move_pos(move_pos_black, stuk_to_move_black.location, stuk_to_move_black.upgraded, new_game)
            //fmt.Println(movable_pos_black)
            stuk_moved = false
            for _, value := range movable_pos_black {
              if user_want_move_black[1] == value {
                index_moving := find_index(user_want_move_black[0], new_game)
                if new_game.stukken_black.stukken[index_moving].upgraded == false && value.Y < stuk_to_move_black.location.Y {
                  new_game.stukken_black.stukken[index_moving].location = value
                  stuk_moved = true
                  fmt.Println("succes")
                } else if new_game.stukken_black.stukken[index_moving].upgraded {
                  new_game.stukken_black.stukken[index_moving].location = value
                  stuk_moved = true
                  fmt.Println("succes")
                }
              }
            }
            if stuk_moved == false {
              fmt.Println("stuk not moved")
            }
          }
          //upgrade stuk to Dam
          upgrade_stuk_pos_black := check_to_upgrade(new_game)
          if len(upgrade_stuk_pos_black) > 0 {
            upgrade_index := find_index(upgrade_stuk_pos_black[0], new_game)
            new_game.stukken_black.stukken[upgrade_index].upgraded = true
            new_game.stukken_black.stukken[upgrade_index].display_string = "%"
          }

          display_stukken(new_game.stukken_white, new_game.stukken_black, empty_board)

          new_game.finished = check_finished(new_game)
        }
      }
    winner := find_winner(new_game)
    fmt.Println("congrats ", winner, ", you have won the game")
}

func play_game(player_1 Player, player_2 Player) Game {
	stukken_white := Stukken{}
	stukken_black := Stukken{}
  var new_game Game = Game{turn: 0, player1: player_1, player2: player_2, finished: false, stukken_white: stukken_white, stukken_black: stukken_black}
  return new_game
}

func display_stukken(stukken_white Stukken, stukken_black Stukken, empty_board Board) {
  x_as_label := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Println("row")
	for _, value := range stukken_black.stukken {
    empty_board.display_stukken[value.location.Y][value.location.X] = value.display_string
		//fmt.Println(value)
  }
  for _, value := range stukken_white.stukken {
    empty_board.display_stukken[value.location.Y][value.location.X] = value.display_string
		//fmt.Println(value)
	}
  for index, value := range empty_board.display_stukken {
    fmt.Println(index, value)
  }

  fmt.Println(" ", x_as_label, "colum")
}

func (stuk Stuk) make_move_pos() []Point2 {
  var return_array []Point2
  //not upgrade stuk has 4 possible moves to make
  if stuk.upgraded == false {
    X := stuk.location.X
    Y := stuk.location.Y
    return_array = append(return_array, Point2{X - 1, Y - 1})
    return_array = append(return_array, Point2{X - 1, Y + 1})
    return_array = append(return_array, Point2{X + 1, Y - 1})
    return_array = append(return_array, Point2{X + 1, Y + 1})
  } else {
    return_array = make_move_pos_upgrade(stuk)
  }
  return return_array
}

func make_move_pos_upgrade(stuk Stuk) []Point2 {
  var return_array []Point2
  X := stuk.location.X
  Y := stuk.location.Y
  for i := 1; i < 10; i++ {
    higher_x := X + i
    if higher_x >= 0 && higher_x < 10 {
      higher_y := Y + i
      if higher_y >= 0 && higher_y < 10 {
        return_array = append(return_array, Point2{higher_x, higher_y})
      }
      lower_y := Y - i
      if lower_y >= 0 && lower_y < 10 {
        return_array = append(return_array, Point2{higher_x, lower_y})
      }
    }
    lower_x := X - i
    if lower_x >= 0 && lower_x < 10 {
      higher_y := Y + i
      if higher_y >= 0 && higher_y < 10 {
        return_array = append(return_array, Point2{lower_x, higher_y})
      }
      lower_y := Y + i
      if lower_y >= 0 && lower_y < 10 {
        return_array = append(return_array, Point2{lower_x, lower_y})
      }
    }
  }
  //fmt.Println("upgraded stuk move pos: ", return_array)
  return return_array
}

func check_move_pos(move_pos []Point2, move_from Point2, is_upgraded bool, game Game) []Point2 {
  var return_array []Point2
  //new function for checking valid places to move as upgraded stuk
  if is_upgraded{
    for _, value := range move_pos {
      if check_move_pos_upgraded(value, move_from, game) {
        return_array = append(return_array, value)
      }
    }
  } else {
    for _, value := range move_pos {
      if value.X < 0 || value.X > 9 {
        continue
      } else if value.Y < 0 || value.Y > 9 {
        continue
      }
      collision := false
      for _, value_other := range game.stukken_white.stukken {
        if value_other.location.X == value.X && value_other.location.Y == value.Y {
          collision = true
          break
        }
      }
      for _, value_other := range game.stukken_black.stukken {
        if value_other.location.X == value.X && value_other.location.Y == value.Y {
          collision = true
          break
        }
      }
      if collision {
        continue
      } else {
        return_array = append(return_array, value)
      }
    }
  }
  return return_array
}

func check_move_pos_upgraded(move_to Point2, move_from Point2, game Game) bool {
  stuk_X := move_from.X
  stuk_Y := move_from.Y
  delta_x := move_to.X - stuk_X
  delta_y := move_to.Y - stuk_Y
  // build location slice from stukken white and black
  var white_stukken_locations []Point2
  for _, value := range game.stukken_white.stukken {
    white_stukken_locations = append(white_stukken_locations, value.location)
  }
  var black_stukken_locations []Point2
  for _, value := range game.stukken_black.stukken {
    black_stukken_locations = append(black_stukken_locations, value.location)
  }
  var not_hit_stuk bool = true
    if delta_x > 1 {
      if delta_y > 1 {
        for i := 1; i <= delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            not_hit_stuk = false
          } else if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            not_hit_stuk = false
          }
        }
      } else {
        for i := 1; i <= delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            not_hit_stuk = false
          } else if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            not_hit_stuk = false
          }
        }
      }
    } else {
      if delta_y > 1 {
        for i := 1; i <= -delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            not_hit_stuk = false
          } else if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            not_hit_stuk = false
          }
        }
      } else {
        for i := 1; i <= -delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            not_hit_stuk = false
          } else if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            not_hit_stuk = false
          }
        }
      }
    }
  return not_hit_stuk
}

func check_hit_pos(move_pos []Point2, game Game, hitting_color string) []Point2 {
  var return_array []Point2
  for _, value := range move_pos {
    if value.X < 0 || value.X > 9 {
      continue
    } else if value.Y < 0 || value.Y > 9 {
      continue
    }
    if hitting_color == "white" {
      for _, value_other := range game.stukken_black.stukken {
        if value_other.location.X == value.X && value_other.location.Y == value.Y {
          return_array = append(return_array, value)
          break
        }
      }
    }
    if hitting_color == "black" {
      for _, value_other := range game.stukken_white.stukken {
        if value_other.location.X == value.X && value_other.location.Y == value.Y {
          return_array = append(return_array, value)
          break
        }
      }
    }
  }
  //fmt.Println("hittable locations", return_array)
  return return_array
}

func (stuk Stuk) check_hit(game Game) []Point2 {
  var return_array []Point2
  if stuk.upgraded == false{
    move_pos := stuk.make_move_pos()
    hit_pos := check_hit_pos(move_pos, game, stuk.color)
    for index, value := range hit_pos {
      delta_x := value.X - stuk.location.X
      delta_y := value.Y - stuk.location.Y
      hit_pos[index].X = value.X + delta_x
      hit_pos[index].Y = value.Y + delta_y
    }
    return_array = append(return_array, stuk.location)
    move_after_hit_pos := check_move_pos(hit_pos, Point2{}, false, game)
    if len(move_after_hit_pos) > 0 {
      for _, value := range move_after_hit_pos {
        return_array = append(return_array, value)
      }
    }
    return return_array
  // if stuk upgraded then :
  } else {
    has_hit_option := false
    var move_after_hit_pos_upgraded []Point2
    move_pos_upgraded := stuk.make_move_pos()
    for _, value := range move_pos_upgraded {
      if check_hit_pos_upgraded(value, stuk.location, stuk.color, game) {
        has_hit_option = true
        move_after_hit_pos_upgraded = append(move_after_hit_pos_upgraded, value)
      }
    }
    if has_hit_option {
      return_array = append(return_array, stuk.location)
      for _, value_append := range move_after_hit_pos_upgraded {
        return_array = append(return_array, value_append)
      }
    }
    return return_array
  }
}

func check_hit_pos_upgraded(move_after_hit_pos_upgraded Point2, stuk_upgraded Point2, hitting_color string, game Game) bool {
  var return_bool = false
  stuk_X := stuk_upgraded.X
  stuk_Y := stuk_upgraded.Y
  delta_x := move_after_hit_pos_upgraded.X - stuk_X
  delta_y := move_after_hit_pos_upgraded.Y - stuk_Y
  // move 1 over can never hit
  if delta_x >= -1 && delta_x <= 1 {
    return false
  }
  if delta_y >= -1 && delta_y <= 1 {
    return false
  }
  // build location slice from stukken white and black
  var white_stukken_locations []Point2
  for _, value := range game.stukken_white.stukken {
    white_stukken_locations = append(white_stukken_locations, value.location)
  }
  var black_stukken_locations []Point2
  for _, value := range game.stukken_black.stukken {
    black_stukken_locations = append(black_stukken_locations, value.location)
  }
  var has_hit bool = false
  if hitting_color == "white" {
    if delta_x > 1 {
      if delta_y > 1 {
        for i := 1; i <= delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            return false
          } else if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      } else {
        for i := 1; i <= delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            return false
          } else if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      }
    } else {
      if delta_y > 1 {
        for i := 1; i <= -delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            return false
          } else if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      } else {
        for i := 1; i <= -delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            return false
          } else if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      }
    }
    if has_hit {
      return true
    }
  //if not white then black :
  } else {
    if delta_x > 1 {
      if delta_y > 1 {
        for i := 1; i <= delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            return false
          } else if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      } else {
        for i := 1; i <= delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            return false
          } else if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      }
    } else {
      if delta_y > 1 {
        for i := 1; i <= -delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            return false
          } else if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      } else {
        for i := 1; i <= -delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            return false
          } else if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            if has_hit {
              return false
            } else {
              has_hit = true
            }
          }
        }
      }
    }
    if has_hit {
      return true
    }
  }
  return return_bool
}

func get_location_removing_stuk_upgraded(move_vector [2]Point2, moving_color string, game Game) Point2 {
  stuk_X := move_vector[0].X
  stuk_Y := move_vector[0].Y
  delta_x := move_vector[1].X - move_vector[0].X
  delta_y := move_vector[1].Y - move_vector[0].Y
  // build location slice from stukken white and black
  var white_stukken_locations []Point2
  for _, value := range game.stukken_white.stukken {
    white_stukken_locations = append(white_stukken_locations, value.location)
  }
  var black_stukken_locations []Point2
  for _, value := range game.stukken_black.stukken {
    black_stukken_locations = append(black_stukken_locations, value.location)
  }
  if moving_color == "white" {
    if delta_x > 1 {
      if delta_y > 1 {
        for i := 1; i <= delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            return Point2{stuk_X + i, stuk_Y + i}
          }
        }
      } else {
        for i := 1; i <= delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            return Point2{stuk_X + i, stuk_Y - i}
          }
        }
      }
    } else {
      if delta_y > 1 {
        for i := 1; i <= -delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            return Point2{stuk_X - i, stuk_Y + i}
          }
        }
      } else {
        for i := 1; i <= -delta_x; i++ {
          if contains(black_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            return Point2{stuk_X - i, stuk_Y - i}
          }
        }
      }
    }
  //if not white then black :
  } else {
    if delta_x > 1 {
      if delta_y > 1 {
        for i := 1; i <= delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y + i}) {
            return Point2{stuk_X + i, stuk_Y + i}
          }
        }
      } else {
        for i := 1; i <= delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X + i, stuk_Y - i}) {
            return Point2{stuk_X + i, stuk_Y - i}
          }
        }
      }
    } else {
      if delta_y > 1 {
        for i := 1; i <= -delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y + i}) {
            return Point2{stuk_X - i, stuk_Y + i}
          }
        }
      } else {
        for i := 1; i <= -delta_x; i++ {
          if contains(white_stukken_locations, Point2{stuk_X - i, stuk_Y - i}) {
            return Point2{stuk_X - i, stuk_Y - i}
          }
        }
      }
    }
  }
  return Point2{}
}

func ask_names() [2]string {
  var return_array [2]string
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("player 1 (white O), enter your name: ")
  name1, _ := reader.ReadString('\n')
  fmt.Printf("player 2 (black #), enter your name: ")
  name2, _ := reader.ReadString('\n')
  return_array[0] = name1
  return_array[1] = name2
  return return_array
}

func ask_move(color_ask string) [2]Point2 {
  var return_array [2]Point2
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("%s stuk on colum: ", color_ask)
  move_from_x_str, _ := reader.ReadString('\n')
  fmt.Printf("%s stuk on row: ", color_ask)
  move_from_y_str, _ := reader.ReadString('\n')
  fmt.Print("move to colum: ")
  move_to_x_str, _ := reader.ReadString('\n')
  fmt.Print("move to row: ")
  move_to_y_str, _ := reader.ReadString('\n')
  move_from_x, _ := strconv.Atoi(move_from_x_str[:1])
  move_from_y, _ := strconv.Atoi(move_from_y_str[:1])
  move_to_x, _ := strconv.Atoi(move_to_x_str[:1])
  move_to_y, _ := strconv.Atoi(move_to_y_str[:1])
  //fmt.Println(move_from_x)
  //fmt.Println(move_from_y)
  //fmt.Println(move_to_x)
  //fmt.Println(move_to_y)
  move_from_point := Point2{move_from_x, move_from_y}
  move_to_point := Point2{move_to_x, move_to_y}
  //fmt.Println(move_to_point)
  return_array[0] = move_from_point
  return_array[1] = move_to_point
  return return_array
}

func find_stuk(position Point2, game Game) Stuk {
  for _, value := range game.stukken_white.stukken {
    if value.location == position {
      return value
    }
  }
  for _, value := range game.stukken_black.stukken {
    if value.location == position {
      return value
    }
  }
  fmt.Println("fout stuk")
  return Stuk{}
}

func find_index(position Point2, game Game) int {
  for index, value := range game.stukken_white.stukken {
    if value.location == position {
      return index
    }
  }
  for index, value := range game.stukken_black.stukken {
    if value.location == position {
      return index
    }
  }
  fmt.Println("fout stuk index")
  return 0
}

func check_finished(game Game) bool {
  if len(game.stukken_white.stukken) == 0 {
    return true
  } else if len(game.stukken_black.stukken) == 0 {
    return true
  }
  return false
}

func find_winner(game Game) string {
  return "naam van de winner"
}

func contains(s []Point2, e Point2) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func check_to_upgrade(game Game) []Point2 {
  var return_array []Point2
  for _, value := range game.stukken_white.stukken {
    if value.location.Y == 9 && value.upgraded == false {
      return_array = append(return_array, value.location)
    }
  }
  for _, value := range game.stukken_black.stukken {
    if value.location.Y == 0 && value.upgraded == false {
      return_array = append(return_array, value.location)
    }
  }
  return return_array
}
