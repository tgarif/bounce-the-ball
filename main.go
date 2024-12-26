//go:build js && wasm
// +build js,wasm

package main

import (
	"bouncing-ball-game/engine"
	"fmt"
	"log"
	"syscall/js"
)

// Global physics engine
var physicsEngine engine.PhysicsEngine

func wasmUpdate(this js.Value, args []js.Value) interface{} {
	dt := args[0].Float() // Delta time
	physicsEngine.Update(dt)
	return nil
}

func wasmSerializeState(this js.Value, args []js.Value) interface{} {
	data, err := physicsEngine.SerializeState()
	if err != nil {
		log.Fatalf("Failed to serialize game state: %v", err)
	}
	// Convert serialized data to a Uint8Array for JavaScript
	array := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(array, data)
	return array
}

func wasmUpdateInput(_ js.Value, args []js.Value) interface{} {
	// Expecting an object with MoveLeft, MoveRight, and Jump as boolean properties
	input := engine.InputState{
		MoveLeft:  args[0].Get("MoveLeft").Bool(),
		MoveRight: args[0].Get("MoveRight").Bool(),
		Jump:      args[0].Get("Jump").Bool(),
	}
	physicsEngine.Input = input
	return nil
}

func main() {
	// Initialize Lua Virtual Machine
	luaVM := engine.NewLuaVM()

	// Load gravity configuration
	gravity, err := luaVM.LoadGravity("gravity.lua")
	if err != nil {
		log.Fatalf("Failed to load gravity configuration: %v", err)
	}

	// Load lua objects
	objects, err := luaVM.LoadGameObjects("objects.lua")
	if err != nil {
		log.Fatalf("Failed to load objects: %v", err)
	}

	// Initialize physics engine
	physicsEngine = engine.NewPhysicsEngine(objects)
	physicsEngine.Gravity = gravity

	// Expose WebAssembly functions
	js.Global().Set("wasmUpdate", js.FuncOf(wasmUpdate))
	js.Global().Set("wasmSerializeState", js.FuncOf(wasmSerializeState))
	js.Global().Set("wasmUpdateInput", js.FuncOf(wasmUpdateInput))

	fmt.Println("WebAssembly module loaded")
	select {} // Keep the Go program running
}
