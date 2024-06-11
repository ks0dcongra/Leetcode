package leetcode

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	oldColor := image[sr][sc]
	image[sr][sc] = newColor

	for i := 0; i < 4; i++ {
		if (sr+dir[i][0] >= 0 && sr+dir[i][0] < len(image)) &&
			(sc+dir[i][1] >= 0 && sc+dir[i][1] < len(image[0])) &&
			(image[sr+dir[i][0]][sc+dir[i][1]] == oldColor) {
			floodFill(image, sr+dir[i][0], sc+dir[i][1], newColor)
		}
	}

	return image
}
