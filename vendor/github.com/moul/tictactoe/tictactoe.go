package tictactoe

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type TicTacToe struct {
	Map    map[int]map[int]string
	Player string
	Size   int
}

type Movement struct {
	Y     int
	X     int
	Score float64
}

func NewTicTacToe() TicTacToe {
	return NewTicTacToeWithSize(3)
}

func NewTicTacToeWithSize(size int) TicTacToe {
	ttt := TicTacToe{
		Size: size,
		Map:  make(map[int]map[int]string, size),
	}
	for i := 0; i < size; i++ {
		ttt.Map[i] = make(map[int]string, size)
	}
	return ttt
}

func (t *TicTacToe) SetPlayer(symbol string) {
	t.Player = symbol
}

func (t *TicTacToe) Set(y, x int, symbol string) {
	t.Map[y][x] = symbol
}

func (t *TicTacToe) ShowMap() string {
	lines := []string{}
	for i := 0; i < t.Size; i++ {
		line := []string{}
		for j := 0; j < t.Size; j++ {
			if t.Map[i][j] == "" {
				line = append(line, ".")
			} else {
				line = append(line, t.Map[i][j])
			}
		}
		lines = append(lines, strings.Join(line, " "))
	}
	return strings.Join(lines, "\n")
}

func (t *TicTacToe) Winner() string {
	// horizontal
	for y := 0; y < t.Size; y++ {
		symbol := t.Map[y][0]
		ok := true
		for x := 1; x < t.Size; x++ {
			if symbol != t.Map[y][x] {
				ok = false
				break
			}
		}
		if ok {
			return symbol
		}
	}

	// vertical
	for x := 0; x < t.Size; x++ {
		symbol := t.Map[0][x]
		ok := true
		for y := 1; y < t.Size; y++ {
			if symbol != t.Map[y][x] {
				ok = false
				break
			}
		}
		if ok {
			return symbol
		}
	}

	// diagonal
	symbol := t.Map[0][0]
	ok := true
	for i := 1; i < t.Size; i++ {
		if t.Map[i][i] != symbol {
			ok = false
			break
		}
	}
	if ok {
		return symbol
	}
	symbol = t.Map[t.Size-1][0]
	ok = true
	for i := 1; i < t.Size; i++ {
		if t.Map[t.Size-1-i][i] != symbol {
			ok = false
			break
		}
	}
	if ok {
		return symbol
	}

	return ""
}

func (t *TicTacToe) AvailableMoves() []Movement {
	moves := []Movement{}
	for y := 0; y < t.Size; y++ {
		for x := 0; x < t.Size; x++ {
			if t.Map[y][x] == "" {
				move := Movement{Y: y, X: x}
				/*
					if y == t.Size/2 && x == t.Size/2 {
						move.Score += 50
					}
				*/
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (t *TicTacToe) ScoreMoves(currentPlayer string, deepness int) []Movement {
	// if map is finished, return nil
	if t.Winner() != "" {
		return nil
	}

	moves := t.AvailableMoves()

	if deepness > 7 {
		// useless to go deeper
		return moves
	}

	value := math.Pow(float64(t.Size*t.Size+1), float64(t.Size*t.Size-deepness))

	for idx, move := range moves {
		t.Set(move.Y, move.X, currentPlayer)
		switch t.Winner() {
		case t.Player:
			moves[idx].Score = value
		case t.Opponent():
			moves[idx].Score = -value
		default:
			for _, subMove := range t.ScoreMoves(t.NextPlayer(currentPlayer), deepness+1) {
				moves[idx].Score += subMove.Score
			}
		}
		t.Set(move.Y, move.X, "")
	}

	return moves
}

func (t *TicTacToe) NextPlayer(currentPlayer string) string {
	switch currentPlayer {
	case "X":
		return "O"
	case "O":
		return "X"
	}
	return ""
}

func (t *TicTacToe) Opponent() string {
	return t.NextPlayer(t.Player)
}

func (t *TicTacToe) Next() (*Movement, error) {
	if t.Winner() != "" {
		return nil, fmt.Errorf("game is already finished")
	}

	// first move is random
	if len(t.AvailableMoves()) == t.Size*t.Size {
		return &Movement{Y: rand.Intn(t.Size), X: rand.Intn(t.Size)}, nil
	}

	moves := t.ScoreMoves(t.Player, 1)

	if len(moves) == 0 {
		return nil, fmt.Errorf("no such move")
	}

	maxIdx := 0
	maxScore := moves[0].Score

	for idx, move := range moves {
		if move.Score > maxScore {
			maxScore = move.Score
			maxIdx = idx
		}
	}
	move := moves[maxIdx]
	// fmt.Println(move, maxScore, maxIdx)
	return &move, nil
}
