// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package panicendpoint

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	customMulti "github.com/uber/zanzibar/examples/example-gateway/endpoints/panic"

	defaultExample "github.com/uber/zanzibar/examples/example-gateway/middlewares/default/default_example"
	defaultExample2 "github.com/uber/zanzibar/examples/example-gateway/middlewares/default/default_example2"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/panic/module"
)

// ServiceCFrontHelloHandler is the handler for "/multi/serviceC_f/hello"
type ServiceCFrontHelloHandler struct {
	Dependencies *module.Dependencies
	endpoint     *zanzibar.RouterEndpoint
}

// NewServiceCFrontHelloHandler creates a handler
func NewServiceCFrontHelloHandler(deps *module.Dependencies) *ServiceCFrontHelloHandler {
	handler := &ServiceCFrontHelloHandler{
		Dependencies: deps,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.ContextExtractor, deps.Default,
		"panic", "panic",
		zanzibar.NewStack([]zanzibar.MiddlewareHandle{
			deps.Middleware.DefaultExample2.NewMiddlewareHandle(
				defaultExample2.Options{},
			),
			deps.Middleware.DefaultExample.NewMiddlewareHandle(
				defaultExample.Options{},
			),
		}, handler.HandleRequest).Handle,
	)

	return handler
}

// Register adds the http handler to the gateway's http router
func (h *ServiceCFrontHelloHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Handle(
		"GET", "/multi/serviceC_f/hello",
		http.HandlerFunc(h.endpoint.HandleRequest),
	)
}

// HandleRequest handles "/multi/serviceC_f/hello".
func (h *ServiceCFrontHelloHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) context.Context {
	defer func() {
		if r := recover(); r != nil {
			stacktrace := string(debug.Stack())
			e := errors.Errorf("enpoint panic: %v, stacktrace: %v", r, stacktrace)
			ctx = h.Dependencies.Default.ContextLogger.ErrorZ(
				ctx,
				"Endpoint failure: endpoint panic",
				zap.Error(e),
				zap.String("stacktrace", stacktrace),
				zap.String("endpoint", h.endpoint.EndpointName))

			h.Dependencies.Default.ContextMetrics.IncCounter(ctx, zanzibar.MetricEndpointPanics, 1)
			res.SendError(502, "Unexpected workflow panic, recovered at endpoint.", nil)
		}
	}()

	// log endpoint request to downstream services
	if ce := h.Dependencies.Default.ContextLogger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		for _, k := range req.Header.Keys() {
			if val, ok := req.Header.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		ctx = h.Dependencies.Default.ContextLogger.DebugZ(ctx, "endpoint request to downstream", zfields...)
	}

	w := customMulti.NewServiceCFrontHelloWorkflow(h.Dependencies)
	if span := req.GetSpan(); span != nil {
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	ctx, response, cliRespHeaders, err := w.Handle(ctx, req.Header)
	// map useful client response headers to server response
	if cliRespHeaders != nil {
		if val, ok := cliRespHeaders.Get(zanzibar.ClientResponseDurationKey); ok {
			if duration, err := time.ParseDuration(val); err == nil {
				res.DownstreamFinishTime = duration
			}
			cliRespHeaders.Unset(zanzibar.ClientResponseDurationKey)
		}
		if val, ok := cliRespHeaders.Get(zanzibar.ClientTypeKey); ok {
			res.ClientType = val
			cliRespHeaders.Unset(zanzibar.ClientTypeKey)
		}
	}

	if err != nil {
		res.SendError(500, "Unexpected server error", err)
		return ctx

	}

	bytes, err := json.Marshal(response)
	if err != nil {
		res.SendError(500, "Unexpected server error", errors.Wrap(err, "Unable to marshal resp json"))
		return ctx
	}
	res.WriteJSONBytes(200, cliRespHeaders, bytes)
	return ctx
}