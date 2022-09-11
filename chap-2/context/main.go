package main

import (
	"context"
	"fmt"
)

type TraceID string

const ZeroTraceID = ""

type TraceIDKey struct{}

func SetTraceID(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, TraceIDKey{}, tid)
}

func GetTraceID(ctx context.Context) TraceID {
	if v, ok := ctx.Value(TraceIDKey{}).(TraceID); ok {
		return v
	}
	return ZeroTraceID
}

func main() {
	ctx := context.Background()
	fmt.Printf("trace id: %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id: %q\n", GetTraceID(ctx))
}
