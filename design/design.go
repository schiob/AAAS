package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Algorithms", func() { // API defines the microservice endpoint and
	Title("Algorithms as a service")                               // other global properties. There should be one
	Description("A simple implementation of different algorithms") // and exactly one API definition appearing in
	Scheme("http")                                                 // the design.
	Host("localhost:8080")

	Origin("*", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})
})

var _ = Resource("dijkstra", func() { // Resources group related API endpoints
	BasePath("/dijkstra")     // together. They map to REST resources for REST
	DefaultMedia(GraphResult) // services.

	Action("compute", func() { // Actions define a single API endpoint together
		Description("Calculate shortest path") // with its path, parameters (both path
		Routing(POST(""))                      // parameters and querystring values) and payload
		Payload(GraphPayload)
		Response(OK)                     // Responses define the shape and status code
		Response(BadRequest, ErrorMedia) // of HTTP responses.
	})

	Origin("*", func() {
		Methods("GET", "DELETE", "OPTIONS")
		Headers("Authorization", "Content-Type", "Content-Disposition")
		Credentials()
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
		Headers("Authorization", "Content-Type")
	})
	Files("/swagger.json", "swagger/swagger.json")
})

var edge = Type("edge", func() {
	Attribute("origin", Integer, "Id of origin node")
	Attribute("destination", Integer, "Id of destination node")
	Attribute("weight", Integer, "Weight of the edge")
	Required("origin", "destination", "weight")
})

var GraphPayload = Type("GraphPayload", func() {
	Attribute("start", Integer, "Id of starting node")
	Attribute("directed", Boolean, "True if directed graph")
	Attribute("node_num", Integer, "number of nodes")
	Attribute("edge_num", Integer, "number of edges")
	Attribute("edges", ArrayOf(edge), "Edges of the graph")
	Required("start", "directed", "node_num", "edge_num", "edges")
})

// BottleMedia defines the media type used to render bottles.
var GraphResult = MediaType("application/vnd.graphresult+json", func() {
	Description("A graph result")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("start", Integer, "ID of start point")
		Attribute("shortest", ArrayOf(Integer), "Resut for every node")
		Attribute("fathers", ArrayOf(Integer), "Father of every node")
		Required("start", "shortest", "fathers")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("start")    // Media types may have multiple views and must
		Attribute("shortest") // have a "default" view.
		Attribute("fathers")
	})
})
