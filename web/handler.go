package web

import (
	"html/template"
	"net/http"
	"strconv"

	"tic-tac-toe/game"
)

var g = game.NewGame()
var mode = "ai" // ai or human
// func HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))

// 	// data := map[string]interface{}{
// 	// 	"Board": g.Board,
// 	// }

// 	data := GamePage{
// 		Board:   g.Board ,
// 		Message:  "Player X Wins!",
// 		GameOver: true,
// 	}
// 	tmpl.Execute(w, data)
// }

// func MoveHandler(w http.ResponseWriter, r *http.Request) {

// 	cell := r.URL.Query().Get("cell")

// 	idx, _ := strconv.Atoi(cell)

// 	if g.Board[idx] == game.Empty {

// 		g.MakeMove(idx, game.X)

// 		ai := g.BestMoveRandom()

// 		if ai >= 0 {
// 			g.MakeMove(ai, game.O)
// 		}
// 	}

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

type GamePage struct {
	Board       [9]game.Player
	Message     string
	GameOver    bool
	WinnerCells []int
	AITurn      bool
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))

// 	winner, tie := g.CheckWinner()

// 	message := ""
// 	gameOver := false

// 	if winner != game.Empty {
// 		message = "Player " + string(winner) + " Wins!"
// 		gameOver = true
// 	}

// 	if tie {
// 		message = "It's a Draw!"
// 		gameOver = true
// 	}

// 	data := GamePage{
// 		Board:    g.Board,
// 		Message:  message,
// 		GameOver: gameOver,
// 	}

// 	tmpl.Execute(w, data)
// }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))

// 	message := ""

// 	if g.GameOver {

// 		if g.Winner != game.Empty {
// 			message = "Player " + g.Winner.String() + " Wins!"
// 		} else {
// 			message = "Game Draw!"
// 		}

// 	} else {
// 		message = "Turn: " + g.CurrentPlayer.String()
// 	}

// 	data := map[string]interface{}{
// 		"Board":   g.Board,
// 		"Message": message,
// 		"Mode":    mode,
// 	}

// 	tmpl.Execute(w, data)
// }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))

// 	winner, tie, cells := g.CheckWinner()

// 	message := ""

// 	if winner != game.Empty {
// 		message = "Player " + winner.String() + " Wins!"
// 		g.GameOver = true
// 		g.Winner = winner
// 	} else if tie {
// 		message = "Game Draw!"
// 		g.GameOver = true
// 	} else {
// 		message = "Turn: " + g.CurrentPlayer.String()
// 	}

// 	// detect AI thinking state
// 	aiTurn := false
// 	if mode == "ai" && g.CurrentPlayer == game.O && !g.GameOver {
// 		aiTurn = true
// 	}

// 	data := GamePage{
// 		Board:       g.Board,
// 		Message:     message,
// 		GameOver:    g.GameOver,
// 		WinnerCells: cells,
// 		AITurn:      aiTurn,
// 	}

//		err := tmpl.Execute(w, data)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	winner, tie, cells := g.CheckWinner()

	message := ""

	if winner != game.Empty {
		message = "Player " + winner.String() + " Wins!"
		g.GameOver = true
		g.Winner = winner
	} else if tie {
		message = "Game Draw!"
		g.GameOver = true
	} else {
		message = "Turn: " + g.CurrentPlayer.String()
	}

	aiTurn := false
	if mode == "ai" && g.CurrentPlayer == game.O && !g.GameOver {
		aiTurn = true
	}

	data := GamePage{
		Board:       g.Board,
		Message:     message,
		GameOver:    g.GameOver,
		WinnerCells: cells,
		AITurn:      aiTurn,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
func MoveHandler(w http.ResponseWriter, r *http.Request) {

	cell := r.URL.Query().Get("cell")

	idx, err := strconv.Atoi(cell)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = g.PlayMove(idx)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// AI Move
	if mode == "ai" && !g.GameOver && g.CurrentPlayer == game.O {

		aiMove := g.BestMoveMinimax(game.O, game.X)

		if aiMove >= 0 {
			g.PlayMove(aiMove)
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {

	g.Reset()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ModeHandler(w http.ResponseWriter, r *http.Request) {

	m := r.URL.Query().Get("mode")

	if m == "ai" || m == "human" {
		mode = m
		g.Reset()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
