package test

type Animal interface {
	Eat()
	Run()
}

func Action(a Animal)  {
	a.Run()
	a.Eat()
}
