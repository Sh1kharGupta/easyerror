package result

import (
	"reflect"
	. "github.com/Sh1kharGupta/easyerror"
)

// Can catch panics by Unwrap() on Err{Error}. Please see README for usage.
func Catch[T any](ret *Result[T]) {
	r := recover()
	if r == nil {
		return
	}
	vp := reflect.ValueOf(r)
	if vp.Kind() != reflect.Pointer {
		panic(r)
	}
	v := vp.Elem()
	// TODO: improve these conditions to be more stringent.
	if v.Kind() != reflect.Struct || v.NumField() != 1 || v.Type().Field(0).Name != "Error" {
		panic(r)
	}
	f := v.Field(0)
	if f.Kind() != reflect.Interface {
		panic(r)
	}
	e, ok := f.Interface().(error)
	if !ok {
		panic(r)
	}
	*ret = &Err[T]{e}
}

// Ok{Some{Value}} -> Some{Ok{Value}}
// Ok{None{}} -> None{}
// Err{Error} -> Some{Err{Error}}
func Transpose[T any](input Result[Option[T]]) Option[Result[T]] {
	if input.IsOk() {
		if input.Unwrap().IsSome() {
			return &Some[Result[T]]{&Ok[T]{input.Unwrap().Unwrap()}}
		}
		return &None[Result[T]]{}
	}
	return &Some[Result[T]]{&Err[T]{input.UnwrapErr()}}
}

// Similar to the bound Map() method except this one can work with multiple types.
// Please see the `Result` interface for further documentation.
func Map[T1, T2 any](input Result[T1], transformFunc func(T1) T2) Result[T2] {
	if input.IsOk() {
		return &Ok[T2]{transformFunc(input.Unwrap())}
	}
	return &Err[T2]{input.UnwrapErr()}
}

// Similar to the bound MapOr() method except this one can work with multiple types.
// Please see the `Result` interface for further documentation.
func MapOr[T1, T2 any](input Result[T1], defaultValue T2, transformFunc func(T1) T2) T2 {
	if input.IsOk() {
		return transformFunc(input.Unwrap())
	}
	return defaultValue
}

// Similar to the bound MapOrElse() method except this one can work with multiple types.
// Please see the `Result` interface for further documentation.
func MapOrElse[T1, T2 any](input Result[T1], defaultFunc func() T2, transformFunc func(T1) T2) T2 {
	if input.IsOk() {
		return transformFunc(input.Unwrap())
	}
	return defaultFunc()
}

// Similar to the bound And() method except this one can work with multiple types.
// Please see the `Result` interface for further documentation.
func And[T1, T2 any](first Result[T1], second Result[T2]) Result[T2] {
	if first.IsOk() {
		return second
	}
	return &Err[T2]{first.UnwrapErr()}
}

// Similar to the bound AndThen() method except this one can work with multiple types.
// Please see the `Result` interface for further documentation.
func AndThen[T1, T2 any](first Result[T1], transformFunc func(T1) Result[T2]) Result[T2] {
	if first.IsErr() {
		return &Err[T2]{first.UnwrapErr()}
	}
	return transformFunc(first.Unwrap())
}

// Convert a function's return (value T, err error) to:-
// Ok{value} if err is nil
// Err{err} if err is not nil
func Convert[T any](value T, err error) Result[T] {
	if err == nil {
		return &Ok[T]{value}
	}
	return &Err[T]{err}
}
