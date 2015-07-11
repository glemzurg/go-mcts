/*
	package mcts is an implementation of a Monte Carlo Tree Search.

	More information and example code can be found here: http://mcts.ai/
*/
package mcts

import (
	"math/rand"
	"sort"
)

// Uct is an Upper Confidence Bound Tree search through game stats for an optimal move, given a starting game state.
func Uct(state GameState, iterations uint, simulations uint, ucbC float64, playerId uint64, scorer Scorer) Move {

	// Find the best move given a fixed number of state explorations.
	var root *treeNode = newTreeNode(nil, nil, state, ucbC)
	for i := 0; i < int(iterations); i++ {

		// Start at the top of the tree again.
		var node *treeNode = root

		// Select. Find the node we wish to explore next.
		// While we have complete nodes, dig deeper for a new state to explore.
		for len(node.untriedMoves) == 0 && len(node.children) > 0 {
			// This node has no more moves to try but it does have children.
			// Move the focus to its most promising child.
			node = node.selectChild()
		}

		// Expand.
		// Can we explore more about this particular state? Are there untried moves?
		if len(node.untriedMoves) > 0 {
			node = node.makeRandomUntriedMove() // This creates a new child node with cloned game state.
		}

		// Simulation.
		// From the new child, make many simulated random steps to get a fuzzy idea of how good
		// the move that created the child is.
		var simulatedState GameState = node.state.Clone()
		for j := 0; j < int(simulations); j++ {
			// What moves can further the game state?
			var availableMoves []Move = simulatedState.AvailableMoves()
			// Is the game over?
			if len(availableMoves) == 0 {
				break
			}
			// Pick a random move (could be any player).
			var randomIndex int = rand.Intn(len(availableMoves))
			var move Move = availableMoves[randomIndex]
			simulatedState.MakeMove(move)
		}

		// Backpropagate.
		// Our simulated state may be good or bad in the eyes of our player of interest.
		var outcome float64 = scorer(playerId, simulatedState)
		node.addOutcome(outcome) // Will internally propogate up the tree.
	}

	// The best move to take is going to be the root nodes most visited child.
	sort.Sort(byVisits(root.children))
	return root.children[0].move // Descending by visits.
}
