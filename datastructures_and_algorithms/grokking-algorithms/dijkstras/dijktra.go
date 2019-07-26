package main

import (
	"fmt"
	"math"
)

type graph map[string]map[string]int

var processed map[string]bool

var costs map[string]int
var parents map[string]string

func fillCosts(g graph) {
	costs = make(map[string]int)
	for k := range g {
		costs[k] = math.MaxInt8
	}
	costs["START"] = 0
}

func shortestPath(g graph) string {
	parents = make(map[string]string)
	processed = make(map[string]bool)
	fillCosts(g)
	next := "START"
	for next != "" {
		neighbours := g[next]
		for k, v := range neighbours {
			newCost := costs[next] + v
			if costs[k] > newCost {
				costs[k] = newCost
				parents[k] = next
			}
		}
		processed[next] = true
		next = lowCostNode()
	}

	return ""
}

func lowCostNode() string {
	lowestCost := int(math.MaxInt8)
	nextNode := ""
	for node, cost := range costs {
		if _, ok := processed[node]; ok != true {
			if cost < lowestCost {
				lowestCost = cost
				nextNode = node
			}
		}
	}
	return nextNode
}

func printShortestPath(g graph) {
	node := "FINISH"
	fmt.Printf("\n%s", node)
	node = parents[node]
	for node != "START" {
		fmt.Printf("<-%s", node)
		node = parents[node]
	}
	fmt.Printf("<-%s\n", node)
}



