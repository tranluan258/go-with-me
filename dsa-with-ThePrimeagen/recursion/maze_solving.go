package recursion

type Point struct {
	x int
	y int
}

var direction [][]int = [][]int{
	{0, -1},
	{0, 1},
	{1, 0},
	{-1, 0},
}

func walk(maze [][]string, wall string, curr Point, end Point, path *[]Point, seen [][]bool) bool {
	// Base case
	// 1. off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// 2. wall
	if maze[curr.y][curr.x] == wall {
		return false
	}

	// 3 end game
	if curr.x == end.x && curr.y == end.y {
		// push the end
		*path = append(*path, curr)
		return true
	}

	// if seed = true
	if seen[curr.y][curr.x] {
		return false
	}

	// Push the path
	*path = append(*path, curr)
	seen[curr.y][curr.x] = true

	for _, v := range direction {
		if walk(maze, wall, Point{x: curr.x + v[0], y: curr.y + v[1]}, end, path, seen) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func maze_solve(maze [][]string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze[0])+1)
	for i := range seen {
		seen[i] = make([]bool, len(maze))
	}
	path := make([]Point, 0)

	walk(maze, wall, start, end, &path, seen)
	return path
}
