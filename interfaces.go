package easyerror

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	Expect(string) T
	Unwrap() T
	UnwrapOr(T) T
	UnwrapOrElse(func() T) T
	OkOr(error) Result[T]
	OkOrElse(func() error) Result[T]
	Filter(func(T) bool) Option[T]
	Map(func(T) T) Option[T]
	MapOr(T, func(T) T) T
	MapOrElse(func() T, func(T) T) T
	ZipWith(Option[T], func(T, T) T) Option[T]
	And(Option[T]) Option[T]
	Or(Option[T]) Option[T]
	Xor(Option[T]) Option[T]
	AndThen(func(T) Option[T]) Option[T]
	OrElse(func() Option[T]) Option[T]
}

type Result[T any] interface {
	IsOk() bool
	IsErr() bool
	Expect(string) T
	Unwrap() T
	UnwrapOr(T) T
	UnwrapOrElse(func() T) T
	ExpectErr(string) error
	UnwrapErr() error
	Err() Option[error]
	Ok() Option[T]
	Map(func(T) T) Result[T]
	MapErr(func(error) error) Result[T]
	MapOr(T, func(T) T) T
	MapOrElse(func() T, func(T) T) T
	And(Result[T]) Result[T]
	Or(Result[T]) Result[T]
	AndThen(func(T) Result[T]) Result[T]
	OrElse(func(error) Result[T]) Result[T]
}
