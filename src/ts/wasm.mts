// @ts-ignore
await import("../wasm_exec.js");

// Inline declaration for Go
declare class Go {
  importObject: WebAssembly.Imports;
  run(instance: WebAssembly.Instance): void;
}

export async function initializeWasm() {
  const go = new Go();
  const wasmModule = await WebAssembly.instantiateStreaming(
    fetch("/assets/main.wasm"),
    go.importObject,
  );
  go.run(wasmModule.instance);

  return {
    wasmUpdate: window.wasmUpdate,
    wasmSerializeState: window.wasmSerializeState,
    wasmUpdateInput: window.wasmUpdateInput,
  };
}
