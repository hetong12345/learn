package data

type Student struct {
	Name  string
	Score int
}

func (var1 *Student) CompareTo(var2 Comparable) int {
	return var1.Score - (var2.(*Student).Score)
}
