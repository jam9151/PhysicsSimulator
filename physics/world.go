package physics

type World struct {
	Balls   []*Ball
	Gravity Vector2D
	FloorY  float64 // Y-coordinate of the floor
	Width   float64 // Screen width
	Height  float64 // Screen height
}

// Add a ball to the world
func (w *World) AddBall(ball *Ball) {
	w.Balls = append(w.Balls, ball)
}

// Update all balls in the world
func (w *World) Update(dt float64) {
	for i, ball := range w.Balls {
		ball.ApplyGravity(w.Gravity)
		ball.Update(dt)
		ball.CheckCollision(w.FloorY, w.Width, w.Height) // Wall/Floor collisions

		// Check for ball-to-ball collisions
		for j := i + 1; j < len(w.Balls); j++ {
			ball.CheckCollisionWithBalls(w.Balls[j])
		}
	}
}
