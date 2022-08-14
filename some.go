package easyerror

type Some[T any] struct {
	Value T
}

func (self *Some[T]) IsSome() bool {
	return true;
}

func (self *Some[T]) IsNone() bool {
	return false;
}

func (self *Some[T]) Expect(msg string) T {
	return self.Value
}

func (self *Some[T]) Unwrap() T {
	return self.Value
}

func (self *Some[T]) UnwrapOr(defaultValue T) T {
	return self.Value
}

func (self *Some[T]) UnwrapOrElse(defaultFunc func() T) T {
	return self.Value
}

func (self *Some[T]) OkOr(err error) Result[T] {
	return &Ok[T]{self.Value}
}

func (self *Some[T]) OkOrElse(errorFunc func() error) Result [T] {
	return &Ok[T]{self.Value}
}

func (self *Some[T]) Filter(filterFunc func(T) bool) Option[T] {
	if filterFunc(self.Value) {
		return self
	}
	return &None[T]{}
}

func (self *Some[T]) Map(transformFunc func(T) T) Option[T] {
	return &Some[T]{transformFunc(self.Value)}
}

func (self *Some[T]) MapOr(defaultValue T, transformFunc func(T) T) T {
	return transformFunc(self.Value)
}

func (self *Some[T]) MapOrElse(defaultFunc func() T, transformFunc func(T) T) T {
	return transformFunc(self.Value)
}

func (self *Some[T]) ZipWith(second Option[T], transformFunc func(T, T) T) Option[T] {
	if second.IsSome() {
		return &Some[T]{transformFunc(self.Value, second.Unwrap())}
	}
	return second
}

func (self *Some[T]) And(second Option[T]) Option[T] {
	return second
}

func (self *Some[T]) Or(second Option[T]) Option[T] {
	return self
}

func (self *Some[T]) Xor(second_ Option[T]) Option[T] {
	// https://github.com/golang/go/issues/53376
	second := second_
	if second.IsSome() {
		return &None[T]{}
	}
	return self
}

func (self *Some[T]) AndThen(transformFunc func(T) Option[T]) Option[T] {
	return transformFunc(self.Value)
}

func (self *Some[T]) OrElse(defaultFunc func() Option[T]) Option[T] {
	return self
}
