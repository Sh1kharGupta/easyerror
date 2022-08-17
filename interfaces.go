package easyerror

type Option[T any] interface {
	// Some{Value} -> true
	// None{} -> false
	IsSome() bool

	// Some{Value} -> false
	// None{} -> true
	IsNone() bool

	// Some{Value} -> Value
	// None{} -> panic(self) - can be caught using option.Catch()
	Unwrap() T

	// Some{Value} -> Value
	// None{} -> given arg
	UnwrapOr(T) T

	// Some{Value} -> Value
	// None{} -> return value of given function
	UnwrapOrElse(func() T) T

	// Some{Value} -> Ok{Value}
	// None{} -> Err{given error}
	OkOr(error) Result[T]

	// Some{Value} -> Ok{Value}
	// None{} -> Err{return value of given function}
	OkOrElse(func() error) Result[T]

	// Some{Value} -> Some{Value} if func(Value) is true, else None{}
	// None{} -> None{}
	Filter(func(T) bool) Option[T]

	// Some{Value} -> Some{func(Value)}
	// None{} -> None{}
	Map(func(T) T) Option[T]

	// Some{Value} -> func(Value)
	// None{} -> given arg value
	MapOr(T, func(T) T) T

	// Some{Value} -> second_func(Value)
	// None{} -> return value of first function
	MapOrElse(func() T, func(T) T) T

	// Some{Value1} + Some{Value2} -> Some{func(Value1, Value2)}
	// All other cases -> None{}
	ZipWith(Option[T], func(T, T) T) Option[T]

	// Some{Value} & Any -> Any
	// None{} & Any -> None{}
	And(Option[T]) Option[T]

	// Some{Value} | Any -> Some{Value}
	// None{} | Any -> Any
	Or(Option[T]) Option[T]

	// Some{Value1} ^ Some{Value2} -> None{}
	// Some{Value} ^ None{} -> Some{Value}
	// None{} ^ Any -> Any
	Xor(Option[T]) Option[T]

	// Some{Value} -> func(Value)
	// None{} -> None{}
	AndThen(func(T) Option[T]) Option[T]

	// Some{Value} -> Some{Value}
	// None{} -> return value of given function
	OrElse(func() Option[T]) Option[T]
}

type Result[T any] interface {
	// Ok{Value} -> true
	// Err{Error} -> false
	IsOk() bool

	// Ok{Value} -> false
	// Err{Error} -> true
	IsErr() bool

	// Ok{Value} -> Value
	// Err{Error} -> panic(Err{fmt.Errorf(given string, Error)})
	// > wraps the underlying Error with given string.
	// > can be caught with result.Catch()
	Expect(string) T

	// Ok{Value} -> Value
	// Err{Error} -> panic(self) - can be caught with result.Catch()
	Unwrap() T

	// Ok{Value} -> Value
	// Err{Error} -> given arg value
	UnwrapOr(T) T

	// Ok{Value} -> Value
	// Err{Error} -> return value of given function
	UnwrapOrElse(func() T) T

	// Ok{Value} -> panic(generic string)
	// Err{Error} -> Error
	UnwrapErr() error

	// Ok{Value} -> None{}
	// Err{Error} -> Some{Error}
	Err() Option[error]

	// Ok{Value} -> Some{Value}
	// Err{Error} -> None{}
	Ok() Option[T]

	// Ok{Value} -> Ok{func(Value)}
	// Err{Error} -> Err{Error}
	Map(func(T) T) Result[T]

	// Ok{Value} -> Ok{Value}
	// Err{Error} -> Err{func(Error)}
	MapErr(func(error) error) Result[T]

	// Ok{Value} -> func(Value)
	// Err{Error} -> given arg value
	MapOr(T, func(T) T) T

	// Ok{Value} -> second_func(Value)
	// Err{Error} -> return value of first function
	MapOrElse(func() T, func(T) T) T

	// Ok{Value} & Any -> Any
	// Err{Error} & Any -> Err{Error}
	And(Result[T]) Result[T]

	// Ok{Value} | Any -> Ok{Value}
	// Err{Error} | Any -> Any
	Or(Result[T]) Result[T]

	// Ok{Value} -> func(Value)
	// Err{Error} -> Err{Error}
	AndThen(func(T) Result[T]) Result[T]

	// Ok{Value} -> Ok{Value}
	// Err{Error} -> func(Error)
	OrElse(func(error) Result[T]) Result[T]
}
