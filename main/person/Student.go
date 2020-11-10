package person

type Student struct {
	id        string
	name      string
	age       int
	className string
	score     float64
}

func (student Student) eat() {
	println(student.name + " is eating")
}
