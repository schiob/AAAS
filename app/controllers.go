// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Algorithms": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/schiob/dijkstra/design
// --out=$(GOPATH)/src/github.com/schiob/dijkstra
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// DijkstraController is the controller interface for the Dijkstra actions.
type DijkstraController interface {
	goa.Muxer
	Compute(*ComputeDijkstraContext) error
}

// MountDijkstraController "mounts" a Dijkstra resource controller on the given service.
func MountDijkstraController(service *goa.Service, ctrl DijkstraController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/dijkstra", ctrl.MuxHandler("preflight", handleDijkstraOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewComputeDijkstraContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*GraphPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Compute(rctx)
	}
	h = handleDijkstraOrigin(h)
	service.Mux.Handle("POST", "/dijkstra", ctrl.MuxHandler("compute", h, unmarshalComputeDijkstraPayload))
	service.LogInfo("mount", "ctrl", "Dijkstra", "action", "Compute", "route", "POST /dijkstra")
}

// handleDijkstraOrigin applies the CORS response headers corresponding to the origin.
func handleDijkstraOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Disposition")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalComputeDijkstraPayload unmarshals the request body into the context request data Payload field.
func unmarshalComputeDijkstraPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &graphPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}