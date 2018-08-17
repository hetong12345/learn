package data

type Student struct {
	Name  string
	Score int
}

func (p *Student) Compare(c2 Comparable) int {
	return p.Score - c2.(*Student).Score
}

//func Compare(var1,var2 Comparable) int {
//	return var1.Sub(var1,var2)
//}
//func Sub(var1,var2 Comparable) int {
//	return Sub(var1,var2)
//}

type Integer int

func (i Integer) Compare(c2 Comparable) int {
	return int(i) - int(c2.(Integer))
}
