package option

import (
	"reflect"
	. "github.com/Sh1kharGupta/easyerror"
)

// Returned by Zip().
type Pair[T1, T2 any] struct {
	First T1
	Second T2
}

// Can catch panics by Unwrap() on None{}. Please see README for usage.
func Catch[T any](ret *Option[T]) {
	r := recover()
	if r == nil {
		return
	}
	vp := reflect.ValueOf(r)
	if vp.Kind() != reflect.Pointer {
		panic(r)
	}
	v := vp.Elem()
	// TODO: improve this condition to be more stringent.
	if v.Kind() != reflect.Struct || v.NumField() != 0 {
		panic(r)
	}
	*ret = &None[T]{}
}

// Some{Ok{Value}} -> Ok{Some{Value}}
// Some{Err{Error}} -> Err{Error}
// None{} -> Ok{None{}}
func Transpose[T any](input Option[Result[T]]) Result[Option[T]] {
	if input.IsSome() {
		if input.Unwrap().IsOk() {
			return &Ok[Option[T]]{&Some[T]{input.Unwrap().Unwrap()}}
		}
		return &Err[Option[T]]{input.Unwrap().UnwrapErr()}
	}
	return &Ok[Option[T]]{&None[T]{}}
}

// Some{Some{Value}} -> Some{Value}
// None{} -> None{}
func Flatten[T any](input_ Option[Option[T]]) Option[T] {
	// https://github.com/golang/go/issues/53376
	input := input_
	return input.UnwrapOr(&None[T]{})
}

// Similar to the bound Map() method except this one can work with multiple types.
// Please see the `Option` interface for further documentation.
func Map[T1, T2 any](input Option[T1], transformFunc func(T1) T2) Option[T2] {
	if input.IsSome() {
		return &Some[T2]{transformFunc(input.Unwrap())}
	}
	return &None[T2]{}
}

// Similar to the bound MapOr() method except this one can work with multiple types.
// Please see the `Option` interface for further documentation.
func MapOr[T1, T2 any](input Option[T1], defaultValue T2, transformFunc func(T1) T2) T2 {
	if input.IsSome() {
		return transformFunc(input.Unwrap())
	}
	return defaultValue
}

// Similar to the bound MapOrElse() method except this one can work with multiple types.
// Please see the `Option` interface for further documentation.
func MapOrElse[T1, T2 any](input Option[T1], defaultFunc func() T2, transformFunc func(T1) T2) T2 {
	if input.IsSome() {
		return transformFunc(input.Unwrap())
	}
	return defaultFunc()
}

// Some{Value1} + Some{Value2} -> Some{Pair{Value1, Value2}}
// Any other case -> None{}
func Zip[T1, T2 any](first Option[T1], second Option[T2]) Option[Pair[T1, T2]] {
	if first.IsSome() && second.IsSome() {
		return &Some[Pair[T1, T2]]{Pair[T1, T2]{first.Unwrap(), second.Unwrap()}}
	}
	return &None[Pair[T1, T2]]{}
}

// Similar to the bound ZipWith() method except this one can work with multiple types.
// Please see the `Option` interface for further documentation.
func ZipWith[T1, T2, T3 any](first Option[T1], second Option[T2], transformFunc func(T1, T2) T3) Option[T3] {
	if first.IsSome() && second.IsSome() {
		return &Some[T3]{transformFunc(first.Unwrap(), second.Unwrap())}
	}
	return &None[T3]{}
}

// Similar to the bound And() method except this one can work with multiple types.
// Please see the `Option` interface for further documentation.
func And[T1, T2 any](first_ Option[T1], second_ Option[T2]) Option[T2] {
	first := first_
	second := second_
	if first.IsSome() {
		return second
	}
	return &None[T2]{}
}

// Similar to the bound AndThen() method except this one can work with multiple types.
// Please see the `Option` interface for further documentation.
func AndThen[T1, T2 any](first Option[T1], transformFunc func(T1) Option[T2]) Option[T2] {
	if first.IsNone() {
		return &None[T2]{}
	}
	return transformFunc(first.Unwrap())
}

// Convert a function's return (value T, err error) to:-
// Some{value} if err is nil
// None{} if err is not nil
func Convert[T any](value T, err error) Option[T] {
	if err == nil {
		return &Some[T]{value}
	}
	return &None[T]{}
}
