/*
	Package nim implements a game state usuable by go-mcts

	From: http://mcts.ai/
	A state of the game Nim. In Nim, players alternately take 1,2 or 3 chips with the winner being the player
	to take the last chip. In Nim any initial state of the form 4n+k for k = 1,2,3 is a win for player 1
	(by choosing k) chips. Any initial state of the form 4n is a win for player 2.
*/
package nim

import (
	"github.com/glemzurg/go-mcts"
	"log"
)

const (
	_MAX_PICKABLE_CHIPS = 3 // The max chips a player can pick on their move.
)

// NimSate is the state of a Nim game.
type NimState struct {
	// Public members are used by the scoring function and running the game.
	JustMovedPlayerId uint64 // Who just moved, if they took the last chip they win.
	ActivePlayerId    uint64 // This is the player who's current turn it is.

	// Private members are just used for running the game.
	playerIds []uint64 // The players in this game, the first player moves first.
	chips     uint64   // How many chips are left?
}

// NewNimState creates a new nim game.
func NewNimState(chips uint64, playerIds []uint64) *NimState {
	return &NimState{
		// No player is the just moved played.
		ActivePlayerId: playerIds[0],                         // The first player starts.
		playerIds:      []uint64{playerIds[0], playerIds[1]}, // Only two players allowed.
		chips:          chips,
	}
}

// Log reports the current game state.
func (g *NimState) Log() {
	log.Printf("CHIPS: %d", g.chips)
}

// LogWinner reports the winner.
func (g *NimState) LogWinner() {
	log.Printf("PLAYER %d WINS!", g.JustMovedPlayerId)
}

// Clone makes a deep copy of the game state.
func (g *NimState) Clone() mcts.GameState {
	// Deep copy the player ids.
	var playerIds []uint64 = make([]uint64, len(g.playerIds))
	copy(playerIds, playerIds)
	// Return the new state.
	return &NimState{
		JustMovedPlayerId: g.JustMovedPlayerId,
		ActivePlayerId:    g.ActivePlayerId,
		playerIds:         playerIds,
		chips:             g.chips,
	}
}

// AvailableMoves returns all the available moves.
func (g *NimState) AvailableMoves() []mcts.Move {
	var maxChipsPickable uint64 = g.chips
	if maxChipsPickable > _MAX_PICKABLE_CHIPS {
		maxChipsPickable = _MAX_PICKABLE_CHIPS
	}
	var moves []mcts.Move
	var pickedChips uint64
	for pickedChips = 1; pickedChips <= maxChipsPickable; pickedChips++ {
		moves = append(moves, newNimMove(g.ActivePlayerId, pickedChips))
	}
	return moves // Will be nil if the game is over (no pickable chips).
}

// MakeMove makes a move in the game state, changing it.
func (g *NimState) MakeMove(move mcts.Move) {
	// Convert the move to a form we can use.
	var nimMove *NimMove = move.(*NimMove)
	g.chips -= nimMove.chips
	// It is now the next player's turn.
	g.JustMovedPlayerId = g.ActivePlayerId
	g.nextPlayerActive()
}

// nextPlayerActive makes the next player the active player in the game state.
func (g *NimState) nextPlayerActive() {
	// There are only two players so whichever player is not active should be the new active player.
	for _, playerId := range g.playerIds {
		if playerId != g.ActivePlayerId {
			g.ActivePlayerId = playerId
		}
	}
}

// RandomizeUnknowns has no effect since Nim has no random hidden information.
func (g *NimState) RandomizeUnknowns() {}
