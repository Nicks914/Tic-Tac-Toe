package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Simple Tic-Tac-Toe implementation with two AI modes:
// - "minimax" (optimal / unbeatable)
// - "easy" (random moves)
// Supports Human vs AI, Human vs Human, AI vs AI.

type Player int

const (
	Empty Player = iota
	X
	O
)

func (p Player) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

type Game struct {
	board [9]Player
}

// printBoard prints indices and current board side-by-side for guidance.
func (g *Game) printBoard() {
	fmt.Println()
	for r := 0; r < 3; r++ {
		// row values
		for c := 0; c < 3; c++ {
			fmt.Printf(" %s ", g.board[r*3+c])
			if c < 2 {
				fmt.Print("|")
			}
		}
		fmt.Print("    ")
		// indices
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

// availableMoves returns list of indices (0-based) that are empty.
func (g *Game) availableMoves() []int {
	moves := make([]int, 0, 9)
	for i := 0; i < 9; i++ {
		if g.board[i] == Empty {
			moves = append(moves, i)
		}
	}
	return moves
}

// makeMove sets a move for a player on position idx (0-based).
func (g *Game) makeMove(idx int, p Player) {
	g.board[idx] = p
}

// undoMove sets idx back to Empty.
func (g *Game) undoMove(idx int) {
	g.board[idx] = Empty
}

// checkWinner returns the winner (X or O) or Empty if none; returns true if tie as second return
func (g *Game) checkWinner() (Player, bool) {
	winning := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	for _, w := range winning {
		a, b, c := g.board[w[0]], g.board[w[1]], g.board[w[2]]
		if a != Empty && a == b && b == c {
			return a, false
		}
	}
	// tie?
	for _, s := range g.board {
		if s == Empty {
			return Empty, false // no winner, not tie (game continues)
		}
	}
	return Empty, true // tie
}

// minimax implementation: returns score for current board for 'maximizingPlayer'
func (g *Game) minimax(depth int, maximizing bool, aiPlayer, humanPlayer Player) int {
	winner, tie := g.checkWinner()
	if winner == aiPlayer {
		return 10 - depth // prefer quicker win
	}
	if winner == humanPlayer {
		return depth - 10 // prefer slower loss
	}
	if tie {
		return 0
	}

	if maximizing {
		best := -1000
		for _, mv := range g.availableMoves() {
			g.makeMove(mv, aiPlayer)
			score := g.minimax(depth+1, false, aiPlayer, humanPlayer)
			g.undoMove(mv)
			if score > best {
				best = score
			}
		}
		return best
	} else {
		best := 1000
		for _, mv := range g.availableMoves() {
			g.makeMove(mv, humanPlayer)
			score := g.minimax(depth+1, true, aiPlayer, humanPlayer)
			g.undoMove(mv)
			if score < best {
				best = score
			}
		}
		return best
	}
}

// bestMoveMinimax chooses the optimal move for aiPlayer against humanPlayer.
func (g *Game) bestMoveMinimax(aiPlayer, humanPlayer Player) int {
	bestScore := -10000
	bestIdx := -1
	for _, mv := range g.availableMoves() {
		g.makeMove(mv, aiPlayer)
		score := g.minimax(0, false, aiPlayer, humanPlayer)
		g.undoMove(mv)
		if score > bestScore {
			bestScore = score
			bestIdx = mv
		}
	}
	// fallback
	if bestIdx == -1 {
		avs := g.availableMoves()
		if len(avs) > 0 {
			return avs[rand.Intn(len(avs))]
		}
	}
	return bestIdx
}

// bestMoveRandom selects a random available index.
func (g *Game) bestMoveRandom() int {
	avs := g.availableMoves()
	if len(avs) == 0 {
		return -1
	}
	return avs[rand.Intn(len(avs))]
}

// readIntFromStdin reads an int from stdin with prompt.
func readIntFromStdin(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		n, err := strconv.Atoi(text)
		if err == nil {
			return n
		}
		fmt.Println("Please enter a valid number.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var g Game

	fmt.Println("=== Tic-Tac-Toe (Go) ===")
	fmt.Println("Modes:")
	fmt.Println("1) Human vs AI")
	fmt.Println("2) Human vs Human")
	fmt.Println("3) AI vs AI")
	mode := readIntFromStdin("Select mode (1-3): ")
	for mode < 1 || mode > 3 {
		mode = readIntFromStdin("Select mode (1-3): ")
	}

	var aiMode string
	if mode == 1 || mode == 3 {
		fmt.Println("AI difficulty options:")
		fmt.Println("1) minimax (unbeatable)")
		fmt.Println("2) easy (random)")
		d := readIntFromStdin("Choose AI mode (1-2): ")
		for d < 1 || d > 2 {
			d = readIntFromStdin("Choose AI mode (1-2): ")
		}
		if d == 1 {
			aiMode = "minimax"
		} else {
			aiMode = "easy"
		}
	}

	var humanIsX bool = true
	// Who is X and who starts?
	if mode == 1 {
		fmt.Println("Do you want to play as X (goes first) or O?")
		fmt.Println("1) Play as X (first)")
		fmt.Println("2) Play as O (second)")
		ch := readIntFromStdin("Choose (1-2): ")
		for ch != 1 && ch != 2 {
			ch = readIntFromStdin("Choose (1-2): ")
		}
		humanIsX = (ch == 1)
	}

	// Reset board
	for i := range g.board {
		g.board[i] = Empty
	}

	current := X // X starts always
	if mode == 1 {
		// if human chose O, AI starts
		if !humanIsX {
			current = X
		} else {
			current = X
		}
	}

	// Game loop
	for {
		g.printBoard()
		winner, tie := g.checkWinner()
		if winner != Empty || tie {
			if tie {
				fmt.Println("It's a tie!")
			} else {
				fmt.Printf("Winner: %s\n", winner)
			}
			break
		}

		switch mode {
		case 1: // Human vs AI
			var humanPlayer, aiPlayer Player
			if humanIsX {
				humanPlayer = X
				aiPlayer = O
			} else {
				humanPlayer = O
				aiPlayer = X
			}

			if current == humanPlayer {
				// human turn
				for {
					move := readIntFromStdin(fmt.Sprintf("Your move (%s). Choose cell 1-9: ", humanPlayer))
					if move < 1 || move > 9 {
						fmt.Println("Invalid cell. Choose 1 to 9.")
						continue
					}
					idx := move - 1
					if g.board[idx] != Empty {
						fmt.Println("Cell already taken. Choose another.")
						continue
					}
					g.makeMove(idx, humanPlayer)
					break
				}
			} else {
				// AI turn
				fmt.Printf("AI (%s) is thinking...\n", aiPlayer)
				var aiMove int
				if aiMode == "minimax" {
					aiMove = g.bestMoveMinimax(aiPlayer, humanPlayer)
				} else {
					aiMove = g.bestMoveRandom()
				}
				if aiMove == -1 {
					fmt.Println("No moves left.")
					break
				}
				fmt.Printf("AI plays at %d\n", aiMove+1)
				g.makeMove(aiMove, aiPlayer)
			}
		case 2: // Human vs Human
			fmt.Printf("Player %s turn. Enter cell 1-9: ", current)
			move := readIntFromStdin("")
			if move < 1 || move > 9 {
				fmt.Println("Invalid cell. Skipping turn.")
			} else {
				idx := move - 1
				if g.board[idx] == Empty {
					g.makeMove(idx, current)
				} else {
					fmt.Println("Cell taken. Try next turn.")
				}
			}
		case 3: // AI vs AI
			var p1, p2 Player = X, O
			var m1, m2 int
			if aiMode == "minimax" {
				m1 = g.bestMoveMinimax(p1, p2)
				m2 = g.bestMoveMinimax(p2, p1)
			} else {
				m1 = g.bestMoveRandom()
				m2 = g.bestMoveRandom()
			}
			var chosen int
			if current == p1 {
				chosen = m1
			} else {
				chosen = m2
			}
			if chosen == -1 {
				fmt.Println("No moves left.")
				break
			}
			fmt.Printf("AI (%s) plays at %d\n", current, chosen+1)
			g.makeMove(chosen, current)
			// small delay for visibility
			time.Sleep(400 * time.Millisecond)
		}

		// swap turn
		if current == X {
			current = O
		} else {
			current = X
		}
	}

	g.printBoard()
	fmt.Println("Game over. Thanks for playing!")
}
