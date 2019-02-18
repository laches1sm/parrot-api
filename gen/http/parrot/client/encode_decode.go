// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// parrot HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/bdna/the_parrot_api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	parrotsvc "github.com/bdna/the_parrot_api/gen/parrot"
	goahttp "goa.design/goa/http"
)

// BuildAddParrotRequest instantiates a HTTP request object with method and
// path set to call the "parrot" service "add-parrot" endpoint
func (c *Client) BuildAddParrotRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		name   string
		breed  string
		colour string
	)
	{
		p, ok := v.(*parrotsvc.AddParrotPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("parrot", "add-parrot", "*parrotsvc.AddParrotPayload", v)
		}
		name = p.Name
		breed = p.Breed
		colour = p.Colour
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddParrotParrotPath(name, breed, colour)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("parrot", "add-parrot", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeAddParrotResponse returns a decoder for responses returned by the
// parrot add-parrot endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeAddParrotResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("parrot", "add-parrot", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("parrot", "add-parrot", resp.StatusCode, string(body))
		}
	}
}