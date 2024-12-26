package engine

import (
	"embed"
	"fmt"
	"strings"

	"github.com/yuin/gopher-lua"
)

//go:embed configs/*.lua
var embeddedLuaFiles embed.FS

// LuaVM represents the Lua virtual machine
type LuaVM struct {
	state *lua.LState
}

// NewLuaVM creates a new Lua virtual machine
func NewLuaVM() *LuaVM {
	return &LuaVM{
		state: lua.NewState(),
	}
}

// LoadGameObjects loads a Lua source file and returns a slice of GameObjects (both player and balls)
func (vm *LuaVM) LoadGameObjects(filename string) ([]GameObject, error) {
	// Read the Lua source code from embedded files
	luaCode, err := embeddedLuaFiles.ReadFile("configs/" + filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read Lua source file %s: %w", filename, err)
	}

	// Execute the Lua source code
	if err := vm.state.DoString(string(luaCode)); err != nil {
		return nil, fmt.Errorf("failed to execute Lua source code: %w", err)
	}

	// Expecting the Lua script to return a table
	luaTable := vm.state.Get(-1)
	if luaTable.Type() != lua.LTTable {
		return nil, fmt.Errorf("expected a Lua table, got %s", luaTable.Type())
	}

	// Initialize a slice to hold the game objects (player + balls)
	var gameObjects []GameObject

	// Iterate through the table to load all objects (player + balls)
	table := luaTable.(*lua.LTable)
	table.ForEach(func(key, value lua.LValue) {
		if tbl, ok := value.(*lua.LTable); ok {
			// Create a new GameObject from the Lua table
			obj := GameObject{
				ID:           tbl.RawGetString("id").String(),
				Mass:         float64(tbl.RawGet(lua.LString("mass")).(lua.LNumber)),
				Elasticity:   float64(tbl.RawGet(lua.LString("elasticity")).(lua.LNumber)),
				Friction:     float64(tbl.RawGet(lua.LString("friction")).(lua.LNumber)),
				Position:     parseVector(tbl.RawGet(lua.LString("position"))),
				Velocity:     parseVector(tbl.RawGet(lua.LString("velocity"))),
				Acceleration: parseVector(tbl.RawGet(lua.LString("acceleration"))),
			}

			// Parse controls if the object ID starts with "player"
			if strings.HasPrefix(obj.ID, "player") {
				if controls := tbl.RawGet(lua.LString("controls")); controls.Type() == lua.LTTable {
					obj.Controls = parseControls(controls.(*lua.LTable))
				}
			}

			// Add the object to the game objects list
			gameObjects = append(gameObjects, obj)
		}
	})

	return gameObjects, nil
}

// LoadGravity loads the gravity configuration from Lua source code
func (vm *LuaVM) LoadGravity(filename string) (Vector2, error) {
	// Read the Lua source code from embedded files
	luaCode, err := embeddedLuaFiles.ReadFile("configs/" + filename)
	if err != nil {
		return Vector2{}, fmt.Errorf("failed to read Lua source file %s: %w", filename, err)
	}

	// Execute the Lua source code
	if err := vm.state.DoString(string(luaCode)); err != nil {
		return Vector2{}, fmt.Errorf("failed to execute Lua source code: %w", err)
	}

	// Expecting the Lua script to return a table
	luaTable := vm.state.Get(-1)
	if luaTable.Type() != lua.LTTable {
		return Vector2{}, fmt.Errorf("expected a Lua table, got %s", luaTable.Type())
	}

	// Parse the gravity vector
	table := luaTable.(*lua.LTable)
	return parseVector(table.RawGet(lua.LString("force"))), nil
}

// parseVector parses a Lua table into a Vector2
func parseVector(value lua.LValue) Vector2 {
	if tbl, ok := value.(*lua.LTable); ok {
		x := tbl.RawGet(lua.LString("x")).(lua.LNumber)
		y := tbl.RawGet(lua.LString("y")).(lua.LNumber)
		return Vector2{X: float64(x), Y: float64(y)}
	}
	return Vector2{}
}

// parseControls parses a Lua table into a Controls struct
func parseControls(tbl *lua.LTable) *Controls {
	return &Controls{
		Speed:     float64(tbl.RawGet(lua.LString("speed")).(lua.LNumber)),
		JumpForce: float64(tbl.RawGet(lua.LString("jumpForce")).(lua.LNumber)),
	}
}
