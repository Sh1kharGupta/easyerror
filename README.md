# Easier error handling for Golang

Golang code can end up looking like the following.

```go
func myFunction() (string, error) {
    data1, err := doFirstThing()
    if err != nil {
        return "", err
    }
    data2, err := doSecondThing(data1)
    if err != nil {
        return "", err
    }
    data3, err := doThirdThing(data2)
    if err != nil {
        return "", err
    }
    return data3, nil
}

func main() {
    res, err := myFunction()
    if err == nil {
        fmt.Println(res)
    }
}
```

There is an easier way with `easyerror`.

```go
func myFunction() Result[string] {
    return doFirstThing().Map(doSecondThing).Map(doThirdThing)
}

func main() {
    res := myFunction()
    if res.IsOk() {
        fmt.Println(res.Unwrap())
    }
}
```

This is functionally equivalent to the first snippet and much shorter. The following sections explain how this works.

## The `Result` interface

In the above example, `myFunction()` was returning a `Result[string]`. `Result` is an interface (inspired by Rust https://doc.rust-lang.org/std/result/) implemented by two structs: `Ok` and `Err` - one stores a value, the other stores an error.

```go
type Ok[T any] struct {
    Value T
}
type Err[T any] struct {
    Error error
}
```

**Hence `Result` represents either a value or an error.**

Calling `Map(func(T) T)` on `Ok[T]` applies the given function to `Value T` while calling it on `Err[T]` does nothing.

```go
func (self *Ok[T]) Map(transformFunc func(T) T) Result[T] {
    self.Value = transformFunc(self.Value)
    return self
}
func (self *Err[T]) Map(transformFunc func(T) T) Result[T] {
    return self
}
```

Calling `Unwrap()` on `Ok[T]` returns `Value T` while calling it on `Err[T]` causes a panic.

```go
func (self *Ok[T]) Unwrap() T {
    return self.Value
}
func (self *Err[T]) Unwrap() T {
    panic("Can't Unwrap an error!")
}
```

Now if all `doSomeThing()` functions return a `Result`, then chaining the functions is possible!

```go
func doFirstThing() Result[string] {
    // code
}
func doSecondThing(data string) Result[string] {
    // code
}
func doThirdThing(data string) Result[string] {
    // code
}
func myFunction() Result[string] {
    return doFirstThing().Map(doSecondThing).Map(doThirdThing)
}
func main() {
    res := myFunction()
    if res.IsOk() { // Returns true if res is Ok[T].
        fmt.Println(res.Unwrap())
    }
}
```

Here all functions are returning `Result[string]`. What if the types were different?

```go
func doFirstThing() Result[int] {
    // code
}
func doSecondThing(data int) Result[string] {
    // code
}
func doThirdThing(data string) Result[myCustomType] {
    // code
}
```

In that case, the unbound `Map()` function can be used which does the same thing.

```go
func myFunction() Result[myCustomType] {
    return result.Map[string, myCustomType]( // Map string to myCustomType.
        result.Map[int, string]( // Map int to string.
            doFirstThing(),
            doSecondThing,
        ),
        doThirdThing,
    )
}
```

This is slightly longer than the last code segment but still shorter than the very first code segment.

The `Result` interface offers many more methods to ease writing code. Please read the docs (https://doc.rust-lang.org/std/result/) for a detailed view into the same.

## `Unwrap` and `Catch`

Consider the following example function to open a file, read its contents to a string and then return that string.

```go
func myFunction() Result[*string] {
    // openFile() returns Result[FileObj].
    res1 := openFile("myFileName.txt")
    if res1.IsErr() {
        return &Err[*string]{res1.UnwrapErr()} // failed to open file.
    }
    var str string
    // FileObj.readTo() returns Result[int] (num bytes read).
    res2 := res1.Unwrap().readTo(&str)
    if res2.IsErr() {
        return &Err[*string]{res2.UnwrapErr()} // failed to read file.
    }
    return &Ok[*string]{&str} // all good!
}
```

It is possible to make this shorter using `Map()` and unbound `And()` like so.

```go
func myFunction() Result[*string] {
    var str string
    return result.And[int, *string](
        openFile("myFileName.txt").Map(func(f FileObj) Result[int] {
            return f.readTo(&str)
        }),
        &Ok[*string]{&str},
    )
}
```

But this is not very readable. What is needed is a way to exit `myFunction()` as soon as any `Result` is `Err`.

```go
// Doesn't work - throws panic!
func myFunction() Result[*string] {
    var str string
    openFile("myFileName.txt").Unwrap().readTo(&str).Unwrap()
    return &Ok[*string]{&str}
}
```

This is much cleaner but doesn't work because `Unwrap()` will `panic` if `Result` is `Err`. But what if this `panic` was caught by `recover` and then returned `Err[*string]{<cause of panic>}`? That is exactly what `Catch()` does.

```go
// Works!
func myFunction() (ret Result[*string]) {
    defer result.Catch[*string](&ret)
    var str string
    openFile("myFileName.txt").Unwrap().readTo(&str).Unwrap()
    return &Ok[*string]{&str}
}
```

This is functionally equivalent to the first snippet. `Catch()` takes a pointer to `ret` - the variable being returned - as an argument. When a panic happens, `Catch()` recovers from the panic and checks whether the panic was caused because of calling `Unwrap()` on `Err`. If it was, then it takes the `error` from `Err` (call it `e`) and sets `*ret = &Err[T]{e}`. Otherwise, it "re-panics".

This pattern removes a lot of boilerplate code and was inspired by Rust's question mark (?) operator: https://doc.rust-lang.org/book/ch09-02-recoverable-errors-with-result.html#a-shortcut-for-propagating-errors-the--operator

## The `Option` interface

`easyerror` also provides an interface `Option` (again inspired by Rust https://doc.rust-lang.org/std/option/). `Option` is implemented by two structs: `Some` and `None` - one stores a value, the other stores nothing.
```go
type Some[T any] struct {
    Value T
}
type None[T any] struct {
}
```
`Option` (`Some`/`None`) is quite similar to `Result` (`Ok`/`Err`) and there is a large overlap in the methods implemented by both. The only difference is that `None` holds no value while `Err` holds an error value. `Option` is useful in cases where the value of the error does not matter or if a function can return some value or no value, e.g., find the starting index of a substring in a string - in this case, `Some[int]` can convey that the substring was found along with its index while `None[int]` can convey that the substring was not found.

The `Option` interface also provides a variety of methods to ease writing code. Please read the docs (https://doc.rust-lang.org/std/option/) for a detailed view into the same.

## Tips
- There may be times when one is using `defer Catch()` but still wants a panic if a certain `Result` is `Err`. In that case, use `Result.Expect("")`. `Expect()` is similar to `Unwrap()` but it also accepts a string as argument and panics with that string. Panics from `Expect()` go right past `Catch()` as it only catches panics from `Unwrap()`.

## UT Coverage

|Package  |Coverage|Remarks|
|-|-|-|
|easyerror|100%|
|easyerror/option|97.7%|Minor conditions in Catch() are left|
|easyerror/result|92.3%|Minor conditions in Catch() are left|
