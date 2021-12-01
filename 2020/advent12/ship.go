package advent12

type ship struct {
	facing int
	x int
	y int
}

func (s *ship) rotate(isRight bool, degrees int) {

	if !isRight {
		// turn a left turn into a right turn
		degrees = 360-degrees
	}
	s.facing += degrees / 90
	s.facing %= 4

}

func (s *ship) moveForward(distance int) {
	s.move(s.facing, distance)
}

func (s *ship) move(direction int, distance int) {
	switch direction {
	case N:
		s.y += distance
	case E:
		s.x += distance
	case S:
		s.y -= distance
	case W:
		s.x -= distance
	}
}

func (s *ship) moveToWaypoint(wp *waypoint, times int) {
	s.x = s.x + wp.x * times
	s.y = s.y + wp.y * times
}

