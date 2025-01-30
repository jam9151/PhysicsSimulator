package physics

type Ball struct {
	Position, Velocity, Acceleration Vector2D
	Radius                           float64
	Mass                             float64
	Restitution                      float64 // Bounciness factor (0 = no bounce, 1 = perfect bounce)
}

// Apply gravity
func (b *Ball) ApplyGravity(gravity Vector2D) {
	b.Acceleration = gravity
}

// Update ball position and velocity
func (b *Ball) Update(dt float64) {
	b.Velocity = b.Velocity.Add(b.Acceleration.Scale(dt))
	b.Position = b.Position.Add(b.Velocity.Scale(dt))
}

// Handle collision with ground (or walls)
func (b *Ball) CheckCollision(floorY float64) {
	if b.Position.Y+b.Radius >= floorY {
		b.Position.Y = floorY - b.Radius // Prevent sinking into the floor
		b.Velocity.Y *= -b.Restitution   // Invert velocity and apply restitution
	}
}
