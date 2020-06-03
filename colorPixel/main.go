package main

import (
	"fmt"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	fmt.Println(sr, sc)
	color := image[sr][sc]
	if color == newColor {
		return image
	}
	colorPixel(color, newColor, sr, sc, image)
	return image
}

func colorPixel(withColor, newColor, sr, sc int, image [][]int) {
	image[sr][sc] = newColor
	for j := 0; j <= 1; j++ {
		for i := -1; i <= 1; i += 2 {
			x, y := sr, sc
			if j == 0 {
				x += i
			} else {
				y += i
			}
			if y < 0 || x < 0 || y >= len(image[0]) || x >= len(image) {
				continue
			}
			if image[x][y] == withColor {
				colorPixel(withColor, newColor, x, y, image)
			}
		}
	}
}

func main() {
	a := [][]int{{0, 0, 0}, {0, 1, 1}}
	fmt.Println(floodFill(a, 1, 1, 1))
}
