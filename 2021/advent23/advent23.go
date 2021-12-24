package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/tools"
	"log"
	"math"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

//globals
type paths map[move]howToMove

var PATHS paths
var stateCache map[string]int

func initialize() {
	PATHS = precomputePaths()
	stateCache = make(map[string]int)
}

func run1(inputText string) int {
	//create initial state
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{4, 3}},
			room{0, []int{4, 3}},
			room{0, []int{1, 2}},
			room{0, []int{1, 2}}},
		corridor: corridor{0, 0, 0, 0, 0, 0, 0},
	}
	cost := step(initialState)

	return cost
}

func run2(inputText string) int {

	return 0
}

// Don't count doors as spaces; they're just costs
// Don't count rooms as spaces; they're just costs (1 move if 0 occupants, else 2 moves)
// Cannot move into a room unless it is empty or contains 1 of its type

// Rooms, not amphipods; they're 1,2,3,4
const (
	A = 0
	B = 1
	C = 2
	D = 3
)

var COSTS = map[int]int{A: 1, B: 10, C: 100, D: 1000}

type move struct {
	corridor int
	room     int
}
type howToMove struct {
	distance   int
	traversals []int
}
type moveWithCost struct {
	corridor int
	room     int
	cost     int
	who      int
}
type room struct {
	finishers int   // number of desired amphipods
	starters  []int // the amphipods that are in this room, in order from corridor out
}

// the list of movable corridor spaces and what is in them
type corridor []int

type state struct {
	cost     int    // total cost of all moves to reach this state
	rooms    []room // 4 rooms with 2 slots each
	corridor corridor
}

// precalculate moving from any corridor to any room
// thanks  https://github.com/davearussell/advent2021/blob/master/day23 for the idea
func precomputePaths() paths {
	p := make(paths)
	p[move{0, A}] = howToMove{3, []int{1}}
	p[move{0, B}] = howToMove{5, []int{1, 2}}
	p[move{0, C}] = howToMove{7, []int{1, 2, 3}}
	p[move{0, D}] = howToMove{9, []int{1, 2, 3, 4}}
	p[move{1, A}] = howToMove{2, []int{}}
	p[move{1, B}] = howToMove{4, []int{2}}
	p[move{1, C}] = howToMove{6, []int{2, 3}}
	p[move{1, D}] = howToMove{8, []int{2, 4}}
	p[move{2, A}] = howToMove{2, []int{}}
	p[move{2, B}] = howToMove{2, []int{}}
	p[move{2, C}] = howToMove{4, []int{3}}
	p[move{2, D}] = howToMove{6, []int{3, 4}}
	p[move{3, A}] = howToMove{4, []int{2}}
	p[move{3, B}] = howToMove{2, []int{}}
	p[move{3, C}] = howToMove{2, []int{}}
	p[move{3, D}] = howToMove{4, []int{4}}
	p[move{4, A}] = howToMove{6, []int{2, 3}}
	p[move{4, B}] = howToMove{4, []int{3}}
	p[move{4, C}] = howToMove{2, []int{}}
	p[move{4, D}] = howToMove{2, []int{}}
	p[move{5, A}] = howToMove{8, []int{2, 3, 4}}
	p[move{5, B}] = howToMove{6, []int{3, 4}}
	p[move{5, C}] = howToMove{4, []int{4}}
	p[move{5, D}] = howToMove{2, []int{}}
	p[move{6, A}] = howToMove{9, []int{2, 3, 4, 5}}
	p[move{6, B}] = howToMove{7, []int{3, 4, 5}}
	p[move{6, C}] = howToMove{5, []int{4, 5}}
	p[move{6, D}] = howToMove{3, []int{5}}
	return p
}

func findMovers() {
	// return a list of amphipods who can move at all
	// first ones that can move to their room
	// then ones in rooms that can move out
}

func begin(s *state) (cost int) {
	initialize()
	s.fixStartingInFinishPosition()
	return step(s)
}

func step(s *state) (cost int) {
	if s.isWinner() {
		//log.Printf("complete! cost = %d", s.cost)
		return s.cost
	}

	// find moves into a room
	corridorToRoomMoves := s.findMoversInCorridor()
	roomToCorridorMoves := s.findMoversInRooms()
	// FIXME sort the lists of moves by cost
	cost = math.MaxInt
	for _, m := range corridorToRoomMoves {
		// make a state for the move, and recurse on it
		// this means copying the reference types in state
		// add the cost of the move to the state now
		newState := s.copy()
		newState.corridor[m.corridor] = -1
		newState.rooms[m.room].finishers++
		newState.cost += m.cost

		// Store states and cost globally, and if we look at same state but higher cost, don't recurse it; it's already a loser.
		cachedState, exists := stateCache[newState.toKey()]
		if exists && cachedState < newState.cost {
			// we've been here already for less, so stop looking down this chain
			//log.Printf("been there done that")
			break
		}
		stateCache[newState.toKey()] = newState.cost
		//log.Printf("Moving %d from corridor %d to room %d with cost %d", m.who, m.corridor, m.room, m.cost)
		cost = tools.MinInt(cost, step(newState))
	}

	for _, m := range roomToCorridorMoves {
		// make a state for the move, and recurse on it
		// this means copying the reference types in state
		// add the cost of the move to the state now
		newState := s.copy()
		newState.corridor[m.corridor] = m.who
		// take the front starter out of the room
		newState.rooms[m.room].starters = newState.rooms[m.room].starters[1:]
		newState.cost += m.cost

		// Store states and cost globally, and if we look at same state but higher cost, don't recurse it; it's already a loser.
		cachedState, exists := stateCache[newState.toKey()]
		if exists && cachedState < newState.cost {
			//log.Printf("been there done that")
			// we've been here already for less, so stop looking down this chain
			break
		}
		stateCache[newState.toKey()] = newState.cost
		//log.Printf("Moving %d from room %d to corridor %d with cost %d", m.who, m.room, m.corridor, m.cost)
		cost = tools.MinInt(cost, step(newState))
	}

	return cost
}

// find all positions they can move to
// make a state for each one of them

// when a recursive call returns a cost, return min(mycost, itscost)

// returns a list of corridor spots that contain an amphipod that can move

func (s *state) findMoversInCorridor() []moveWithCost {
	moves := []moveWithCost{}
	for i, a := range s.corridor {

		if a < 0 {
			continue // there's nobody there
		}

		// a wants to go to the room of its number
		// is the room open?
		if !s.rooms[a].isClear() {
			continue
		}

		// is the path there clear?
		path, ok := PATHS[move{i, a}]
		if !ok {
			log.Panicf("can't find a path for %v", move{i, a})
		}
		pathClear := true
		for _, traversal := range path.traversals {
			if s.corridor[traversal] >= 0 {
				pathClear = false
				break
			}
		}
		if !pathClear {
			continue
		}
		cost := path.distance * COSTS[a]
		if s.rooms[a].finishers == 0 {
			cost += COSTS[a]
		}
		moves = append(moves, moveWithCost{room: a, corridor: i, cost: cost, who: a})
	}
	return moves
}

// returns list of moves from rooms into the hall
func (s *state) findMoversInRooms() []moveWithCost {
	moves := []moveWithCost{}

	for ir, r := range s.rooms {

		if r.finishers == 2 { // FIXME this will change for step 2; make it a constant
			continue
		}
		if len(r.starters) == 0 {
			continue
		}
		a := r.starters[0]
		// make a list of ALL hallway positions it can go
		// this would be a good place to be smart and filter out useless moves
		for ic, c := range s.corridor {
			if c >= 0 { // can't go there, it's occupied
				continue
			}
			// is the path there clear?
			path, ok := PATHS[move{ic, ir}]
			if !ok {
				log.Panicf("can't find a path for %v", move{ic, a})
			}
			pathClear := true
			for _, traversal := range path.traversals {
				if s.corridor[traversal] >= 0 {
					pathClear = false
					break
				}
			}
			if !pathClear {
				continue
			}
			cost := path.distance * COSTS[a]
			// if moving from the back of the room, add one
			// FIXME breaking point for size 4 rooms
			if len(s.rooms[ir].starters)+s.rooms[ir].finishers == 1 {
				cost += COSTS[a]
			}
			moves = append(moves, moveWithCost{
				corridor: ic,
				room:     ir,
				cost:     cost,
				who:      a,
			})
		}
	}

	return moves
}

// can an amphipod go home, or is this room blocked by strangers?
func (r *room) isClear() bool {
	return len(r.starters) == 0
}

func (s *state) isWinner() bool {
	for _, r := range s.rooms {
		if r.finishers < 2 { // FIXME update for step 2
			return false
		}
	}
	return true
}

func (s *state) copy() *state {
	s2 := state{cost: s.cost}
	s2.corridor = make(corridor, 7)
	// copy won't extend the target slice
	copy(s2.corridor, s.corridor)
	s2.rooms = make([]room, 4)
	copy(s2.rooms, s.rooms)
	return &s2
}

func (s *state) toKey() string {
	return fmt.Sprintf("%v/%v", s.corridor, s.rooms)
}

func (m move) cost(a int) int {
	return PATHS[m].distance * m.room * COSTS[a]
}

func (s *state) fixStartingInFinishPosition() {
	// no input we have has the case where an amphipod is in its room but blocks another, so ignore that case

	for i, r := range s.rooms {
		updatedStarters := []int{}
		for _, starter := range r.starters {
			if starter == i {
				r.finishers++
			} else {
				updatedStarters = append(updatedStarters, starter)
			}
		}
		r.starters = updatedStarters
		s.rooms[i] = r
	}
}

/*
step- get a state
make a list of all posssible moves
create a state for each one
call(step) on each of those
if it wins, stop and return cost
make sure you can't go back and forth forever



*/
