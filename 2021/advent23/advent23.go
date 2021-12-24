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
	//fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

//globals
type paths map[move]howToMove

var LOG bool = false
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
			room{0, []int{D, C}},
			room{0, []int{D, C}},
			room{0, []int{A, B}},
			room{0, []int{A, B}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}

	return begin(initialState)
}

func run2(inputText string) int {
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{D, D, D, C}},
			room{0, []int{D, C, B, C}},
			room{0, []int{A, B, A, B}},
			room{0, []int{A, A, C, B}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	return beginWithDepth(initialState, 4)
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

var ROOMDEPTH = 2

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
	cost        int    // total cost of all moves to reach this state
	rooms       []room // 4 rooms with N slots each
	corridor    corridor
	breadcrumbs []moveWithCost
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

func beginWithDepth(s *state, d int) (cost int) {
	ROOMDEPTH = d
	return begin(s)
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
	// TODO sort the lists of moves by cost, will that help?
	cost = math.MaxInt
	for _, m := range corridorToRoomMoves {
		// make a state for the move, and recurse on it
		// this means copying the reference types in state
		// add the cost of the move to the state now
		newState := s.copy()
		newState.corridor[m.corridor] = -1
		newState.rooms[m.room].finishers++
		newState.cost += m.cost
		newState.breadcrumbs = append(newState.breadcrumbs, m)

		// Store states and cost globally, and if we look at same state but higher cost, don't recurse it; it's already a loser.
		cachedState, exists := stateCache[newState.toKey()]
		if exists && cachedState <= newState.cost {
			// we've been here already for less, so stop looking down this chain
			//log.Printf("been there done that")
			break
		}
		if LOG && newState.isWinner() {
			log.Printf("best winner seen yet. cost=%d", newState.cost)
			for _, bc := range newState.breadcrumbs {
				log.Println(bc.toString())
			}
			log.Println("---")
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
		newState.breadcrumbs = append(newState.breadcrumbs, m)

		// Fuuuuuunky. Trimming room-to-corridor moves with the *same* state and cost causes it to miss some moves.
		// It's probably my hash keys.
		// Store states and cost globally, and if we look at same state but higher cost, don't recurse it; it's already a loser.
		cachedState, exists := stateCache[newState.toKey()]
		if exists && cachedState < newState.cost { // FIXME why does <= not work?
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

		// if moving to the back of the room, add some steps
		cost += COSTS[a] * (ROOMDEPTH - s.rooms[a].finishers - 1)
		moves = append(moves, moveWithCost{room: a, corridor: i, cost: cost, who: a})
	}
	return moves
}

// returns list of moves from rooms into the hall
func (s *state) findMoversInRooms() []moveWithCost {
	moves := []moveWithCost{}

	for ir, r := range s.rooms {

		if r.finishers == ROOMDEPTH {
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
			// if moving from the back of the room, add some steps
			depthInRoom := ROOMDEPTH - len(s.rooms[ir].starters) - s.rooms[ir].finishers
			cost += COSTS[a] * depthInRoom
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
		if r.finishers < ROOMDEPTH {
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
	s2.breadcrumbs = make([]moveWithCost, len(s.breadcrumbs))
	copy(s2.breadcrumbs, s.breadcrumbs)
	return &s2
}

func (s *state) toKey() string {
	return fmt.Sprintf("%v/%v", s.corridor, s.rooms)
}

func (m move) cost(a int) int {
	return PATHS[m].distance * m.room * COSTS[a]
}

// if any starters are in final position, change them from starters to finishers
func (s *state) fixStartingInFinishPosition() {

	for roomNum, r := range s.rooms {
		updatedStarters := []int{}
		// if we run into an amphipod that doesn't belong in the room, stop looking
		// this handles a case where a room contains ABA; the rightmost A is a finished, but the leftmost is not
		// because it has to move to let the B out
		stop := false
		for j := len(r.starters) - 1; j >= 0; j-- {
			starter := r.starters[j]
			if starter != roomNum {
				stop = true
			}
			if starter == roomNum && !stop {
				r.finishers++
			} else {
				// this is being built in reverse order and will need to be re-reversed
				updatedStarters = append(updatedStarters, starter)
			}
		}
		// reverse the order
		for i, j := 0, len(updatedStarters)-1; i < j; i, j = i+1, j-1 {
			updatedStarters[i], updatedStarters[j] = updatedStarters[j], updatedStarters[i]
		}
		r.starters = updatedStarters
		s.rooms[roomNum] = r
	}
}

func (m moveWithCost) toString() string {
	return fmt.Sprintf("moved %d between corridor %d and room %d with cost %d", m.who, m.corridor, m.room, m.cost)
}
