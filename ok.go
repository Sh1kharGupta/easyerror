package easyerror

type Ok[T any] struct {
	Value T
}

func (self *Ok[T]) IsOk() bool {
	return true
}

func (self *Ok[T]) IsErr() bool {
	return false
}

func (self *Ok[T]) Expect(msg string) T {
	return self.Value
}

func (self *Ok[T]) Unwrap() T {
	return self.Value
}

func (self *Ok[T]) UnwrapOr(defaultValue T) T {
	return self.Value
}

func (self *Ok[T]) UnwrapOrElse(defaultFunc func() T) T {
	return self.Value
}

func (self *Ok[T]) ExpectErr(msg string) error {
	panic(msg)
}

func (self *Ok[T]) UnwrapErr() error {
	panic("Can't UnwrapErr on Ok!")
}

func (self *Ok[T]) Err() Option[error] {
	return &None[error]{}
}

func (self *Ok[T]) Ok() Option[T] {
	return &Some[T]{self.Value}
}

func (self *Ok[T]) Map(transformFunc func(T) T) Result[T] {
	return &Ok[T]{transformFunc(self.Value)}
}

func (self *Ok[T]) MapErr(transformFunc func(error) error) Result[T] {
	return self
}

func (self *Ok[T]) MapOr(defaultValue T, transformFunc func(T) T) T {
	return transformFunc(self.Value)
}

func (self *Ok[T]) MapOrElse(defaultFunc func() T, transformFunc func(T) T) T {
	return transformFunc(self.Value)
}

func (self *Ok[T]) And(second Result[T]) Result[T] {
	return second
}

func (self *Ok[T]) Or(second Result[T]) Result[T] {
	return self
}

func (self *Ok[T]) AndThen(transformFunc func(T) Result[T]) Result[T] {
	return transformFunc(self.Value)
}

func (self *Ok[T]) OrElse(transformFunc func(error) Result[T]) Result [T] {
	return self
}
