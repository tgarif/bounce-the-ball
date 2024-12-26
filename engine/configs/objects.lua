return {
  {
    id = "player",
    mass = 20.0,
    position = { x = 400, y = 550 },
    velocity = { x = 0, y = 0 },
    acceleration = { x = 0, y = 0 },
    friction = 0.2,
    elasticity = 0.2,
    controls = {
      speed = 800, -- Horizontal speed
      jumpForce = 1200, -- Jump force
    },
  },
  {
    id = "ball1",
    mass = 1.0,
    position = { x = 100, y = 350 },
    velocity = { x = 400, y = -1000 },
    acceleration = { x = 0, y = 0 },
    elasticity = 1,
    friction = 0,
  },
  {
    id = "ball2",
    mass = 1.0,
    position = { x = 650, y = 500 },
    velocity = { x = 400, y = -1000 },
    acceleration = { x = 0, y = 0 },
    elasticity = 1,
    friction = 0,
  },
}
