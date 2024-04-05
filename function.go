package options

import (
	"context"
)

// Function[T] constraint represents the union of function signatures that can
// be used with the `New` function to create a new `Option[T]`
type Function[T any] interface {
	func(*T) | func(*T) error | func(context.Context, *T) | func(context.Context, *T) error
}

// New takes a function that satisfies the `Function[T]` constraints and returns
// an `Option[T]` interface.
func New[T any, F Function[T]](fn F) Option[T] {
	switch fn := any(fn).(type) {
	case func(*T):
		return OptionFunc[T](func(_ context.Context, target *T) error {
			fn(target)
			return nil
		})

	case func(*T) error:
		return OptionFunc[T](func(_ context.Context, target *T) error {
			return fn(target)
		})

	case func(context.Context, *T):
		return OptionFunc[T](func(ctx context.Context, target *T) error {
			fn(ctx, target)
			return nil
		})

	case func(context.Context, *T) error:
		return OptionFunc[T](fn)

	default:
		panic("invalid option function signature")
	}
}
