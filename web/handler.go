package web

import (
	"html/template"
	"net/http"
	"strconv"

	"tic-tac-toe/game"
)

var g = game.NewGame()
var mode = "ai" // ai or human

type GamePage struct {
	Board       [9]game.Player
	Message     string
	GameOver    bool
	WinnerCells []int
	AITurn      bool
}

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
