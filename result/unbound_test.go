package result

import (
	"errors"
	"testing"
	. "github.com/TheShikharGupta/easyerror"
	. "github.com/TheShikharGupta/easyerror/test_utils"
)

func TestUnbound(t *testing.T) {
	ok := &Ok[int]{123}
	myError := errors.New("MyError")
	myError2 := errors.New("MyError2")
	ok2 := &Ok[string]{"test"}
	err2 := &Err[string]{myError2}
	err := &Err[int]{myError}

	func1 := func(condition int) (ret Result[int]) {
		defer Catch[int](&ret)
		switch condition {
		case 0:
			err2.Unwrap()
		case 1:
			err2.Expect("expect panic")
		case 2:
			panic("raw panic")
		case 3:
			ok2.Unwrap()
			return ok
		}
		return err
	}
	Assert(func1(0).UnwrapErr() == myError2) // Should be caught.
	AssertPanic[Result[int], string](func() Result[int] {return func1(1)}, "expect panic")
	AssertPanic[Result[int], string](func() Result[int] {return func1(2)}, "raw panic")
	Assert(func1(3).Unwrap() == 123)
	Assert(Transpose[int](&Ok[Option[int]]{&Some[int]{123}}).Unwrap().Unwrap() == 123)
	Assert(Transpose[int](&Ok[Option[int]]{&None[int]{}}).IsNone())
	Assert(Transpose[int](&Err[Option[int]]{myError}).Unwrap().UnwrapErr() == myError)
	func2 := func(int) string {return "test"}
	Assert(Map[int, string](ok, func2).Unwrap() == "test")
	Assert(Map[int, string](err, func2).UnwrapErr() == myError)
	Assert(MapOr[int, string](ok, "default", func2) == "test")
	Assert(MapOr[int, string](err, "default", func2) == "default")
	Assert(MapOrElse[int, string](ok, func() string {return "default"}, func2) == "test")
	Assert(MapOrElse[int, string](err, func() string {return "default"}, func2) == "default")
	Assert(And[int, string](ok, ok2) == ok2)
	Assert(And[int, string](ok, err2) == err2)
	Assert(And[int, string](err, ok2).UnwrapErr() == myError)
	Assert(And[int, string](err, err2).UnwrapErr() == myError)
	Assert(AndThen[int, string](ok, func(int) Result[string] {return ok2}) == ok2)
	Assert(AndThen[int, string](ok, func(int) Result[string] {return err2}) == err2)
	Assert(AndThen[int, string](err, func(int) Result[string] {return ok2}).UnwrapErr() == myError)
	Assert(AndThen[int, string](err, func(int) Result[string] {return err2}).UnwrapErr() == myError)
	func3 := func(condition int) (int, error) {
		switch condition {
		case 0:
			return 123, nil
		default:
			return 0, myError
		}
	}
	Assert(Convert[int](func3(0)).Unwrap() == 123)
	Assert(Convert[int](func3(1)).UnwrapErr() == myError)
}
