package options

import "context"

// Option represents a configuration option to be applied to
// targets of the specified type parameter T.
type Option[T any] interface {
	apply(ctx context.Context, target *T) error
}

// OptionFunc[T] is a function handler that implements the `Option` interface.
type OptionFunc[T any] func(context.Context, *T) error

// apply implements the Option interface.
func (fn OptionFunc[T]) apply(ctx context.Context, target *T) error {
	return fn(ctx, target)
}

// Combine takes any number of Options and returns a single option.
func Combine[T any, O Option[T]](opts ...O) Option[T] {
	return OptionFunc[T](func(ctx context.Context, target *T) error {
		return ApplyContext[T](ctx, target, opts...)
	})
}
