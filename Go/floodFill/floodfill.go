package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	oldColor := image[sr][sc]
	if oldColor != color {
		dfs(image, sr, sc, image[sr][sc], color)
	}
	return image
}

func dfs(image [][]int, r, c, oldColor, color int) {
	if image[r][c] != oldColor {
		return
	}

	image[r][c] = color

	// South
	if r >= 1 {
		dfs(image, r-1, c, oldColor, color)
	}

	// East
	if c >= 1 {
		dfs(image, r, c-1, oldColor, color)
	}

	// North
	if r+1 < len(image) {
		dfs(image, r+1, c, oldColor, color)
	}

	// West
	if c+1 < len(image[r]) {
		dfs(image, r, c+1, oldColor, color)
	}
}
