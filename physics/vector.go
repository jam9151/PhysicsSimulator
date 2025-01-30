package physics

import "math"

type Vector2D struct {
	X, Y float64
}

// Vector operations
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{v.X + other.X, v.Y + other.Y}
}

func (v Vector2D) Scale(factor float64) Vector2D {
	return Vector2D{v.X * factor, v.Y * factor}
}

func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2D) Normalize() Vector2D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2D{0, 0}
	}
	return v.Scale(1 / mag)
}
