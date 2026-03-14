package game

import "fmt"

type Game struct {
	Board         [9]Player
	CurrentPlayer Player
	GameOver      bool
	Winner        Player
	ScoreX        int
	ScoreO        int
	Draws         int
}

func NewGame() *Game {
	g := &Game{}
	g.Reset()
	return g
}

func (g *Game) Reset() {
	for i := range g.Board {
		g.Board[i] = Empty
	}

	g.CurrentPlayer = X
	g.GameOver = false
	g.Winner = Empty
}

// Print board
func (g *Game) PrintBoard() {

	fmt.Println()

	for r := 0; r < 3; r++ {

		for c := 0; c < 3; c++ {

			fmt.Printf(" %s ", g.Board[r*3+c])

			if c < 2 {
				fmt.Print("|")
			}
		}

		fmt.Print("    ")

		for c := 0; c < 3; c++ {

			fmt.Printf(" %d ", r*3+c+1)

			if c < 2 {
				fmt.Print("|")
			}
		}

		fmt.Println()

		if r < 2 {
			fmt.Println("---+---+---    ---+---+---")
		}
	}

	fmt.Println()
}

// Get available moves
func (g *Game) AvailableMoves() []int {

	var moves []int

	for i := 0; i < 9; i++ {
		if g.Board[i] == Empty {
			moves = append(moves, i)
		}
	}

	return moves
}

// Make move
func (g *Game) MakeMove(index int, player Player) {
	g.Board[index] = player
}

// Undo move
func (g *Game) UndoMove(index int) {
	g.Board[index] = Empty
}

// Check winner
func (g *Game) CheckWinner() (Player, bool, []int) {

	patterns := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, p := range patterns {

		a, b, c := g.Board[p[0]], g.Board[p[1]], g.Board[p[2]]

		if a != Empty && a == b && b == c {
			return a, false, []int{p[0], p[1], p[2]}
		}
	}

	for _, v := range g.Board {
		if v == Empty {
			return Empty, false, nil
		}
	}

	return Empty, true, nil
}

func (g *Game) PlayMove(index int) error {

	if g.GameOver {
		return fmt.Errorf("game already finished")
	}

	if index < 0 || index > 8 {
		return fmt.Errorf("invalid cell")
	}

	if g.Board[index] != Empty {
		return fmt.Errorf("cell already occupied")
	}

	g.Board[index] = g.CurrentPlayer

	winner, tie, _ := g.CheckWinner()

	if winner != Empty {
		g.GameOver = true
		g.Winner = winner
		return nil
	}

	if tie {
		g.GameOver = true
		return nil
	}

	if g.CurrentPlayer == X {
		g.CurrentPlayer = O
	} else {
		g.CurrentPlayer = X
	}

	return nil
}
