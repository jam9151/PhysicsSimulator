package physics

// Global physics constants
var (
	Gravity = Vector2D{X: 0, Y: 9.8} // Default gravity
)

// Clamp function to prevent values from exceeding limits
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}
