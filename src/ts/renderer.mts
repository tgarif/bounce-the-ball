import { decodeGameState } from "./protobuf.mjs";

export function renderGame(update: (dt: number) => void, serialize: () => Uint8Array) {
  const canvas = document.getElementById("gameCanvas") as HTMLCanvasElement;
  const ctx = canvas.getContext("2d");

  if (!ctx) throw new Error("Canvas not supported by your browser.");

  let lastTimestamp = 0;

  const gameLoop = (timestamp: number) => {
    let dt = (timestamp - lastTimestamp) / 1000; // Convert ms to seconds
    lastTimestamp = timestamp;

    if (dt > 0.05) {
      dt = 0.05; // clamp to 0.05s max to fix the weird movement when switching browser tabs
    }

    // Update physics
    update(dt);

    // Serialize game state and deserialize it for rendering
    const serializedState = serialize();
    const gameState = decodeGameState(serializedState);

    // Clear the canvas
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // Render all objects
    for (const obj of gameState.objects) {
      ctx.beginPath();
      ctx.arc(obj.position.x, obj.position.y, 20, 0, Math.PI * 2); // Render as circles
      ctx.fillStyle = obj.id === "player" ? "blue" : "red";
      ctx.fill();
      ctx.closePath();
    }

    requestAnimationFrame(gameLoop);
  };

  requestAnimationFrame(gameLoop);
}

export function handleInput(
  updateInput: (input: { MoveLeft: boolean; MoveRight: boolean; Jump: boolean }) => void,
) {
  const inputState = { MoveLeft: false, MoveRight: false, Jump: false };

  window.addEventListener("keydown", (e) => {
    if (e.key === "ArrowLeft") inputState.MoveLeft = true;
    if (e.key === "ArrowRight") inputState.MoveRight = true;
    if (e.key === " ") inputState.Jump = true;
    updateInput(inputState);
  });

  window.addEventListener("keyup", (e) => {
    if (e.key === "ArrowLeft") inputState.MoveLeft = false;
    if (e.key === "ArrowRight") inputState.MoveRight = false;
    if (e.key === " ") inputState.Jump = false;
    updateInput(inputState);
  });
}
