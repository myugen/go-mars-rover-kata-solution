package utils

func KeepValueInRangeFor(value, min, max int) int {
	v := value
	for v < min {
		v += max
	}
	return v % max
}
