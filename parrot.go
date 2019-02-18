package parrot

import (
	"context"
	"log"

	parrotsvc "github.com/bdna/the_parrot_api/gen/parrot"
)

// parrot service example implementation.
// The example methods log the requests and return zero values.
type parrotsrvc struct {
	logger *log.Logger
}

// NewParrot returns the parrot service implementation.
func NewParrot(logger *log.Logger) parrotsvc.Service {
	return &parrotsrvc{logger}
}

// AddParrot implements add-parrot.
func (s *parrotsrvc) AddParrot(ctx context.Context, p *parrotsvc.AddParrotPayload) (res int, err error) {
	s.logger.Print("parrot.add-parrot")
	return
}
