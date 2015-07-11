package mcts

import (
	"math"
)

// upperConfidenceBound calculates the value of a child node (relative to its parent) for selection. c is a bias parameter, higher favors
// exploration (value children that have not been explored much), lower favors exploitation (value children for the scores
// they've already accumulated).
func upperConfidenceBound(childAggregateOutcome float64, ucbC float64, parentVisits uint64, childVisits uint64) float64 {
	return childAggregateOutcome/float64(childVisits) + ucbC*math.Sqrt(2*math.Log(float64(parentVisits))/float64(childVisits))
}
