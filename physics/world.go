package physics

type World struct {
	Balls   []*Ball
	Gravity Vector2D
	FloorY  float64 // Y-coordinate of the floor
}

// Add a ball to the world
func (w *World) AddBall(ball *Ball) {
	w.Balls = append(w.Balls, ball)
}

// Update all balls in the world
func (w *World) Update(dt float64) {
	for _, ball := range w.Balls {
		ball.ApplyGravity(w.Gravity)
		ball.Update(dt)
		ball.CheckCollision(w.FloorY) // Check for floor collision
	}
}
