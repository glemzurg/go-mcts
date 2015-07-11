/*
	nim exercies the mcts algorithm for the game of nim.
*/
package main

import (
	"flag"
	"github.com/glemzurg/go-mcts"
	"github.com/glemzurg/go-mcts/examples/games/nim"
	"log"
	"os"
)

func main() {

	// Pass the configuration files as parameters:
	//
	//   bin/nim -ucbc=1.0 -chips=100
	//   bin/nim -h
	//
	var ucbC *float64 = flag.Float64("ucbc", 1.0, "the constant biasing exploitation vs exploration")
	var chips *uint64 = flag.Uint64("chips", 100, "the number of chips in the starting state")
	flag.Parse()

	// Report the files we are using.
	var experimentName string = os.Args[0]
	log.Printf("Experiment: '%s'\n", experimentName)

	// Our two intrepid players.
	var playerA uint64 = 1 // Goes first.
	var playerB uint64 = 2 // Goes second.
	var playerIds []uint64 = []uint64{playerA, playerB}

	// How many iterations do players take when considering moves?
	var iterations uint = 1000

	// How many simulations do players make when valuing the new moves?
	var simulations uint = 100

	// Create the initial game state.
	var state *nim.NimState = nim.NewNimState(*chips, playerIds)

	// Play until the game is over (no more available moves).
	for len(state.AvailableMoves()) > 0 {

		// Log the current game state.
		state.Log()

		// What is the next active player's move?
		var move mcts.Move = mcts.Uct(state, iterations, simulations, *ucbC, state.ActivePlayerId, scoreNim)
		state.MakeMove(move)

		// Report the action taken.
		var nimMove *nim.NimMove = move.(*nim.NimMove)
		nimMove.Log()
	}

	// Report winner.
	state.LogWinner()

	log.Println("Experiment Complete.")
	os.Exit(0)
}

// scoreNim scores the game state from a player's perspective, returning 0.0 (lost), 0.5 (in progress), 1.0 (won)
func scoreNim(playerId uint64, state mcts.GameState) float64 {
	// Is the game over or still in progress?
	var moves []mcts.Move = state.AvailableMoves()
	if len(moves) > 0 {
		// The game is still in progress.
		return 0.5 // Consider it a neutral state (0.0-1.0)
	}

	// The game is over.

	// Convert the state into a form we can use.
	var nimState *nim.NimState = state.(*nim.NimState)
	if playerId == nimState.JustMovedPlayerId {
		// The game is over and we were the last player to move. We win!
		return 1.0
	}
	// We didn't win.
	return 0.0
}
