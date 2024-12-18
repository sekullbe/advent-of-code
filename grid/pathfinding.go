package grid

import (
	"github.com/beefsack/go-astar"
	"github.com/sekullbe/advent/geometry"
)

// kind of a gross kludge, but if I change algorithm later I can
// change this to do something like generate a graph from the board
func (b *Board) PrepareForPathfinding() {
	for _, tile := range b.Grid {
		tile.Board = b
	}
	b.pfReady = true
}

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}

	for _, npt := range t.Board.GetSquareNeighbors(t.Point) {
		if !IsWall(t.Board.AtPoint(npt).Contents) {
			neighbors = append(neighbors, t.Board.AtPoint(npt))
		}
	}
	return neighbors
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	//all moves are either blocked or allowed, so movement cost is constant
	// in future, make this based on Tile.Value if exists?
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return float64(ManhattanDistance(t.Point, toT.Point))
}

func (b *Board) FindPath(from, to geometry.Point2) ([]geometry.Point2, int, bool) {
	if !b.pfReady {
		b.PrepareForPathfinding()
	}
	path, distance, found := astar.Path(b.AtPoint(from), b.AtPoint(to))
	pathPts := make([]geometry.Point2, len(path))
	for i, pather := range path {
		pathPts[i] = pather.(*Tile).Point
	}
	return pathPts, int(distance), found

}
