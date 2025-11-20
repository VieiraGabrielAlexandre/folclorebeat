package utils

// Clamp constrains x to the range [min, max].
func Clamp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
