package main

import (
	"PhysicsSimulator/physics"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const screenWidth = 800
const screenHeight = 600

var world physics.World
var mousePressed bool // Track mouse clicks to prevent spamming

type Game struct{}

// Capture mouse clicks and spawn balls
func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !mousePressed { // Prevent holding from spamming balls
			x, y := ebiten.CursorPosition()
			SpawnBall(float64(x), float64(y))
			mousePressed = true
		}
	} else {
		mousePressed = false
	}

	dt := 1.0 / 60.0 // Fixed timestep
	world.Update(dt)
	return nil
}

// Spawn a ball at (x, y) with a random velocity
func SpawnBall(x, y float64) {
	angle := rand.Float64() * 2 * math.Pi // Random angle in radians
	speed := 200 + rand.Float64()*300     // Random speed between 200-500
	velocity := physics.Vector2D{
		X: math.Cos(angle) * speed,
		Y: math.Sin(angle) * speed,
	}

	world.AddBall(&physics.Ball{
		Position:    physics.Vector2D{X: x, Y: y},
		Velocity:    velocity,
		Radius:      20,
		Mass:        1,
		Restitution: 0.8, // Bounce factor
	})
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, ball := range world.Balls {
		ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Random seed for different results

	world = physics.World{
		Gravity: physics.Vector2D{X: 0, Y: 980}, // Gravity strength
		FloorY:  screenHeight - 50,              // Floor level (lowered slightly)
		Width:   screenWidth,
		Height:  screenHeight,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Click to Spawn a Bouncing Ball")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
