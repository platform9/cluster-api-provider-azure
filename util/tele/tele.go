/*
Copyright 2020 The Kubernetes Authors.

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

package tele

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type tracer struct {
	trace.Tracer
}

// Start is the trace.Tracer implementation
func (t tracer) Start(
	ctx context.Context,
	op string,
	opts ...trace.SpanOption,
) (context.Context, trace.Span) {
	currentCorrIDIface := ctx.Value(corrIDKeyVal)
	corrID := ""
	if currentCorrIDIface == nil {
		corrID = newCorrID()
	} else {
		cid, ok := currentCorrIDIface.(string)
		if ok {
			corrID = cid
		} else {
			corrID = newCorrID()
		}
	}

	ctx = context.WithValue(ctx, corrIDKeyVal, corrID)
	return t.Tracer.Start(ctx, op, opts...)
}

// newCorrID returns a new correlation ID to be put into a context.Context.
// if there was a problem creating a correlation ID, an empty string will
// be returned
func newCorrID() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return uid.String()
}

// Tracer returns an OpenTelemetry Tracer implementation to be used
// to create spans. Use this implementation instead of the "raw" one that
// you could otherwise get from calling `otel.Tracer("whatever")`.
//
// Example usage:
//
//	ctx, span := tele.Tracer().Start(ctx, "myFunction")
//	defer span.End()
//	// use the span and context here
func Tracer() trace.Tracer {
	return tracer{
		Tracer: otel.Tracer("capz"),
	}
}
