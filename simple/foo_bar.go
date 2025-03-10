package simple

type foo struct {
}

func NewFoo() *foo {
	return &foo{}
}

type bar struct {
}

func NewBar() *bar {
	return &bar{}
}

type FooBar struct {
	foo *foo
	bar *bar
}


