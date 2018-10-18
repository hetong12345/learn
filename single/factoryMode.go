package single

type Calc interface {
	//go语言实现设计模式（二）：工厂模式
	SetData(data ...interface{})
	CalcOp() float64
}

type Add struct {
	n1, n2 float64
}

func NewAdd() *Add {
	return new(Add)
}

func (a *Add) SetData(data ...interface{}) {
	if len(data) != 2 {
		panic("too many nums")
	}
	if _, ok := data[0].(float64); !ok {
		panic("error,Need float64 parameters ")
	}
	if _, ok := data[1].(float64); !ok {
		panic("error,Need float64 parameters ")
	}
	a.n1 = data[0].(float64)
	a.n2 = data[1].(float64)
}
func (a *Add) CalcOp() float64 {
	return a.n1 + a.n2
}

type CalcFactory struct {
}

func NewCalcFactory() *CalcFactory {
	instance := new(CalcFactory)
	return instance
}

func (f CalcFactory) CreateOperate(opType string) Calc {
	var op Calc
	switch opType {
	case "+":
		op = NewAdd()
	default:
		panic("error ! dont has this operate")
	}
	return op
}
