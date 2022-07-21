package utils

// KeepValueInRangeFor returns a value in range between min and max
func KeepValueInRangeFor(value, min, max int) int {
	v := value
	for v < min {
		v += max
	}
	return v % max
}
