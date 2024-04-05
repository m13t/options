package options

import (
	"context"
)

// Apply takes a variadic number of options and applies them
// to the given target. Any options passed that expect a context
// will be provided a new background context.
func Apply[T any, O Option[T]](target *T, opts ...O) error {
	return apply(context.Background(), target, opts...)
}

// ApplyContext works the same as the `Apply` function but allows providing a context
func ApplyContext[T any, O Option[T]](ctx context.Context, target *T, opts ...O) error {
	return apply(ctx, target, opts...)
}

func apply[T any, O Option[T]](ctx context.Context, target *T, opts ...O) error {
	for _, opt := range opts {
		if err := opt.apply(ctx, target); err != nil {
			return err
		}
	}

	return nil
}
