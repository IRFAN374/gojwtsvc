package transport

import (
	"context"
	"fmt"
	"time"

	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"

	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
)

type EndpointsSet struct {
	LoginEndpoint endpoint.Endpoint
}


func InstrumentingEndpoints(endpoints EndpointsSet, tracer opentracinggo.Tracer) EndpointsSet {
	return EndpointsSet{
		LoginEndpoint:    opentracing.TraceServer(tracer, "Login")(endpoints.LoginEndpoint),
	}
}


func LatencyMiddleware(dur metrics.Histogram, methodName string) endpoint.Middleware {

	return func(next endpoint.Endpoint) endpoint.Endpoint {

		dur := dur.With("method", methodName)
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				dur.With("success", fmt.Sprint(err == nil)).Observe(float64(time.Since(begin).Seconds()))

			}(time.Now())
			return next(ctx, request)
		}
	}
}

func RequestFrequencyMiddleware(freq metrics.Gauge, methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		freq := freq.With("method", methodName)

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			freq.Add(1)
			response, err = next(ctx, request)
			freq.Add(-1)
			return response, err
		}

	}
}
