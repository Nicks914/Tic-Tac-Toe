package game

import "math/rand"

// Minimax algorithm
func (g *Game) minimax(depth int, maximizing bool, ai, human Player) int {

	winner, tie, _ := g.CheckWinner()

	if winner == ai {
		return 10 - depth
	}

	if winner == human {
		return depth - 10
	}

	if tie {
		return 0
	}

	if maximizing {

		best := -1000

		for _, move := range g.AvailableMoves() {

			g.MakeMove(move, ai)

			score := g.minimax(depth+1, false, ai, human)

			g.UndoMove(move)

			if score > best {
				best = score
			}
		}

		return best
	}

	best := 1000

	for _, move := range g.AvailableMoves() {

		g.MakeMove(move, human)

		score := g.minimax(depth+1, true, ai, human)

		g.UndoMove(move)

		if score < best {
			best = score
		}
	}

	return best
}

// Best move using minimax
func (g *Game) BestMoveMinimax(ai, human Player) int {

	bestScore := -1000
	bestMove := -1

	for _, move := range g.AvailableMoves() {

		g.MakeMove(move, ai)

		score := g.minimax(0, false, ai, human)

		g.UndoMove(move)

		if score > bestScore {

			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

// Random move
func (g *Game) BestMoveRandom() int {

	moves := g.AvailableMoves()

	if len(moves) == 0 {
		return -1
	}

	return moves[rand.Intn(len(moves))]
}

func (g *Game) BestMove(difficulty string, ai, human Player) int {

	switch difficulty {

	case "easy":
		return g.BestMoveRandom()

	case "medium":

		if rand.Intn(2) == 0 {
			return g.BestMoveRandom()
		}

		return g.BestMoveMinimax(ai, human)

	case "hard":
		return g.BestMoveMinimax(ai, human)

	default:
		return g.BestMoveRandom()
	}
}
