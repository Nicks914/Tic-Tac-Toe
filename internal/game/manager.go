package game

import (
	"sync"
)

type Manager struct {
	Games map[string]*Game
	Mu    sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Games: make(map[string]*Game),
	}
}

func (m *Manager) GetGame(sessionID string) *Game {

	m.Mu.Lock()
	defer m.Mu.Unlock()

	g, ok := m.Games[sessionID]

	if !ok {
		g = NewGame()
		m.Games[sessionID] = g
	}

	return g
}
