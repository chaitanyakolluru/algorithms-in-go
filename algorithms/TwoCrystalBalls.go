package algorithms

import "math"

// Given that you are given two crystal balls and a building with 100 floors
// and that they both'll break at a certain height from the ground, what is the
// most optimum way of figuring out the best height at which you'd want to
// throw the balls.

func TwoCrystalBalls(breaks []bool) int {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	// find the point at which the breaks list has true in its element,
	// stepping by square root of len of breaks
	var i int
	for i = jmpAmount; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	// come back one step
	i -= jmpAmount

	// now go for one more square root of breaks and see if you can find the
	// first one to come as true
	for j := 0; j < jmpAmount && i < len(breaks); {
		if breaks[i] {
			return i
		}
		j++
		i++
	}

	return -1
}
