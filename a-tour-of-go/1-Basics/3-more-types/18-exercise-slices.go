package main

import "golang.org/x/tour/pic"
	
/* Implement `Pic`. It should return a slice of length `dy`, each element of which
   is a slice of `dx` 8-bit unsigned integers. When you run the program, it will 
   display your picture, interpreting the integers as grayscale (well, bluescale)
   values.
   
   The choice of image is up to you. Interesting functions include `(x+y)/2`,
   `x*y`, and `x^y`.
   
   (You need to use a loop to allocate each `[]uint8` insite the `[][]uint8`.)

   (Use `uint8(intValue)` to convert between types.)   */

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, 0, dy)

	for x := 0; x < dy; x++ {
		row := make([]uint8, 0, dx)
		
		for y := 0; y < dx; y ++ {
			row = append(row, uint8(x*y))
		}

		image = append(image, row)
	}

	return(image)
}

func main() {
	// The image is only displayed on A Tour of Go's website
	pic.Show(Pic)
}
