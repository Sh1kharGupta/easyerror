package easyerror

type Err[T any] struct {
	Error error
}

func (self *Err[T]) IsOk() bool {
	return false
}

func (self *Err[T]) IsErr() bool {
	return true
}

func (self *Err[T]) Expect(msg string) T {
	panic(msg)
}

func (self *Err[T]) Unwrap() T {
	panic(self)
}

func (self *Err[T]) UnwrapOr(defaultValue T) T {
	return defaultValue
}

func (self *Err[T]) UnwrapOrElse(defaultFunc func() T) T {
	return defaultFunc()
}

func (self *Err[T]) ExpectErr(msg string) error {
	return self.Error
}

func (self *Err[T]) UnwrapErr() error {
	return self.Error
}

func (self *Err[T]) Err() Option[error] {
	return &Some[error]{self.Error}
}

func (self *Err[T]) Ok() Option[T] {
	return &None[T]{}
}

func (self *Err[T]) Map(transformFunc func(T) T) Result[T] {
	return self
}

func (self *Err[T]) MapErr(transformFunc func(error) error) Result[T] {
	return &Err[T]{transformFunc(self.Error)}
}

func (self *Err[T]) MapOr(defaultValue T, transformFunc func(T) T) T {
	return defaultValue
}

func (self *Err[T]) MapOrElse(defaultFunc func() T, transformFunc func(T) T) T {
	return defaultFunc()
}

func (self *Err[T]) And(second Result[T]) Result[T] {
	return self
}

func (self *Err[T]) Or(second Result[T]) Result[T] {
	return second
}

func (self *Err[T]) AndThen(transformFunc func(T) Result[T]) Result[T] {
	return self
}

func (self *Err[T]) OrElse(transformFunc func(error) Result[T]) Result[T] {
	return transformFunc(self.Error)
}
