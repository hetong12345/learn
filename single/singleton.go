package single

import (
	"sync"
)

type Singleton struct {
	name string
	age  int
}

var (
	once     sync.Once
	instance Singleton
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = Singleton{
			name: "singleton",
			age:  18,
		}
	})
	return instance
}
