package option

import (
	"errors"
	"testing"
	. "github.com/Sh1kharGupta/easyerror"
	. "github.com/Sh1kharGupta/easyerror/test_utils"
)

func TestUnbound(t *testing.T) {
	some := &Some[int]{123}
	none := &None[int]{}
	some2 := &Some[string]{"test"}
	none2 := &None[string]{}
	myError := errors.New("MyError")

	func1 := func(condition int) (ret Option[int]) {
		defer Catch[int](&ret)
		switch condition {
		case 0:
			none2.Unwrap()
		case 1:
			none2.Expect("expect panic")
		case 2:
			panic("raw panic")
		case 3:
			some2.Unwrap()
			return some
		}
		return none
	}
	Assert(func1(0) == none) // Should be caught.
	AssertPanic[Option[int], string](func() Option[int] {return func1(1)}, "expect panic")
	AssertPanic[Option[int], string](func() Option[int] {return func1(2)}, "raw panic")
	Assert(func1(3).Unwrap() == 123)
	Assert(Transpose[int](&Some[Result[int]]{&Ok[int]{123}}).Unwrap().Unwrap() == 123)
	Assert(Transpose[int](&Some[Result[int]]{&Err[int]{myError}}).UnwrapErr() == myError)
	Assert(Transpose[int](&None[Result[int]]{}).Unwrap().IsNone())
	Assert(Flatten[int](&Some[Option[int]]{some}).Unwrap() == 123)
	Assert(Flatten[int](&Some[Option[int]]{none}).IsNone())
	Assert(Flatten[int](&None[Option[int]]{}).IsNone())
	func2 := func(int) string {return "test"}
	Assert(Map[int, string](some, func2).Unwrap() == "test")
	Assert(Map[int, string](none, func2).IsNone())
	Assert(MapOr[int, string](some, "default", func2) == "test")
	Assert(MapOr[int, string](none, "default", func2) == "default")
	Assert(MapOrElse[int, string](some, func() string {return "default"}, func2) == "test")
	Assert(MapOrElse[int, string](none, func() string {return "default"}, func2) == "default")
	Assert(Zip[int, string](some, some2).Unwrap() == Pair[int, string]{123, "test"})
	Assert(Zip[int, string](some, none2).IsNone())
	Assert(Zip[int, string](none, some2).IsNone())
	Assert(Zip[int, string](none, none2).IsNone())
	func3 := func(int, string) bool {return true}
	Assert(ZipWith[int, string, bool](some, some2, func3).Unwrap())
	Assert(ZipWith[int, string, bool](some, none2, func3).IsNone())
	Assert(ZipWith[int, string, bool](none, some2, func3).IsNone())
	Assert(ZipWith[int, string, bool](none, none2, func3).IsNone())
	Assert(And[int, string](some, some2) == some2)
	Assert(And[int, string](some, none2) == none2)
	Assert(And[int, string](none, some2).IsNone())
	Assert(And[int, string](none, none2).IsNone())
	Assert(AndThen[int, string](some, func(int) Option[string] {return some2}) == some2)
	Assert(AndThen[int, string](some, func(int) Option[string] {return none2}) == none2)
	Assert(AndThen[int, string](none, func(int) Option[string] {return some2}).IsNone())
	Assert(AndThen[int, string](none, func(int) Option[string] {return none2}).IsNone())
	func4 := func(condition int) (int, error) {
		switch condition {
		case 0:
			return 123, nil
		default:
			return 0, myError
		}
	}
	Assert(Convert[int](func4(0)).Unwrap() == 123)
	Assert(Convert[int](func4(1)).IsNone())
}
