package web

import (
	"html/template"
	"net/http"
	"strconv"
	"tic-tac-toe/internal/game"
	"time"
)

var g = game.NewGame()
var mode = "ai" // ai or human
var manager = game.NewManager()
var difficulty = "hard"
var modes = make(map[string]string)

type GamePage struct {
	Board       [9]game.Player
	Message     string
	GameOver    bool
	WinnerCells []int
	AITurn      bool
	Mode        string
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	tmpl, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

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

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		return
// 	}
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	sessionID := getSessionID(w, r)

	g := manager.GetGame(sessionID)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

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

	// if mode == "ai" && g.CurrentPlayer == game.O && !g.GameOver {
	// 	aiTurn = true
	// }

	if mode == "ai" && g.CurrentPlayer == game.O && !g.GameOver {
		sessionID := getSessionID(w, r)

		mode := modes[sessionID]

		if mode == "" {
			mode = "ai"
		}

	}

	data := GamePage{
		Board:       g.Board,
		Message:     message,
		GameOver:    g.GameOver,
		WinnerCells: cells,
		AITurn:      aiTurn,
		Mode:        mode,
	}

	tmpl.Execute(w, data)
}

// func MoveHandler(w http.ResponseWriter, r *http.Request) {

// 	cell := r.URL.Query().Get("cell")

// 	idx, err := strconv.Atoi(cell)

// 	if err != nil {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	err = g.PlayMove(idx)

// 	if err != nil {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	// AI Move
// 	if mode == "ai" && !g.GameOver && g.CurrentPlayer == game.O {

// 		aiMove := g.BestMoveMinimax(game.O, game.X)

// 		if aiMove >= 0 {
// 			g.PlayMove(aiMove)
// 		}
// 	}

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

func MoveHandler(w http.ResponseWriter, r *http.Request) {

	sessionID := getSessionID(w, r)

	g := manager.GetGame(sessionID)

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

	if mode == "ai" && !g.GameOver && g.CurrentPlayer == game.O {

		aiMove := g.BestMove(difficulty, game.O, game.X)

		if aiMove >= 0 {
			g.PlayMove(aiMove)
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// func ResetHandler(w http.ResponseWriter, r *http.Request) {

// 	g.Reset()

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

func ResetHandler(w http.ResponseWriter, r *http.Request) {

	sessionID := getSessionID(w, r)

	g := manager.GetGame(sessionID)

	g.Reset()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// func ModeHandler(w http.ResponseWriter, r *http.Request) {

// 	m := r.URL.Query().Get("mode")

// 	if m == "ai" || m == "human" {
// 		mode = m
// 		g.Reset()
// 	}

//		http.Redirect(w, r, "/", http.StatusSeeOther)
//	}
func ModeHandler(w http.ResponseWriter, r *http.Request) {

	sessionID := getSessionID(w, r)

	m := r.URL.Query().Get("mode")

	if m == "ai" || m == "human" {

		modes[sessionID] = m

		g := manager.GetGame(sessionID)

		g.Reset()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getSessionID(w http.ResponseWriter, r *http.Request) string {

	cookie, err := r.Cookie("session_id")

	if err == nil {
		return cookie.Value
	}

	id := strconv.FormatInt(time.Now().UnixNano(), 10)

	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: id,
		Path:  "/",
	})

	return id
}

func DifficultyHandler(w http.ResponseWriter, r *http.Request) {

	d := r.URL.Query().Get("level")

	if d == "easy" || d == "medium" || d == "hard" {
		difficulty = d
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
