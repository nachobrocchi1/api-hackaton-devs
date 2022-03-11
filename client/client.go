package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

type DevsClient interface {
	Call(ctx context.Context, request *Request) (*Response, error)
}

type devsClient struct {
	clientEP endpoint.Endpoint
}

func NewDevsClient(uri string, t time.Duration, log log.Logger) DevsClient {
	ep := makeClientEndpoint(uri, t, log)
	return &devsClient{
		clientEP: ep,
	}
}

func (c *devsClient) Call(ctx context.Context, request *Request) (*Response, error) {
	response, err := c.clientEP(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*Response), nil
}

func makeClientEndpoint(uri string, t time.Duration, log log.Logger) endpoint.Endpoint {
	url, _ := url.Parse(uri)

	opts := []httptransport.ClientOption{
		httptransport.SetClient(&http.Client{Timeout: t}),
	}

	return httptransport.NewClient(
		"GET",
		url,
		encodeRequest,
		decodeResponse,
		opts...,
	).Endpoint()

}

func encodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Host", "aHost")
	var buf bytes.Buffer
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
func decodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	response := new(Response)
	if err := json.Unmarshal(buf.Bytes(), response); err != nil {
		return nil, err
	}

	return response, nil
}
