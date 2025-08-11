package moner

type ErrorMonad[IN, OUT any] func(IN) (OUT, error)

func Fmap[A, B, C any](
	fn func(B) C,
	value ErrorMonad[A, B],
) ErrorMonad[A, C] {
	return func(a A) (C, error) {
		v, err := value(a)
		if err != nil {
			var c C
			return c, err
		}

		result := fn(v)
		return result, nil
	}
}

func Bind[A, B, C any](
  f1 ErrorMonad[A, B],
  f2 ErrorMonad[B, C],
) ErrorMonad[A, C] {
	return func(a A) (C, error) {
		value, err := f1(a)
		if err != nil {
			var c C
			return c, err
		}
		return f2(value)
	}
}

func Compose(f ...ErrorMonad[any, any]) ErrorMonad[any, any] {
	b := f[0]
	for _, fn := range f[1:] {
		b = Bind(b, fn)
	}
	return b
}

func WrapFn[I, O any](
	fn func(I) (O, error),
) ErrorMonad[any, any] {
	return func(inp any) (any, error) {
		return fn(inp.(I))
	}
}
