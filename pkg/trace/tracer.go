/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trace

import (
	"context"

	oteltrace "go.opentelemetry.io/otel/trace"
)

// Tracer is a custom OpenTelemetry tracer implementation. It is
// identical to the underlying OpenTelemetry one, except that it
// adds correlation IDs to the contexts that are passed to it when
// a span is created
type Tracer struct {
	oteltrace.Tracer
}

// StartSpan starts a new span with the given context and tracer, using
// the given spanName for the span's name. This function also ensures
// that a new correlation ID is set on the passed context if none
// existed already. Callers should make sure to do two things with
// the return values of this function.
//
// First, only use the new context returned from this function.
// Do not use the one you passed. The new one will represent the new
// span and contain the new correlation ID (if one was created and set).
//
// Second, make sure to close the span when you're done with it.
// Usually this is done with a defer.
//
// See below for an example:
//
//	ctx := context.Background()
//	ctx, corrID, newSpan, err := myTracer.Start(ctx, "doingAnOp")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer newSpan.End()
//	log.Println("about to do something with correlation ID ", corrID)
//	doSomething(ctx)
func (t *Tracer) Start(
	ctx context.Context,
	opName string,
	opt ...oteltrace.SpanOption,
) (context.Context, CorrID, oteltrace.Span, error) {
	ctx, corrID, err := ctxWithCorrID(ctx)
	if err != nil {
		return nil, CorrID(""), nil, err
	}
	ctx, span := t.Tracer.Start(ctx, opName, opt...)

	return ctx, corrID, span, nil
}
