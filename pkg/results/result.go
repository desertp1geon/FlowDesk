package results

type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		Value: value,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

func (r Result[T]) IsOk() bool         { return r.Err == nil }
func (r Result[T]) IsErr() bool        { return r.Err != nil }
func (r Result[T]) Unwrap() (T, error) { return r.Value, r.Err }
