package physics

import "math"

type Ball struct {
	Position, Velocity, Acceleration Vector2D
	Radius                           float64
	Mass                             float64
	Restitution                      float64 // Bounciness factor (0 = no bounce, 1 = perfect bounce)
}

// Apply gravity: acceleration affects velocity
func (b *Ball) ApplyGravity(gravity Vector2D) {
	b.Acceleration = b.Acceleration.Add(gravity)
}

// Update velocity and position (Euler integration)
func (b *Ball) Update(dt float64) {
	b.Velocity = b.Velocity.Add(b.Acceleration.Scale(dt))
	b.Position = b.Position.Add(b.Velocity.Scale(dt))
	b.Acceleration = Vector2D{0, 0} // Reset acceleration after update
}

// Handle collision with ground, walls, and ceiling
func (b *Ball) CheckCollision(floorY, screenWidth, screenHeight float64) {
	// Ground collision
	if b.Position.Y+b.Radius >= floorY {
		b.Position.Y = floorY - b.Radius // Prevent sinking
		b.Velocity.Y *= -b.Restitution   // Reverse velocity and apply bounce factor
	}

	// Ceiling collision
	if b.Position.Y-b.Radius <= 0 {
		b.Position.Y = b.Radius // Prevent sticking to the ceiling
		b.Velocity.Y *= -b.Restitution
	}

	// Left wall collision
	if b.Position.X-b.Radius <= 0 {
		b.Position.X = b.Radius // Prevent sticking to the left wall
		b.Velocity.X *= -b.Restitution
	}

	// Right wall collision
	if b.Position.X+b.Radius >= screenWidth {
		b.Position.X = screenWidth - b.Radius // Prevent sticking to the right wall
		b.Velocity.X *= -b.Restitution
	}
}

func (b *Ball) CheckCollisionWithBalls(other *Ball) {
	// Calculate distance between the two balls
	diff := other.Position.Subtract(b.Position)
	distance := diff.Magnitude()

	// Check if collision occurs
	minDistance := b.Radius + other.Radius
	if distance < minDistance && distance > 0 { // Ensure they are colliding
		// Push balls apart (separation correction)
		overlap := minDistance - distance
		correction := diff.Normalize().Scale(overlap / 2)
		b.Position = b.Position.Subtract(correction)
		other.Position = other.Position.Add(correction)

		// Elastic Collision Response (Momentum Conservation)
		normal := diff.Normalize()
		relativeVelocity := other.Velocity.Subtract(b.Velocity)
		speed := relativeVelocity.Dot(normal)

		if speed > 0 { // Ensure they are moving toward each other
			restitution := math.Min(b.Restitution, other.Restitution) // Use the lower restitution
			impulseMagnitude := (2 * speed) / (b.Mass + other.Mass)
			impulse := normal.Scale(impulseMagnitude * restitution)

			// Apply impulse to both balls
			b.Velocity = b.Velocity.Add(impulse.Scale(other.Mass))
			other.Velocity = other.Velocity.Subtract(impulse.Scale(b.Mass))
		}
	}
}
