package exporter

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	existStatusFailure = "failure"
	exitStatusSuccess  = "success"
)

var (
	requestStartedCounter    prometheus.Counter
	requestCompletedCounter  prometheus.Counter
	resolverStartedCounter   *prometheus.CounterVec
	resolverCompletedCounter *prometheus.CounterVec
	timeToResolveField       *prometheus.HistogramVec
	timeToHandleRequest      *prometheus.HistogramVec
)

type (
	GraphQLTracer struct{}
)

var _ interface {
	graphql.HandlerExtension
	graphql.OperationInterceptor
	graphql.ResponseInterceptor
	graphql.FieldInterceptor
} = GraphQLTracer{}

func GraphQLRegister() {
	GraphQLRegisterOn(prometheus.DefaultRegisterer)
}

func GraphQLRegisterOn(registerer prometheus.Registerer) {
	requestStartedCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "graphql_request_started_total",
			Help: "Total number of requests started on the graphql server.",
		},
	)

	requestCompletedCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "graphql_request_completed_total",
			Help: "Total number of requests completed on the graphql server.",
		},
	)

	resolverStartedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "graphql_resolver_started_total",
			Help: "Total number of resolver started on the graphql server.",
		},
		[]string{"object", "field"},
	)

	resolverCompletedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "graphql_resolver_completed_total",
			Help: "Total number of resolver completed on the graphql server.",
		},
		[]string{"object", "field"},
	)

	timeToResolveField = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "graphql_resolver_duration_ms",
		Help:    "The time taken to resolve a field by graphql server.",
		Buckets: prometheus.ExponentialBuckets(1, 2, 11),
	}, []string{"exitStatus", "object", "field"})

	timeToHandleRequest = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "graphql_request_duration_ms",
		Help:    "The time taken to handle a request by graphql server.",
		Buckets: prometheus.ExponentialBuckets(1, 2, 11),
	}, []string{"exitStatus"})

	registerer.MustRegister(
		requestStartedCounter,
		requestCompletedCounter,
		resolverStartedCounter,
		resolverCompletedCounter,
		timeToResolveField,
		timeToHandleRequest,
	)
}

func GraphQLUnRegister() {
	GraphQLUnRegisterFrom(prometheus.DefaultRegisterer)
}

func GraphQLUnRegisterFrom(registerer prometheus.Registerer) {
	registerer.Unregister(requestStartedCounter)
	registerer.Unregister(requestCompletedCounter)
	registerer.Unregister(resolverStartedCounter)
	registerer.Unregister(resolverCompletedCounter)
	registerer.Unregister(timeToResolveField)
	registerer.Unregister(timeToHandleRequest)
}

func (a GraphQLTracer) ExtensionName() string {
	return "Prometheus"
}

func (a GraphQLTracer) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a GraphQLTracer) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	requestStartedCounter.Inc()
	return next(ctx)
}

func (a GraphQLTracer) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	response := next(ctx)
	errList := graphql.GetErrors(ctx)

	var exitStatus string
	if len(errList) > 0 {
		exitStatus = existStatusFailure
	} else {
		exitStatus = exitStatusSuccess
	}

	oc := graphql.GetOperationContext(ctx)
	observerStart := oc.Stats.OperationStart

	timeToHandleRequest.With(prometheus.Labels{"exitStatus": exitStatus}).
		Observe(float64(time.Since(observerStart).Nanoseconds() / int64(time.Millisecond)))

	requestCompletedCounter.Inc()

	return response
}

func (a GraphQLTracer) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)

	resolverStartedCounter.WithLabelValues(fc.Object, fc.Field.Name).Inc()

	observerStart := time.Now()

	res, err := next(ctx)

	var exitStatus string
	if err != nil {
		exitStatus = existStatusFailure
	} else {
		exitStatus = exitStatusSuccess
	}

	timeToResolveField.WithLabelValues(exitStatus, fc.Object, fc.Field.Name).
		Observe(float64(time.Since(observerStart).Nanoseconds() / int64(time.Millisecond)))

	resolverCompletedCounter.WithLabelValues(fc.Object, fc.Field.Name).Inc()

	return res, err
}
