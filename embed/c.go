package main

type A struct {
	B
}

type B struct {
	C
	Foo int
}

type C struct {
	Bar int
}


func main() {
	a := A{}
	a.Bar = 1
	println("Bar", a.Bar)
}