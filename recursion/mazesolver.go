package recursion

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Point struct {
	x int
	y int
}

func Walk(maze [][]string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// base cases

	// off the map
	if curr.x < 0 || curr.x >= len(maze[0]) ||
		curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// on a wall
	if maze[curr.y][curr.x] == wall {
		return false
	}

	// its the end
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}

	// we've seen this one already
	if seen[curr.y][curr.x] {
		return false
	}

	// actual recursion
	// 3 recurse steps

	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// recurse
	for i := 0; i < len(dir); i++ {
		xyentries := dir[i]
		if Walk(maze, wall, Point{x: curr.x + xyentries[0], y: curr.y + xyentries[1]}, end, seen, path) {
			return true
		}
	}

	// post
	pp := *path
	*path = pp[:]

	return false
}

func MazeSolver(maze [][]string, wall string, start, end Point) []Point {
	seen := [][]bool{}
	path := &[]Point{}
	for i := 0; i < len(maze); i++ {
		seen = append(seen, []bool{false})
	}

	Walk(maze, wall, start, end, seen, path)

	return *path
}
