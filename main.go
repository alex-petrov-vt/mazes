package main

import (
	"fmt"
	"time"

	"github.com/alex-petrov-vt/mazes/algorithms"
	"github.com/alex-petrov-vt/mazes/distances"
	"github.com/alex-petrov-vt/mazes/grid"
)

func main() {
	g := grid.NewGrid(6, 6)
	g = algorithms.CreateSidewinderMaze(g, time.Now().UnixNano())
	fmt.Printf("%s", g)

	d := distances.FindDistances(g.Cells[0][0])
	distanceGrid := distances.NewDistanceGrid(6, 6)
	distanceGrid.Grid = g
	distanceGrid.Distances = d
	fmt.Printf("%s", distanceGrid)

	shortestPath := d.FindShortestPath(g.Cells[5][5])
	distanceGrid.Distances = shortestPath
	fmt.Printf("%s", distanceGrid)

	longestPath := d.FindLongestPath()
	distanceGrid.Distances = longestPath
	fmt.Printf("%s", distanceGrid)
}
