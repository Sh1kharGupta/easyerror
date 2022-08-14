package test_utils

func Assert(expr bool) {
	if !expr {
		panic("Test failed!")
	}
}

func AssertPanic[T1 any, T2 comparable](f func() T1, panicWith T2) {
	defer func(){
		r := recover()
		Assert(r != nil)
		v := r.(T2)
		Assert(v == panicWith)
	}()
	f()
}
