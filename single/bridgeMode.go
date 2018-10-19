package single

import "fmt"

type Company interface {
	Running()
}

type BigCompany struct {
	Worker
}

func (pB *BigCompany) Running() {
	fmt.Println("员工都是螺丝钉")
	pB.Worker.Leaving()
	fmt.Println("员工跑路后随时可以找人顶替")
}

type SmallCompany struct {
	Worker
}

func (pS *SmallCompany) Running() {
	fmt.Println("随便一个员工都是骨干")
	pS.Worker.Leaving()
	fmt.Println("员工跑路后，公司运转受阻")
}

type Worker interface {
	Leaving()
}

type GoodWorker struct {
}

func (pG *GoodWorker) Leaving() {
	fmt.Println("好员工跑路后")
}

type NormalWorker struct {
}

func (pN *NormalWorker) Leaving() {
	fmt.Println("普通员工跑路后")
}
