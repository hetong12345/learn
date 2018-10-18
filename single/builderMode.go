package single

type Glasses struct {
	Price int64
	From  string
}

type Builder interface {
	BuildPrice(int64)
	BuildForm()
	GetGlasses() Glasses
}

type ShenZhenBuilder struct {
	glasses Glasses
}

type ShanWeiBuilder struct {
	glasses Glasses
}

func (pS *ShenZhenBuilder) BuildPrice(iP int64) {
	pS.glasses.Price = iP * 10
}

func (pS *ShenZhenBuilder) BuildForm() {
	pS.glasses.From = "shenzhen"
}

func (pS *ShenZhenBuilder) GetGlasses() Glasses {
	return pS.glasses
}

func (pS *ShanWeiBuilder) BuildPrice(iP int64) {
	pS.glasses.Price = iP * 2
}

func (pS *ShanWeiBuilder) BuildForm() {
	pS.glasses.From = "shanwei"
}

func (pS *ShanWeiBuilder) GetGlasses() Glasses {
	return pS.glasses
}

type LeshiGlasses struct {
	FirstCost int64
}

func (L *LeshiGlasses) GetGlasses(b Builder) Glasses {
	b.BuildPrice(L.FirstCost)
	b.BuildForm()
	return b.GetGlasses()
}
