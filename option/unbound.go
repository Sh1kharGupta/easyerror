package option

import (
	"reflect"
	. "github.com/Sh1kharGupta/easyerror"
)

type Pair[T1, T2 any] struct {
	First T1
	Second T2
}

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

func Transpose[T any](input Option[Result[T]]) Result[Option[T]] {
	if input.IsSome() {
		if input.Unwrap().IsOk() {
			return &Ok[Option[T]]{&Some[T]{input.Unwrap().Unwrap()}}
		}
		return &Err[Option[T]]{input.Unwrap().UnwrapErr()}
	}
	return &Ok[Option[T]]{&None[T]{}}
}

func Flatten[T any](input_ Option[Option[T]]) Option[T] {
	// https://github.com/golang/go/issues/53376
	input := input_
	return input.UnwrapOr(&None[T]{})
}

func Map[T1, T2 any](input Option[T1], transformFunc func(T1) T2) Option[T2] {
	if input.IsSome() {
		return &Some[T2]{transformFunc(input.Unwrap())}
	}
	return &None[T2]{}
}

func MapOr[T1, T2 any](input Option[T1], defaultValue T2, transformFunc func(T1) T2) T2 {
	if input.IsSome() {
		return transformFunc(input.Unwrap())
	}
	return defaultValue
}

func MapOrElse[T1, T2 any](input Option[T1], defaultFunc func() T2, transformFunc func(T1) T2) T2 {
	if input.IsSome() {
		return transformFunc(input.Unwrap())
	}
	return defaultFunc()
}

func Zip[T1, T2 any](first Option[T1], second Option[T2]) Option[Pair[T1, T2]] {
	if first.IsSome() && second.IsSome() {
		return &Some[Pair[T1, T2]]{Pair[T1, T2]{first.Unwrap(), second.Unwrap()}}
	}
	return &None[Pair[T1, T2]]{}
}

func ZipWith[T1, T2, T3 any](first Option[T1], second Option[T2], transformFunc func(T1, T2) T3) Option[T3] {
	if first.IsSome() && second.IsSome() {
		return &Some[T3]{transformFunc(first.Unwrap(), second.Unwrap())}
	}
	return &None[T3]{}
}

func And[T1, T2 any](first_ Option[T1], second_ Option[T2]) Option[T2] {
	first := first_
	second := second_
	if first.IsSome() {
		return second
	}
	return &None[T2]{}
}

func AndThen[T1, T2 any](first Option[T1], transformFunc func(T1) Option[T2]) Option[T2] {
	if first.IsNone() {
		return &None[T2]{}
	}
	return transformFunc(first.Unwrap())
}

func Convert[T any](value T, err error) Option[T] {
	if err == nil {
		return &Some[T]{value}
	}
	return &None[T]{}
}
