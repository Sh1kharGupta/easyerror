package result

import (
	"reflect"
	. "github.com/TheShikharGupta/easyerror"
)

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

func Transpose[T any](input Result[Option[T]]) Option[Result[T]] {
	if input.IsOk() {
		if input.Unwrap().IsSome() {
			return &Some[Result[T]]{&Ok[T]{input.Unwrap().Unwrap()}}
		}
		return &None[Result[T]]{}
	}
	return &Some[Result[T]]{&Err[T]{input.UnwrapErr()}}
}

func Map[T1, T2 any](input Result[T1], transformFunc func(T1) T2) Result[T2] {
	if input.IsOk() {
		return &Ok[T2]{transformFunc(input.Unwrap())}
	}
	return &Err[T2]{input.UnwrapErr()}
}

func MapOr[T1, T2 any](input Result[T1], defaultValue T2, transformFunc func(T1) T2) T2 {
	if input.IsOk() {
		return transformFunc(input.Unwrap())
	}
	return defaultValue
}

func MapOrElse[T1, T2 any](input Result[T1], defaultFunc func() T2, transformFunc func(T1) T2) T2 {
	if input.IsOk() {
		return transformFunc(input.Unwrap())
	}
	return defaultFunc()
}

func And[T1, T2 any](first Result[T1], second Result[T2]) Result[T2] {
	if first.IsOk() {
		return second
	}
	return &Err[T2]{first.UnwrapErr()}
}

func AndThen[T1, T2 any](first Result[T1], transformFunc func(T1) Result[T2]) Result[T2] {
	if first.IsErr() {
		return &Err[T2]{first.UnwrapErr()}
	}
	return transformFunc(first.Unwrap())
}

func Convert[T any](value T, err error) Result[T] {
	if err == nil {
		return &Ok[T]{value}
	}
	return &Err[T]{err}
}
