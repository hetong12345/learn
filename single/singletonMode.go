package single

import (
	"sync"
)

type Singleton struct {
	//go语言实现设计模式（一）：单例模式
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
