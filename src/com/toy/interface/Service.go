package _interface

type Servicer interface {
	Start()
	Log(string)
}

type Register interface {
	Do()
}
