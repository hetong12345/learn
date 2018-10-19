package single

import "fmt"

type NewCompany interface {
	Showing()
}

type BaseCompany struct {
}

func (pB *BaseCompany) Showing() {
	fmt.Println("公司有老板，有前台，有人事...")
}

type DevelopingCompany struct {
	NewCompany
}

func (pD *DevelopingCompany) AddWorker() {
	fmt.Println("还有开发、测试、财务人员")
}

func (pD *DevelopingCompany) Showing() {
	fmt.Println("发展中公司中：")
	pD.NewCompany.Showing()
	pD.AddWorker()
}

type NewBigCompany struct {
	NewCompany
}

func (pD *NewBigCompany) AddWorker() {
	fmt.Println("除此之外，个职能人员应有尽有")
}

func (pD *NewBigCompany) Showing() {
	fmt.Println("大型公司中：")
	pD.NewCompany.Showing()
	pD.AddWorker()
}
