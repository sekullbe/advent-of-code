package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"regexp"
	"strconv"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 1000))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText, 21))
}

func run1(inputText string, winScore int) int {
	players := parseInput(inputText)
	p1 := players[0]
	p2 := players[1]
	d100 := newDie(100)
	rolls := 0
	for p1.score < winScore && p2.score < winScore {
		for _, p := range players {
			roll := d100.rolls(3)
			p.move(roll)
			//fmt.Printf("player %d universes %d, moves to position %d, score now %d\n", i, roll, p.pos(), p.score)
			rolls += 3
			if p.score >= winScore {
				fmt.Println("Winner!")
				break
			}
		}
	}
	winnerScore := 0
	loserScore := 0
	if p1.score > 1000 {
		winnerScore = p1.score
		loserScore = p2.score
	} else {
		winnerScore = p2.score
		loserScore = p1.score
	}
	_ = winnerScore
	return loserScore * rolls
}

func run2(inputText string, winScore int) int {

	// each die roll is 3 copies
	// calculate all possible paths to end game- each one produces 3^r universes
	// there will be at most 18+21= 39 universes; each player always rolls 1
	// and at least 13 universes; each player always rolls 3

	players := parseInput(inputText)
	p1w, p2w := run2DoIt(players[0].boardpos, players[1].boardpos, winScore)
	if p1w > p2w {
		return p1w
	} else {
		return p2w
	}

}

func run2DoIt(p1start, p2start, winScore int) (int, int) {

	// how many universes were created by rolling each result
	rollToUniverse := map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}
	allgames := make(map[gamestate]int)
	wins := [3]int{0, 0, 0}
	playerTurn := 1
	allgames[newGameState(p1start, p2start)] = 1
	for len(allgames) > 0 {
		newStepGames := make(map[gamestate]int)
		for gs, universes := range allgames {
			for _, dieroll := range []int{3, 4, 5, 6, 7, 8, 9} {
				createdUniverses := rollToUniverse[dieroll]
				newUniverses := universes * createdUniverses
				// apply the roll to the gs, which generates a new gs
				newgs := gs
				np := move(gs.players[playerTurn].pos, dieroll)
				newgs.players[playerTurn].pos = np
				newgs.players[playerTurn].score += np
				if newgs.win(winScore) {
					wins[playerTurn] += newUniverses
					continue
				}
				newStepGames[newgs] += newUniverses
			}

		}
		playerTurn++
		if playerTurn == 3 {
			playerTurn = 1
		}
		allgames = newStepGames
	}

	fmt.Printf("p1: %d   p2: %d\n", wins[1], wins[2])
	return wins[1], wins[2]
}

func parseInput(input string) []*player {
	// this assumes players are in order and doesn't read the ids
	var players []*player
	re := regexp.MustCompile(`\d+$`)
	for id, line := range parsers.SplitByLines(input) {
		start, _ := strconv.Atoi(re.FindString(line))
		p := newPlayer(id+1, start)
		players = append(players, p)
	}
	return players
}
