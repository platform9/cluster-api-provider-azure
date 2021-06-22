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

	"github.com/google/uuid"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type corrIDKey string

// CorrID is a correlation ID that the cluster API provider
// sends with all API requests to Azure. Do not create one
// of these manually. Instead, use the CtxWithCorrelationID function
// to create one of these within a context.Context.
type CorrID string

const corrIDKeyVal corrIDKey = "x-ms-correlation-id"

// StartSpan starts a new span with the given context and tracer, using
// the given spanName for the span's name. This function also ensures
// that a new correlation ID is set on the passed context if none
// existed already. Callers should make sure to do two things with
// the return values of this function. First, only use the new context
// returned from this function. Do not use the one you passed. The
// new one will represent the new span and contain the new correlation
// ID (if one was created and set). Second, make sure to close the
// span when you're done with it. Usually this is done with a defer.
// See below for an example:
//
//	ctx := context.Background()
//	ctx, corrID, newSpan, err := StartSpan(ctx, tele.Tracer(), "myNewSpan")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer newSpan.End()
//	doSomething(ctx)
func StartSpan(
	ctx context.Context,
	tracer oteltrace.Tracer,
	spanName string,
	opts ...oteltrace.SpanOption,
) (context.Context, CorrID, oteltrace.Span, error) {
	ctx, span := tracer.Start(ctx, spanName, opts...)
	ctx, corrID, err := ctxWithCorrID(ctx)
	if err != nil {
		return nil, CorrID(""), nil, err
	}
	return ctx, corrID, span, nil
}

// ctxWithCorrID creates a CorrID and creates a new context.Context
// with the new CorrID in it. It returns the _new_ context and the
// newly creates CorrID. After you call this function, prefer to
// use the newly created context over the old one. Common usage is
// below:
//
// 	ctx := context.Background()
//	ctx, newCorrID := CtxWithCorrID(ctx)
//	fmt.Println("new corr ID: ", newCorrID)
//	doSomething(ctx)
func ctxWithCorrID(ctx context.Context) (context.Context, CorrID, error) {
	currentCorrIDIface := ctx.Value(corrIDKeyVal)
	if currentCorrIDIface != nil {
		currentCorrID, ok := currentCorrIDIface.(CorrID)
		if ok {
			return ctx, currentCorrID, nil
		}
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, CorrID(""), err
	}
	newCorrID := CorrID(uid.String())
	ctx = context.WithValue(ctx, corrIDKeyVal, newCorrID)
	return ctx, newCorrID, nil
}
