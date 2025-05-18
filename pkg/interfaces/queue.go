package interfaces

import "context"

type IQueue interface {
	Enqueue(ctx context.Context, taskName string, payload any, opts ...any) error
}
