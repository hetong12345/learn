package main

type student struct {
	name  string
	score int
}

type person struct {
	name   string
	age    byte
	isDead bool
}

func main() {
	p := person{name: "zzy", age: 100}
	p.isDead=false
	//isDead(p)
}

func isDead(p interface{}) {
	if p.(person).age < 101 {
		p.(person).isDead = true
	}
}
