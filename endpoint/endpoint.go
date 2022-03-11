package endpoint

import (
	"api-hackaton-devs/entity"
	"api-hackaton-devs/service"
	"context"
	"time"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log/level"
)

func MakeServiceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {

		req := in.(entity.Request)

		return svc.GetHackatonWithBestDevs(ctx, &req)
	}
}

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func LoggingEndpointMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			now := time.Now()
			level.Info(logger).Log("Endpoint", "calling endpoint", "started", now)

			defer func(begin time.Time) {
				level.Info(logger).Log("took", time.Since(begin))

				if err != nil {
					level.Info(logger).Log("Result", "NOK")
					level.Error(logger).Log("endpoint_error", err)
				} else {
					level.Info(logger).Log("Result", "OK")
				}
			}(now)
			return next(ctx, request)
		}
	}
}
