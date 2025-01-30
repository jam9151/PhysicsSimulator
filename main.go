package main

import (
	"PhysicsSimulator/physics"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var world physics.World

type Game struct{}

func (g *Game) Update() error {
	world.Update(1.0 / 60.0) // Simulate at 60 FPS
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, ball := range world.Balls {
		ebitenutil.DrawCircle(screen, ball.Position.X, ball.Position.Y, ball.Radius, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	world = physics.World{
		Gravity: physics.Vector2D{X: 0, Y: 9.8}, // Gravity pulls downward
		FloorY:  550,                            // Floor position
	}

	// Add a bouncing ball
	world.AddBall(&physics.Ball{
		Position:    physics.Vector2D{X: 400, Y: 100},
		Velocity:    physics.Vector2D{X: 0, Y: 0},
		Radius:      20,
		Mass:        1,
		Restitution: 0.8, // 80% velocity retained after bounce
	})

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Bouncing Ball Simulator")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
