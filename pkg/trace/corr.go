package trace

import (
	"context"

	"github.com/google/uuid"
)

type corrIDKey string

// CorrID is a correlation ID that the cluster API provider
// sends with all API requests to Azure. Do not create one
// of these manually. Instead, use the CtxWithCorrelationID function
// to create one of these within a context.Context.
type CorrID string

const corrIDKeyVal corrIDKey = "x-ms-correlation-id"

// CtxWithCorrID creates a CorrID and creates a new context.Context
// with the new CorrID in it. It returns the _new_ context and the
// newly creates CorrID. After you call this function, prefer to
// use the newly created context over the old one. Common usage is
// below:
//
// 	ctx := context.Background()
//	ctx, newCorrID := CtxWithCorrID(ctx)
//	fmt.Println("new corr ID: ", newCorrID)
//	doSomething(ctx)
func CtxWithCorrID(ctx context.Context) (context.Context, CorrID, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, CorrID(""), err
	}
	newCorrID := CorrID(uid.String())
	ctx = context.WithValue(ctx, corrIDKeyVal, newCorrID)
	return ctx, newCorrID, nil
}
