package main

type player struct {
	id       int
	score    int
	boardpos int // stored as 0-9 but the real number is 1-10
}

func newPlayer(id int, position int) *player {
	p := player{}
	p.id = id
	p.score = 0
	p.boardpos = position
	return &p
}

func (p *player) move(moves int) {
	p.boardpos += moves
	p.boardpos %= 10
	if p.boardpos == 0 {
		p.boardpos = 10
	}
	p.score += p.boardpos
}

func (p player) pos() int {
	return p.boardpos + 1
}

type playerstate struct {
	pos, score int
}

type gamestate struct {
	players [3]playerstate
}

func newGameState(p1pos, p2pos int) gamestate {
	return gamestate{players: [3]playerstate{{0, 0}, {p1pos, 0}, {p2pos, 0}}}
}

func (gs gamestate) win(win int) bool {
	return gs.players[1].score >= win || gs.players[2].score >= win
}

func move(pos, moves int) int {
	pos += moves
	pos %= 10
	if pos == 0 {
		pos = 10
	}
	return pos
}
