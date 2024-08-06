package arrays

import "math"

func TwoCrystalBall(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jumpAmount
	for ; i < len(breaks); i++ {
		if breaks[i] {
			break
		}
	}

	// jump back
	i = i - jumpAmount

	for j := 0; j <= jumpAmount && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}
