package utils

// Between
// check if value is in [min, max)
func Between(value, min, max float32) bool {
	return min <= value && value < max
}
