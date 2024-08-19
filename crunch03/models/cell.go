package models

type Cell struct {
	x, y          int
	Symbol        rune
	NeighborCount int
	IsAlive       bool
	IsFootprint   bool
}
