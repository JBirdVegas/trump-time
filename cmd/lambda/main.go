package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jbirdvegas/tilltrump.com/internal/timex"
)

func main() {
	ctx := context.Background()
	h := handler{}
	lambda.StartHandlerFunc(h.HandleRequest, lambda.WithContext(ctx), lambda.WithUseNumber(true))
}

type handler struct{}

func (h *handler) HandleRequest(ctx context.Context, _ any) (timex.PresTracker, error) {
	return timex.TillNextPresident(), nil
}
