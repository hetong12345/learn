package data

type Student struct {
	Name  string
	Score int
}
type TwoStudent [2]Student

func (p *TwoStudent) Compare() int {
	return p[0].Score - p[1].Score
}

//func Compare(var1,var2 Comparable) int {
//	return var1.Sub(var1,var2)
//}
//func Sub(var1,var2 Comparable) int {
//	return Sub(var1,var2)
//}

type Integer [2]int

func (i *Integer) Compare() int {
	return i[0] - i[1]
}

func Ints(i, j int) int {
	return i - j
}
