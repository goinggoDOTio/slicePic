package main

import "golang.org/x/tour/pic"

// Pic function
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for k := range pic[i] {
			pic[i][k] = uint8((k + i) / 2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
