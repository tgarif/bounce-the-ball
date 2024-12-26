declare global {
  interface Window {
    wasmUpdate: (dt: number) => void;
    wasmSerializeState: () => Uint8Array;
    wasmUpdateInput: (input: { MoveLeft: boolean; MoveRight: boolean; Jump: boolean }) => void;
  }
}

export {};
