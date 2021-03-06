// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// parrot HTTP server
//
// Command:
// $ goa gen github.com/bdna/the_parrot_api/design

package server

import (
	"context"
	"net/http"

	parrotsvc "github.com/bdna/the_parrot_api/gen/parrot"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the parrot service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	AddParrot http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the parrot service endpoints.
func New(
	e *parrotsvc.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"AddParrot", "POST", "/post-parrot/{name}/{breed}/{colour}"},
		},
		AddParrot: NewAddParrotHandler(e.AddParrot, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "parrot" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.AddParrot = m(s.AddParrot)
}

// Mount configures the mux to serve the parrot endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAddParrotHandler(mux, h.AddParrot)
}

// MountAddParrotHandler configures the mux to serve the "parrot" service
// "add-parrot" endpoint.
func MountAddParrotHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/post-parrot/{name}/{breed}/{colour}", f)
}

// NewAddParrotHandler creates a HTTP handler which loads the HTTP request and
// calls the "parrot" service "add-parrot" endpoint.
func NewAddParrotHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAddParrotRequest(mux, dec)
		encodeResponse = EncodeAddParrotResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add-parrot")
		ctx = context.WithValue(ctx, goa.ServiceKey, "parrot")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
