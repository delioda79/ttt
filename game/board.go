package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

const (
	none uint8 = iota
	Nought
	Cross
)

// Board represents a board for tic tac toe
type Board interface {
	CheckWin() (bool, uint8)
	AddNought(x, y int) error
	AddCross(x, y int) error
	GetStatus() map[int]map[int]uint8
	LogStatus()
	Run(ch chan []byte)
	Reset()
}

// NoughtCross is a TicTacToe standard board game
type NoughtCross struct {
	board  map[int]map[int]uint8
	winner uint8
	player uint8
}

func (nc *NoughtCross) checkHor() (bool, uint8) {
ROWS:
	for r := 0; r < len(nc.board); r++ {
		var plyr uint8
		if r == 2 {
			log.Println("Bottom", nc.board[r])
		}
	COLS:
		for c := 0; c < len(nc.board[r]); c++ {
			if nc.board[r][c] == 0 {
				if r == 2 {
					log.Println("Bottom skipping row", nc.board[r])
				}
				continue ROWS
			}
			if nc.board[r][c] != 0 && plyr == 0 {
				if r == 2 {
					log.Println("Bottom continue col found", nc.board[r])
				}
				plyr = nc.board[r][c]
				continue COLS
			}

			if plyr != nc.board[r][c] {
				if r == 2 {
					log.Println("Bottom skipping row second", nc.board[r])
				}
				continue ROWS
			}
		}
		if plyr != 0 {
			fmt.Println("Hor win", plyr, nc.board)
			return true, plyr
		}
	}
	return false, 0
}

func (nc *NoughtCross) checkVert() (bool, uint8) {
COLS:
	for c := 0; c < len(nc.board); c++ {
		var plyr uint8
	ROWS:
		for r := 0; r < len(nc.board); r++ {
			if nc.board[r][c] == 0 {
				continue COLS
			}
			if nc.board[r][c] != 0 && plyr == 0 {
				plyr = nc.board[r][c]
				continue ROWS
			}

			if plyr != nc.board[r][c] {
				continue COLS
			}
		}
		if plyr != 0 {
			return true, plyr
		}
	}
	return false, 0
}

func (nc *NoughtCross) checkDiag() (bool, uint8) {
	var plyr uint8
	for r := 0; r < len(nc.board); r++ {
		if nc.board[r][r] == 0 {
			break
		}
		if nc.board[r][r] != 0 && plyr == 0 {
			plyr = nc.board[r][r]
			continue
		}

		if plyr != nc.board[r][r] {
			break
		}
	}

	if plyr != none {
		return true, plyr
	}

	for r := 0; r < len(nc.board); r++ {
		if nc.board[r][len(nc.board)-r-1] == 0 {
			return false, 0
		}
		if nc.board[r][len(nc.board)-r-1] != 0 && plyr == 0 {
			plyr = nc.board[r][r]
			continue
		}

		if plyr != nc.board[r][len(nc.board)-r-1] {
			return false, 0
		}
	}
	if plyr != none {
		return true, plyr
	}
	return false, none
}

// CheckWin checks if somebody has won
func (nc *NoughtCross) CheckWin() (bool, uint8) {
	win, plr := nc.checkHor()
	if win {
		return win, plr
	}

	win, plr = nc.checkVert()
	if win {
		return win, plr
	}
	return nc.checkDiag()
}

// AddNought adds a nought to the board
func (nc *NoughtCross) AddNought(r, c int) error {
	if nc.winner != none {
		return errors.New("Game finished")
	}
	if nc.player == Cross {
		return errors.New("Not your turn")
	}
	if r >= 0 && r < len(nc.board) && c >= 0 && c < len(nc.board) {
		if nc.board[r][c] == 0 {
			nc.board[r][c] = Nought
			log.Printf("Added %d %d %d and map is\n", r, c, nc.board[r][c])
			nc.LogStatus()
			won, winner := nc.CheckWin()
			if won {
				nc.winner = winner
			}
			nc.player = Cross
			return nil
		}
		log.Println("the ", r, c, " is ", nc.board[r][c])
	}

	msg := fmt.Sprintf(
		"Position not allowed %v %v %v %v %v %v %v %v",
		r, c, len(nc.board),
		r < len(nc.board),
		c < len(nc.board),
		r < len(nc.board) && c < len(nc.board),
		r >= 0,
		c >= 0,
	)
	return errors.New(msg)
}

// AddCross adds a cross to the board
func (nc *NoughtCross) AddCross(r, c int) error {
	if nc.winner != none {
		return errors.New("Game finished")
	}
	if nc.player == Nought {
		return errors.New("Not your turn")
	}
	if r < len(nc.board) && c < len(nc.board) {
		if nc.board[r][c] == 0 {
			nc.board[r][c] = Cross
			log.Printf("Added %d %d %d and map is\n", r, c, nc.board[r][c])
			nc.LogStatus()
			won, winner := nc.CheckWin()
			if won {
				nc.winner = winner
			}
			nc.player = Nought
			return nil
		}
	}

	msg := fmt.Sprintf("Position not allowed %v %v %v %v %v %v", r, c, len(nc.board), r < len(nc.board), c < len(nc.board), r < len(nc.board) && c < len(nc.board))
	return errors.New(msg)
}

// LogStatus returns teh current game status
func (nc *NoughtCross) LogStatus() {
	fmt.Println()
	for r := 0; r < len(nc.board); r++ {
		for c := 0; c < len(nc.board[r]); c++ {
			fmt.Printf("%d", nc.board[r][c])
		}
		fmt.Println()
	}
}

// GetStatus returns teh current game status
func (nc NoughtCross) GetStatus() map[int]map[int]uint8 {
	return nc.board
}

// Run runs the game
func (nc *NoughtCross) Run(ch chan []byte) {
	for {
		mv := &Move{}
		json.Unmarshal(<-ch, mv)
		fmt.Printf("Move %+v\n", mv)
		var err error
		if mv.Player == Nought {
			err = nc.AddNought(mv.Y, mv.X)
		} else {
			err = nc.AddCross(mv.Y, mv.X)
		}

		if err != nil {
			ch <- []byte(err.Error())
		} else {
			sts := nc.GetStatus()
			bts, err := json.Marshal(sts)
			if err != nil {
				ch <- []byte(err.Error())
			} else {
				ch <- bts
			}
		}
	}
}

// Run runs the game
func (nc *NoughtCross) Reset() {
	nc.board = newBoard()
	nc.winner = none
	nc.player = Cross
}

// Run runs the game
func newBoard() map[int]map[int]uint8 {
	board := map[int]map[int]uint8{}
	for r := 0; r < 3; r++ {
		board[r] = map[int]uint8{}
		for c := 0; c < 3; c++ {
			board[r][c] = none
		}
	}

	return board
}

// NewNoughtCross returns a new nought and crosses game
func NewNoughtCross() Board {

	return &NoughtCross{board: newBoard(), player: Cross}
}
