// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// parrot service
//
// Command:
// $ goa gen github.com/bdna/the_parrot_api/design

package parrotsvc

import (
	"context"
)

// goa is trash
type Service interface {
	// AddParrot implements add-parrot.
	AddParrot(context.Context, *AddParrotPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "parrot"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"add-parrot"}

// AddParrotPayload is the payload type of the parrot service add-parrot method.
type AddParrotPayload struct {
	// Parrot name
	Name string
	// Parrot colour
	Colour string
	// Parrot breed
	Breed string
}
