import { initializeWasm } from "./wasm.mjs";
import { renderGame, handleInput } from "./renderer.mjs";

async function main() {
  try {
    const wasm = await initializeWasm();
    const update = wasm.wasmUpdate;
    const serialize = wasm.wasmSerializeState;
    const updateInput = wasm.wasmUpdateInput;

    // Attach input handlers
    handleInput(updateInput);

    // Start the rendering loop
    renderGame(update, serialize);
  } catch (error) {
    console.error("Failed to initialize the game:", error);
  }
}

main();
