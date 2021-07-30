package session

import "github.com/df-mc/dragonfly/server/player"

var sessions = map[string]*Session{}

func New(p *player.Player) *Session {
	session := Session{Player: p}.DefaultFlags()
	sessions[p.Name()] = session
	return session
}

func Get(p *player.Player) *Session {
	return sessions[p.Name()]
}

func (s Session) Close() {
	delete(sessions, s.Player.Name())
}