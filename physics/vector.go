package physics

import "math"

type Vector2D struct {
	X, Y float64
}

// Subtract two vectors (used for direction)
func (v Vector2D) Subtract(other Vector2D) Vector2D {
	return Vector2D{v.X - other.X, v.Y - other.Y}
}

// Add two vectors
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{v.X + other.X, v.Y + other.Y}
}

// Scale a vector by a factor
func (v Vector2D) Scale(factor float64) Vector2D {
	return Vector2D{v.X * factor, v.Y * factor}
}

// Get vector magnitude (length)
func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize the vector (turn it into a unit vector)
func (v Vector2D) Normalize() Vector2D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2D{0, 0} // Prevent division by zero
	}
	return v.Scale(1 / mag)
}

// Dot product of two vectors
func (v Vector2D) Dot(other Vector2D) float64 {
	return v.X*other.X + v.Y*other.Y
}
