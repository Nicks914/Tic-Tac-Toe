package game

// Player represents a game player
type Player int

const (
	Empty Player = iota
	X
	O
)

// Convert player to string
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
