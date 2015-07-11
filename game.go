package mcts

// Move represents a move in the game.
type Move interface {
	Probability() float64 // If this this move has random components indicate the chance it takes place (0.0-1.0).
}

// GameState is the interface a game supports to satisfy the MCTS.
type GameState interface {
	Clone() GameState       // Clone the game state, a deep copy.
	AvailableMoves() []Move // Return all the viable moves given the current game state. For a finished game, nil.
	MakeMove(move Move)     // Take an action, changing the game state.
	RandomizeUnknowns()     // Any game state that is unknown (like order of cards), randomize.
}
