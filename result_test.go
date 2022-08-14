package easyerror

import (
    "errors"
    "testing"
    . "github.com/Sh1kharGupta/easyerror/test_utils"
)

func TestResult(t *testing.T) {
    ok := &Ok[int]{123}
    myError := errors.New("MyError")
    err := &Err[int]{myError}
    Assert(ok.IsOk())
    Assert(!err.IsOk())
    Assert(!ok.IsErr())
    Assert(err.IsErr())
    Assert(ok.Expect("no panic") == 123)
    AssertPanic[int, string](func() int {return err.Expect("panic")}, "panic")
    Assert(ok.Unwrap() == 123)
    AssertPanic[int, *Err[int]](func() int{return err.Unwrap()}, err)
    Assert(ok.UnwrapOr(456) == 123)
    Assert(err.UnwrapOr(456) == 456)
    Assert(ok.UnwrapOrElse(func() int {return 456}) == 123)
    Assert(err.UnwrapOrElse(func() int {return 456}) == 456)
    AssertPanic[error, string](func() error {return ok.ExpectErr("panic")}, "panic")
    Assert(err.ExpectErr("no panic") == myError)
    AssertPanic[error, string](func() error {return ok.UnwrapErr()}, "Can't UnwrapErr on Ok!")
    Assert(err.UnwrapErr() == myError)
    Assert(ok.Err().IsNone())
    Assert(err.Err().Unwrap() == myError)
    Assert(ok.Ok().Unwrap() == 123)
    Assert(err.Ok().IsNone())
    func1 := func(x int) int {return x * 2}
    myError2 := errors.New("MyError2")
    func2 := func(x error) error {return myError2}
    Assert(ok.Map(func1).Unwrap() == 246)
    Assert(err.Map(func1) == err)
    Assert(ok.MapErr(func2) == ok)
    Assert(err.MapErr(func2).UnwrapErr() == myError2)
    Assert(ok.MapOr(456, func1) == 246)
    Assert(err.MapOr(456, func1) == 456)
    Assert(ok.MapOrElse(func() int {return 456}, func1) == 246)
    Assert(err.MapOrElse(func() int {return 456}, func1) == 456)
    ok2 := &Ok[int]{456}
    err2 := &Err[int]{myError2}
    Assert(ok.And(ok2) == ok2)
    Assert(ok.And(err2) == err2)
    Assert(err.And(ok2) == err)
    Assert(err.And(err2) == err)
    Assert(ok.Or(ok2) == ok)
    Assert(ok.Or(err2) == ok)
    Assert(err.Or(ok2) == ok2)
    Assert(err.Or(err2) == err2)
    Assert(ok.AndThen(func(int) Result[int] {return ok2}) == ok2)
    Assert(ok.AndThen(func(int) Result[int] {return err2}) == err2)
    Assert(err.AndThen(func(int) Result[int] {return ok2}) == err)
    Assert(err.AndThen(func(int) Result[int] {return err2}) == err)
    Assert(ok.OrElse(func(error) Result[int] {return ok2}) == ok)
    Assert(ok.OrElse(func(error) Result[int] {return err2}) == ok)
    Assert(err.OrElse(func(error) Result[int] {return ok2}) == ok2)
    Assert(err.OrElse(func(error) Result[int] {return err2}) == err2)
}
