package main

import (
	"github.com/goadesign/goa"
	"github.com/schiob/AAAS/app"
)

// DijkstraController implements the dijkstra resource.
type DijkstraController struct {
	*goa.Controller
}

// NewDijkstraController creates a dijkstra controller.
func NewDijkstraController(service *goa.Service) *DijkstraController {
	return &DijkstraController{Controller: service.NewController("DijkstraController")}
}

// Compute runs the compute action.
func (c *DijkstraController) Compute(ctx *app.ComputeDijkstraContext) error {
	// DijkstraController_Compute: start_implement
	payload := ctx.Payload
	adjacent := map[int]map[int]int{}
	start := payload.Start
	M := payload.EdgeNum
	N := payload.NodeNum

	for _, edge := range payload.Edges {
		x := edge.Origin
		y := edge.Destination
		s := edge.Weight
		mm, ok := adjacent[x]
		if !ok {
			mm = make(map[int]int)
			adjacent[x] = mm
		}
		mm[y] = s
	}

	if !payload.Directed {
		for _, edge := range payload.Edges {
			y := edge.Origin
			x := edge.Destination
			s := edge.Weight
			mm, ok := adjacent[x]
			if !ok {
				mm = make(map[int]int)
				adjacent[x] = mm
			}
			mm[y] = s
		}
	}

	dist, father := dijkstra(start, N, M, adjacent)
	res := &app.Graphresult{Start: start,
		Shortest: dist,
		Fathers:  father,
	}
	// DijkstraController_Compute: end_implement
	return ctx.OK(res)
}
