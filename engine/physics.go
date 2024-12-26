package engine

import (
	"bouncing-ball-game/pb"
	"math"

	"google.golang.org/protobuf/proto"
)

type Vector2 struct {
	X, Y float64
}

type Controls struct {
	Speed     float64 // Horizontal movement speed
	JumpForce float64 // Force applied when jumping
}

type InputState struct {
	MoveLeft  bool
	MoveRight bool
	Jump      bool
}

type GameObject struct {
	ID           string    // Unique identifier (e.g., "ball", "player")
	Position     Vector2   // Current position
	Velocity     Vector2   // Current velocity
	Acceleration Vector2   // Current acceleration
	Mass         float64   // Mass of the object
	Elasticity   float64   // Elasticity (bounciness), 1 = perfectly elastic
	Friction     float64   // Coefficient of friction
	Controls     *Controls // Player-specific controls (nil for non-player objects)
	Grounded     bool
}

type PhysicsEngine struct {
	Objects []GameObject
	Gravity Vector2 // Configurable via Lua
	Input   InputState
}

func NewPhysicsEngine(objects []GameObject) PhysicsEngine {
	return PhysicsEngine{
		Objects: objects,
		Gravity: Vector2{X: 0, Y: 0},
		Input:   InputState{},
	}
}

func (engine *PhysicsEngine) Update(dt float64) {
	for i := range engine.Objects {
		obj := &engine.Objects[i]

		// Apply controls if the object has them
		engine.ApplyControls(obj)

		// Apply gravity
		obj.Acceleration.X += engine.Gravity.X
		obj.Acceleration.Y += engine.Gravity.Y

		// Update velocity
		obj.Velocity.X += obj.Acceleration.X * dt
		obj.Velocity.Y += obj.Acceleration.Y * dt

		// Apply friction (simplified for this example)
		obj.Velocity.X *= (1 - obj.Friction*dt)
		obj.Velocity.Y *= (1 - obj.Friction*dt)

		// Update position
		obj.Position.X += obj.Velocity.X * dt
		obj.Position.Y += obj.Velocity.Y * dt

		// Reset acceleration for the next frame
		obj.Acceleration = Vector2{0, 0}

		// Handle collisions with boundaries (canvas edges)
		engine.HandleBoundaryCollision(obj)

		// Check for collisions with the player and ball
		engine.HandleObjectCollisions(obj)
	}
}

// ApplyControls applies player inputs to the object
func (engine *PhysicsEngine) ApplyControls(obj *GameObject) {
	if obj.Controls == nil {
		return
	}

	// Apply horizontal movement
	if engine.Input.MoveLeft {
		obj.Acceleration.X -= obj.Controls.Speed
	}
	if engine.Input.MoveRight {
		obj.Acceleration.X += obj.Controls.Speed
	}

	// Apply jump (only if on the ground, assuming ground is Y=canvasHeight)
	if engine.Input.Jump && obj.Grounded {
		obj.Velocity.Y -= obj.Controls.JumpForce
		obj.Grounded = false // Player is no longer on the ground after jumping
	}
}

func (engine *PhysicsEngine) HandleBoundaryCollision(obj *GameObject) {
	canvasWidth, canvasHeight, objectRadius := 800.0, 600.0, 20.0 // Canvas size and radius

	// Horizontal collision (left and right walls)
	if obj.Position.X-objectRadius < 0 || obj.Position.X+objectRadius > canvasWidth {
		obj.Velocity.X = -obj.Velocity.X * obj.Elasticity
		// Clamp position to ensure object stays within the boundaries
		obj.Position.X = clamp(obj.Position.X, objectRadius, canvasWidth-objectRadius)
	}

	// Vertical collision (ground and ceiling)
	if obj.Position.Y-objectRadius < 0 { // Collision with ceiling
		obj.Velocity.Y = -obj.Velocity.Y * obj.Elasticity
		// Clamp position to ensure object stays within the boundaries
		obj.Position.Y = clamp(obj.Position.Y, objectRadius, canvasHeight-objectRadius)
	} else if obj.Position.Y+objectRadius > canvasHeight { // Collision with ground
		obj.Velocity.Y = -obj.Velocity.Y * obj.Elasticity // Reverse velocity and apply elasticity
		// Clamp position to ensure object stays within the boundaries
		obj.Position.Y = canvasHeight - objectRadius
		obj.Grounded = true // The player is grounded
	} else {
		obj.Grounded = false // The player is in the air, not grounded
	}
}

// HandleObjectCollisions checks for collisions between objects (ball and player in this case)
func (engine *PhysicsEngine) HandleObjectCollisions(obj *GameObject) {
	objectRadius := 20.0

	// Loop through all objects and check for collision with player
	for i := range engine.Objects {
		other := &engine.Objects[i]
		if other.ID != obj.ID { // Don't check collision with itself
			// Calculate the distance between the centers of the player and the ball
			distance := math.Sqrt(math.Pow(obj.Position.X-other.Position.X, 2) + math.Pow(obj.Position.Y-other.Position.Y, 2))

			// Check if the distance is less than the sum of their radii (collision detection)
			if distance < (objectRadius + objectRadius) {
				engine.ResolveCollision(obj, other)
			}
		}
	}
}

// ResolveCollision handles the collision between two objects (player and ball)
func (engine *PhysicsEngine) ResolveCollision(obj1, obj2 *GameObject) {
	objectRadius := 20.0

	// Calculate the vector between the two objects
	dx := obj2.Position.X - obj1.Position.X
	dy := obj2.Position.Y - obj1.Position.Y
	distance := math.Sqrt(dx*dx + dy*dy)

	// Calculate the direction of the collision (normal)
	normX := dx / distance
	normY := dy / distance

	// Calculate the overlap (the amount of penetration)
	overlap := (objectRadius + objectRadius) - distance

	// Move objects apart by half the overlap to resolve collision
	obj1.Position.X -= normX * overlap / 2
	obj1.Position.Y -= normY * overlap / 2
	obj2.Position.X += normX * overlap / 2
	obj2.Position.Y += normY * overlap / 2

	// Calculate relative velocity
	relVelX := obj2.Velocity.X - obj1.Velocity.X
	relVelY := obj2.Velocity.Y - obj1.Velocity.Y

	// Calculate the relative velocity along the normal
	velocityAlongNormal := relVelX*normX + relVelY*normY

	// Only resolve the collision if the objects are moving towards each other
	if velocityAlongNormal < 0 {
		// Calculate the impulse scalar (for collision response)
		e := math.Min(obj1.Elasticity, obj2.Elasticity)
		impulse := (-(1 + e) * velocityAlongNormal) / (1/obj1.Mass + 1/obj2.Mass)

		// Apply the impulse to each object
		obj1.Velocity.X -= impulse * normX / obj1.Mass
		obj1.Velocity.Y -= impulse * normY / obj1.Mass
		obj2.Velocity.X += impulse * normX / obj2.Mass
		obj2.Velocity.Y += impulse * normY / obj2.Mass
	}
}

// Helper to clamp a value between min and max
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// SerializeState serializes the current game state using Protocol Buffers
func (engine *PhysicsEngine) SerializeState() ([]byte, error) {
	state := &pb.PhysicsEngineState{}
	for _, obj := range engine.Objects {
		state.Objects = append(state.Objects, &pb.GameObject{
			Id: obj.ID,
			Position: &pb.Vector2{
				X: obj.Position.X,
				Y: obj.Position.Y,
			},
			Velocity: &pb.Vector2{
				X: obj.Velocity.X,
				Y: obj.Velocity.Y,
			},
			Acceleration: &pb.Vector2{
				X: obj.Acceleration.X,
				Y: obj.Acceleration.Y,
			},
			Mass:       obj.Mass,
			Elasticity: obj.Elasticity,
			Friction:   obj.Friction,
			Grounded:   obj.Grounded, // Add this field if needed
		})
	}
	return proto.Marshal(state)
}
