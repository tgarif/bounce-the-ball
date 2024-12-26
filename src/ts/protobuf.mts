import { game } from "./game_pb.mjs";

export function decodeGameState(data: Uint8Array): game.PhysicsEngineState {
  // Use the Protocol Buffers `decode` method to parse the binary data
  return game.PhysicsEngineState.deserializeBinary(data);
}
