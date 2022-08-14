package easyerror

// Implements the Option interface. Holds no value.
// See the interface for documentation of methods.
type None[T any] struct {}

func (self *None[T]) IsSome() bool {
	return false;
}

func (self *None[T]) IsNone() bool {
	return true;
}

func (self *None[T]) Expect(msg string) T {
	panic(msg)
}

func (self *None[T]) Unwrap() T {
	panic(self)
}

func (self *None[T]) UnwrapOr(defaultValue T) T {
	return defaultValue
}

func (self *None[T]) UnwrapOrElse(defaultFunc func() T) T {
	return defaultFunc()
}

func (self *None[T]) OkOr(err error) Result[T] {
	return &Err[T]{err}
}

func (self *None[T]) OkOrElse(errorFunc func() error) Result[T] {
	return &Err[T]{errorFunc()}
}

func (self *None[T]) Filter(filterFunc func(T) bool) Option[T] {
	return self
}

func (self *None[T]) Map(transformFunc func(T) T) Option[T] {
	return self
}

func (self *None[T]) MapOr(defaultValue T, transformFunc func(T) T) T {
	return defaultValue
}

func (self *None[T]) MapOrElse(defaultFunc func() T, transformFunc func(T) T) T {
	return defaultFunc()
}

func (self *None[T]) ZipWith(second Option[T], transformFunc func(T, T) T) Option[T] {
	return self
}

func (self *None[T]) And(second Option[T]) Option[T] {
	return self
}

func (self *None[T]) Or(second Option[T]) Option[T] {
	return second
}

func (self *None[T]) Xor(second Option[T]) Option[T] {
	return second
}

func (self *None[T]) AndThen(transformFunc func(T) Option[T]) Option[T] {
	return self
}

func (self *None[T]) OrElse(defaultFunc func() Option[T]) Option[T] {
	return defaultFunc()
}
