package single

import "fmt"

type Super struct {
	//go语言实现设计模式（三）：模版模式
	GetContent func() string
}

func (d Super) DoOperate() {
	fmt.Println("start exec :", d.GetContent())
}

type LocalDoc struct {
	Super
}

func NewLocalDoc() *LocalDoc {
	c := new(LocalDoc)
	c.Super.GetContent = c.GetContent
	return c
}

func (e *LocalDoc) GetContent() string {
	return "this is a LocalDoc."
}

type NetDoc struct {
	Super
}

func NewNetDoc() *NetDoc {
	c := new(NetDoc)
	c.Super.GetContent = c.GetContent
	return c
}

func (c *NetDoc) GetContent() string {
	return "this is net doc."
}
