package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type Builder struct {
	baseurl string
	params  []multimap
	headers []multimap
	method  string
	client  *http.Client
	handler ResponseHandler
}

type multimap struct {
	key    string
	values []string
}

type kvpair struct {
	key, value string
}

// Param sets a query parameter on a request. It overwrites the existing values of a key.
func (rb *Builder) Param(key string, values ...string) *Builder {
	rb.params = append(rb.params, multimap{key, values})
	return rb
}

// Header sets a header on a request. It overwrites the existing values of a key.
func (rb *Builder) Header(key string, values ...string) *Builder {
	rb.headers = append(rb.headers, multimap{key, values})
	return rb
}

// Method sets the HTTP method for a request.
func (rb *Builder) Method(method string) *Builder {
	rb.method = method
	return rb
}

// Handle sets the response handler for a Builder.
// To use multiple handlers, use ChainHandlers.
func (rb *Builder) Handle(h ResponseHandler) *Builder {
	rb.handler = h
	return rb
}

// Request builds a new http.Request with its context set.
func (rb *Builder) Request(ctx context.Context) (req *http.Request, err error) {
	u, err := url.Parse(rb.baseurl)
	if err != nil {
		return nil, fmt.Errorf("could not initialize with base URL %q: %w", u, err)
	}
	if len(rb.params) > 0 {
		q := u.Query()
		for _, kv := range rb.params {
			q[kv.key] = kv.values
		}
		u.RawQuery = q.Encode()
	}
	var body io.Reader
	method := http.MethodGet

	if rb.method != "" {
		method = rb.method
	}
	log.WithFields(log.Fields{
		"method": method,
		"body":   body,
		"url":    u.String(),
	}).Info("HTTP Request")

	req, err = http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		log.Error("Failed to create new request")
		return nil, err
	}

	for _, kv := range rb.headers {
		req.Header[http.CanonicalHeaderKey(kv.key)] = kv.values
	}

	return req, nil
}

// ToString sets the Builder to write the response body to the provided string pointer.
func (rb *Builder) ToString(sp *string) *Builder {
	return rb.Handle(ToString(sp))
}

func URL(baseurl string) *Builder {
	var rb Builder
	rb.baseurl = baseurl
	return &rb
}

// Do calls the underlying http.Client and validates and handles any resulting response. The response body is closed after all validators and the handler run.
func (rb *Builder) Do(req *http.Request) (err error) {
	cl := http.DefaultClient
	if rb.client != nil {
		cl = rb.client
	}

	log.Info("Executing request")
	res, err := cl.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	h := consumeBody
	if rb.handler != nil {
		h = rb.handler
	}
	if err = h(res); err != nil {
		log.Error("Failed to handle response")
		return err
	}
	return nil
}

func consumeBody(res *http.Response) (err error) {
	const maxDiscardSize = 640 * 1 << 10
	if _, err = io.CopyN(io.Discard, res.Body, maxDiscardSize); err == io.EOF {
		err = nil
	}
	return err
}

// Fetch builds a request, sends it, and handles the response.
func (rb *Builder) Fetch(ctx context.Context) (err error) {
	req, err := rb.Request(ctx)
	if err != nil {
		return err
	}
	return rb.Do(req)
}

func (rb *Builder) ToJSON(v any) *Builder {
	return rb.Handle(ToJSON(v))
}
