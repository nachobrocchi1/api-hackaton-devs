package handler

import (
	"api-hackaton-devs/entity"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoint.Endpoint) *httptransport.Server {
	server := httptransport.NewServer(
		ep,
		encodeRequest,
		decodeResponse)
	return server
}
func encodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request entity.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	if response.(*entity.Response) == nil {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}

	return nil
}
