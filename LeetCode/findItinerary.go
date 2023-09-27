package main

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}
	
	for key := range graph {
		sort.Sort(sort.Reverse(sort.StringSlice(graph[key])))
	}
	
	var itinerary []string
	
	var dfs func(airport string)
	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			next := graph[airport][len(graph[airport])-1]
			graph[airport] = graph[airport][:len(graph[airport])-1]
			dfs(next)
		}
		itinerary = append(itinerary, airport)
	}
	
	dfs("JFK")
	
	for i := 0; i < len(itinerary)/2; i++ {
		itinerary[i], itinerary[len(itinerary)-1-i] = itinerary[len(itinerary)-1-i], itinerary[i]
	}
	
	return itinerary
}