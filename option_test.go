package easyerror

import (
	"errors"
	"testing"
	. "github.com/Sh1kharGupta/easyerror/test_utils"
)

func TestOption(t *testing.T) {
	some := &Some[int]{123}
	none := &None[int]{}
	myError := errors.New("MyError")
	Assert(some.IsSome())
	Assert(!none.IsSome())
	Assert(!some.IsNone())
	Assert(none.IsNone())
	Assert(some.Unwrap() == 123)
	Assert(Recover[*None[int]](func() {none.Unwrap()}) == none)
	Assert(some.UnwrapOr(456) == 123)
	Assert(none.UnwrapOr(456) == 456)
	Assert(some.UnwrapOrElse(func() int {return 456}) == 123)
	Assert(none.UnwrapOrElse(func() int {return 456}) == 456)
	Assert(some.OkOr(myError).Unwrap() == 123)
	Assert(none.OkOr(myError).UnwrapErr() == myError)
	Assert(some.OkOrElse(func() error {return myError}).Unwrap() == 123)
	Assert(none.OkOrElse(func() error {return myError}).UnwrapErr() == myError)
	Assert(some.Filter(func(x int) bool {return true}) == some)
	Assert(some.Filter(func(x int) bool {return false}).IsNone())
	Assert(none.Filter(func(x int) bool {return true}) == none)
	Assert(none.Filter(func(x int) bool {return false}) == none)
	func1 := func(x int) int {return x * 2}
	Assert(some.Map(func1).Unwrap() == 246)
	Assert(none.Map(func1) == none)
	Assert(some.MapOr(456, func1) == 246)
	Assert(none.MapOr(456, func1) == 456)
	Assert(some.MapOrElse(func() int {return 456}, func1) == 246)
	Assert(none.MapOrElse(func() int {return 456}, func1) == 456)
	func2 := func(x, y int) int {return x + y}
	some2 := &Some[int]{456}
	none2 := &None[int]{}
	Assert(some.ZipWith(some2, func2).Unwrap() == 579)
	Assert(some.ZipWith(none2, func2) == none)
	Assert(none.ZipWith(some2, func2) == none)
	Assert(none.ZipWith(none2, func2) == none)
	Assert(some.And(some2) == some2)
	Assert(some.And(none2) == none2)
	Assert(none.And(some2) == none)
	Assert(none.And(none2) == none)
	Assert(some.Or(some2) == some)
	Assert(some.Or(none2) == some)
	Assert(none.Or(some2) == some2)
	Assert(none.Or(none2) == none2)
	Assert(some.Xor(some2).IsNone())
	Assert(some.Xor(none2) == some)
	Assert(none.Xor(some2) == some2)
	Assert(none.Xor(none2) == none2)
	Assert(some.AndThen(func(int) Option[int] {return some2}) == some2)
	Assert(some.AndThen(func(int) Option[int] {return none2}) == none2)
	Assert(none.AndThen(func(int) Option[int] {return some2}) == none)
	Assert(none.AndThen(func(int) Option[int] {return none2}) == none)
	Assert(some.OrElse(func() Option[int] {return some2}) == some)
	Assert(some.OrElse(func() Option[int] {return none2}) == some)
	Assert(none.OrElse(func() Option[int] {return some2}) == some2)
	Assert(none.OrElse(func() Option[int] {return none2}) == none2)
}
