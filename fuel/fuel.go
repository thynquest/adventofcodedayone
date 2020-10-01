package fuel

import "math"

//Required : fuel required for a specified mass of a module
func Required(mass float64) float64 {
	if mass <= 0 {
		return 0
	}
	return math.Floor(mass/3) - 2
}
