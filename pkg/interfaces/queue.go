package interfaces

import "context"

type IQueue interface {
	Enqueue(ctx context.Context, taskName string, payload map[string]interface{}, opts ...any) error
}
