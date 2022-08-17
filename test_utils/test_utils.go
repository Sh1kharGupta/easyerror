package test_utils

func Assert(expr bool) {
	if !expr {
		panic("Test failed!")
	}
}

func Recover[T any](f func()) (ret T) {
	defer func(){
		r := recover()
		Assert(r != nil)
		ret = r.(T)
	}()
	f()
	return ret
}
