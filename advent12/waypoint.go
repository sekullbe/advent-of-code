package advent12

// maybe i should make waypoint and ship dependencies, but meh.

type waypoint struct {
	x int
	y int
}

// rotate the waypoint around the ship
// i.e. if waypoint is (2,1), rotating 90deg is (1,-2)
func (w *waypoint) rotate(isRight bool, degrees int) {

	if !isRight {
		// turn a left turn into a right turn
		degrees = 360-degrees
	}
	jumps := degrees / 90
	switch jumps {
	// R (x,y) -> (y,-x) -> (-x,-y) -> (-y,x)
	// L (x,y) -> (-y,x) -> (-x,-y) -> (y,-x)
	case 1:
		w.x, w.y =  w.y, -1*w.x
	case 2:
		w.x, w.y = -1*w.x, -1*w.y
	case 3:
		w.x, w.y = -1*w.y, w.x
	}
}

func (w *waypoint) move(direction int, distance int) {
	switch direction {
	case N:
		w.y += distance
	case E:
		w.x += distance
	case S:
		w.y -= distance
	case W:
		w.x -= distance
	}
}
