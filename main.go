// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"

// 	"tic-tac-toe/game"
// 	"tic-tac-toe/utils"
// )

// func main() {

// 	rand.Seed(time.Now().UnixNano())

// 	g := game.NewGame()
// 	g.Reset()

// 	fmt.Println("=== Tic-Tac-Toe (Go) ===")

// 	fmt.Println("Modes:")
// 	fmt.Println("1) Human vs AI")
// 	fmt.Println("2) Human vs Human")
// 	fmt.Println("3) AI vs AI")

// 	mode := utils.ReadInt("Select mode (1-3): ")

// 	for mode < 1 || mode > 3 {
// 		mode = utils.ReadInt("Select mode (1-3): ")
// 	}

// 	var aiMode string
// 	var humanPlayer, aiPlayer game.Player

// 	if mode == 1 || mode == 3 {

// 		fmt.Println("AI difficulty options:")
// 		fmt.Println("1) minimax (unbeatable)")
// 		fmt.Println("2) easy (random)")

// 		d := utils.ReadInt("Choose AI mode (1-2): ")

// 		for d < 1 || d > 2 {
// 			d = utils.ReadInt("Choose AI mode (1-2): ")
// 		}

// 		if d == 1 {
// 			aiMode = "minimax"
// 		} else {
// 			aiMode = "easy"
// 		}
// 	}

// 	humanIsX := true

// 	if mode == 1 {

// 		fmt.Println("Do you want to play as X (goes first) or O?")
// 		fmt.Println("1) Play as X (first)")
// 		fmt.Println("2) Play as O (second)")

// 		ch := utils.ReadInt("Choose (1-2): ")

// 		for ch != 1 && ch != 2 {
// 			ch = utils.ReadInt("Choose (1-2): ")
// 		}

// 		humanIsX = (ch == 1)

// 		if humanIsX {
// 			humanPlayer = game.X
// 			aiPlayer = game.O
// 		} else {
// 			humanPlayer = game.O
// 			aiPlayer = game.X
// 		}
// 	}

// 	current := game.X

// 	for {

// 		g.PrintBoard()

// 		winner, tie := g.CheckWinner()

// 		if winner != game.Empty || tie {

// 			fmt.Println()

// 			if tie {

// 				fmt.Println("🤝 It's a Draw!")
// 				fmt.Println("Great game! No one wins this round.")

// 			} else {

// 				fmt.Println("🎉 Game Finished!")

// 				if mode == 1 {

// 					if winner == humanPlayer {

// 						fmt.Println("🏆 Congratulations! You won the game!")

// 					} else {

// 						fmt.Println("🤖 AI Wins!")
// 						fmt.Println("Better luck next time!")

// 					}

// 				} else if mode == 2 {

// 					fmt.Printf("🏆 Player %s wins the game!\n", winner)

// 				} else {

// 					fmt.Printf("🤖 AI %s wins the match!\n", winner)

// 				}
// 			}

// 			break
// 		}

// 		switch mode {

// 		case 1: // Human vs AI

// 			if current == humanPlayer {

// 				move := utils.ReadInt(fmt.Sprintf("🎮 Your move (%s) (1-9): ", humanPlayer))

// 				if move < 1 || move > 9 {
// 					fmt.Println("❌ Invalid cell")
// 					continue
// 				}

// 				idx := move - 1

// 				if g.Board[idx] != game.Empty {
// 					fmt.Println("⚠️ Cell already taken")
// 					continue
// 				}

// 				g.MakeMove(idx, humanPlayer)

// 			} else {

// 				fmt.Println("🤖 AI is thinking...")

// 				var aiMove int

// 				if aiMode == "minimax" {
// 					aiMove = g.BestMoveMinimax(aiPlayer, humanPlayer)
// 				} else {
// 					aiMove = g.BestMoveRandom()
// 				}

// 				fmt.Println("🤖 AI plays:", aiMove+1)

// 				g.MakeMove(aiMove, aiPlayer)
// 			}

// 		case 2: // Human vs Human

// 			move := utils.ReadInt(fmt.Sprintf("🎮 Player %s move (1-9): ", current))

// 			if move < 1 || move > 9 {
// 				fmt.Println("❌ Invalid move")
// 				continue
// 			}

// 			idx := move - 1

// 			if g.Board[idx] != game.Empty {
// 				fmt.Println("⚠️ Cell already taken")
// 				continue
// 			}

// 			g.MakeMove(idx, current)

// 		case 3: // AI vs AI

// 			var move int

// 			if aiMode == "minimax" {

// 				if current == game.X {
// 					move = g.BestMoveMinimax(game.X, game.O)
// 				} else {
// 					move = g.BestMoveMinimax(game.O, game.X)
// 				}

// 			} else {

// 				move = g.BestMoveRandom()
// 			}

// 			fmt.Println("🤖 AI", current, "plays:", move+1)

// 			g.MakeMove(move, current)

// 			time.Sleep(500 * time.Millisecond)
// 		}

// 		if current == game.X {
// 			current = game.O
// 		} else {
// 			current = game.X
// 		}
// 	}

// 	g.PrintBoard()

// 	fmt.Println("🏁 Game Over. Thanks for playing!")
// }

package main

import (
	"fmt"
	"net/http"

	"tic-tac-toe/web"
)

func main() {

	fmt.Println("Server started at http://localhost:8080")

	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/move", web.MoveHandler)
	http.HandleFunc("/reset", web.ResetHandler)
	http.HandleFunc("/mode", web.ModeHandler)

	http.ListenAndServe(":8080", nil)
}
