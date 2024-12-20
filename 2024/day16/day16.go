package day16

import (
	. "aoc2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Maze struct {
	Map        [][]rune
	Start, End Vec2
}

func (m Maze) At(v Vec2) rune {
	return m.Map[v.Y][v.X]
}

type Edge struct {
	dir   Vec2
	steps int
}

type EdgeList map[Vec2]map[Vec2]Edge

type Graph struct {
	vertices []Vec2
	edges    EdgeList
}

type DistInfo struct {
	distance       int
	previousVertex Vec2
}

type DistList map[Vec2]DistInfo

var buf = bufio.NewReader(os.Stdin)

var (
	NORTH = Vec2{X: 0, Y: -1}
	EAST  = Vec2{X: 1, Y: 0}
	SOUTH = Vec2{X: 0, Y: 1}
	WEST  = Vec2{X: -1, Y: 0}
)

func ReindeerMazeLowestScore(input []string) string {
	maze := getMazeData(input)
	graph := getMazeGraph(maze)
    start := time.Now()
	dist := findDistancesFromVertex(maze.Start, maze.End, graph)
    fmt.Println(time.Since(start))
	return strconv.Itoa(dist)
}

func findDistancesFromVertex(from, to Vec2, graph Graph) int {
	dists := make(DistList)
	for _, v := range graph.vertices {
		dists[v] = DistInfo{math.MaxInt, Vec2{}}
	}
	dists[from] = DistInfo{0, from}

	visited := []Vec2{}
	unvisited := make([]Vec2, len(graph.vertices))
	copy(unvisited, graph.vertices)
	unvisited = sortByDists(unvisited, dists)

	for len(unvisited) != 0 {
		current := unvisited[0]
		for neighbor, edge := range graph.edges[current] {
			if slices.Contains(visited, neighbor) {
				continue
			}

            turnOffset := 0

            previousDir := graph.edges[dists[current].previousVertex][current].dir
            currentDir := graph.edges[current][neighbor].dir
            if !previousDir.Equals(currentDir) {
                turnOffset = 1000
            }

			newDistance := dists[current].distance + edge.steps + turnOffset

			if newDistance < dists[neighbor].distance {
				dists[neighbor] = DistInfo{
					newDistance,
					current,
				}
			}

            if neighbor.Equals(to) {
                return dists[to].distance
            }
		}
        visited = append(visited, current)
		unvisited = sortByDists(unvisited[1:], dists)
        fmt.Println("Visited #:", len(visited), " | Unvisited #:", len(unvisited))
	}

	return dists[to].distance
}

func sortByDists(vertices []Vec2, dists DistList) []Vec2 {
	return MergeSort(vertices, func(a, b Vec2) int {
		return dists[a].distance - dists[b].distance
	})
}

func getMazeGraph(maze Maze) Graph {
	queue := []Vec2{maze.Start}
	vertices := []Vec2{}
	edges := make(EdgeList)

	// For starting direction
    addEdge(maze.Start, maze.Start, Edge{EAST, 0}, edges)

	for len(queue) != 0 {
		currentVertex := queue[0]
		vertices = append(vertices, currentVertex)

		nextQueue := []Vec2{}

		northFound, northVertex, northSteps := searchForVertex(currentVertex, NORTH, 0, maze)
		if northFound {
			if !slices.Contains(vertices, northVertex) {
				addEdge(currentVertex, northVertex, Edge{NORTH, northSteps}, edges)
				addEdge(northVertex, currentVertex, Edge{SOUTH, northSteps}, edges)
				nextQueue = append(nextQueue, northVertex)
			}
		}

		eastFound, eastVertex, eastSteps := searchForVertex(currentVertex, EAST, 0, maze)
		if eastFound {
			if !slices.Contains(vertices, eastVertex) {
				addEdge(currentVertex, eastVertex, Edge{EAST, eastSteps}, edges)
				addEdge(eastVertex, currentVertex, Edge{WEST, eastSteps}, edges)
				nextQueue = append(nextQueue, eastVertex)
			}
		}

		southFound, southVertex, southSteps := searchForVertex(currentVertex, SOUTH, 0, maze)
		if southFound {
			if !slices.Contains(vertices, southVertex) {
				addEdge(currentVertex, southVertex, Edge{SOUTH, southSteps}, edges)
				addEdge(southVertex, currentVertex, Edge{NORTH, southSteps}, edges)
				nextQueue = append(nextQueue, southVertex)
			}
		}

		westFound, westVertex, westSteps := searchForVertex(currentVertex, WEST, 0, maze)
		if westFound {
			if !slices.Contains(vertices, westVertex) {
				addEdge(currentVertex, westVertex, Edge{WEST, westSteps}, edges)
				addEdge(westVertex, currentVertex, Edge{EAST, westSteps}, edges)
				nextQueue = append(nextQueue, westVertex)
			}
		}

		queue = append(queue, nextQueue...)
		// Shift
		queue = queue[1:]
	}

	return Graph{vertices, edges}
}

func searchForVertex(start, dir Vec2, stepsTaken int, maze Maze) (found bool, vertex Vec2, steps int) {
	search := start.Add(dir)
	stepsTaken++
	if maze.At(search) == '#' {
		return false, Vec2{}, 0
	}

	if isVertex(search, maze) {
		return true, search, stepsTaken
	}

	return searchForVertex(search, dir, stepsTaken, maze)
}

func addEdge(v1, v2 Vec2, edge Edge, edges EdgeList) {
	if edges[v1] == nil {
		edges[v1] = make(map[Vec2]Edge)
	}
	edges[v1][v2] = edge
}

func isVertex(loc Vec2, maze Maze) bool {
	if loc.Equals(maze.Start) || loc.Equals(maze.End) {
		return true
	}

	vEdges, hEdges := false, false
	if maze.At(loc.Add(NORTH)) == '.' || maze.At(loc.Add(SOUTH)) == '.' {
		vEdges = true
	}
	if maze.At(loc.Add(EAST)) == '.' || maze.At(loc.Add(WEST)) == '.' {
		hEdges = true
	}
	return vEdges && hEdges
}

func getMazeData(input []string) Maze {
	maze := Maze{
		[][]rune{},
		Vec2{},
		Vec2{},
	}
	for y, row := range input {
		startX := strings.Index(row, "S")
		if startX != -1 {
			maze.Start = Vec2{X: startX, Y: y}
		}
		endX := strings.Index(row, "E")
		if endX != -1 {
			maze.End = Vec2{X: endX, Y: y}
		}
		maze.Map = append(maze.Map, []rune(row))
	}
	return maze
}

func printMazeGraph(maze Maze, graph Graph) {
	mapCopy := make([][]rune, len(maze.Map))
	for i := range maze.Map {
		newRunes := make([]rune, len(maze.Map[i]))
		copy(newRunes, maze.Map[i])
		mapCopy[i] = newRunes
	}

	for _, v := range graph.vertices {
		mapCopy[v.Y][v.X] = '@'
	}

	for v1, vEdges := range graph.edges {
		for v2, edge := range vEdges {
			for e := v1.Add(edge.dir); !e.Equals(v2); e.AddMut(edge.dir) {
				if edge.dir.Equals(NORTH) {
					mapCopy[e.Y][e.X] = '^'
				}
				if edge.dir.Equals(EAST) {
					mapCopy[e.Y][e.X] = '>'
				}
				if edge.dir.Equals(SOUTH) {
					mapCopy[e.Y][e.X] = 'v'
				}
				if edge.dir.Equals(WEST) {
					mapCopy[e.Y][e.X] = '<'
				}
			}
		}
	}

	for _, line := range mapCopy {
		str := string(line)
		str = strings.ReplaceAll(str, ".", " ")
		str = strings.ReplaceAll(str, "#", " ")
		fmt.Println(str)
	}
}
