package single

import "fmt"

type Handler interface {
	HandleData(string) (interface{}, error)
}

type DBHandler struct {
	DBname string
}

func (pD *DBHandler) HandleData(data string) (interface{}, error) {
	fmt.Println("Get Db data by query sql :", data)
	return nil, nil
}

type AppHandler struct {
	Handler *DBHandler
}

func (pA *AppHandler) HandleData(data string) (interface{}, error) {
	if true { //Get Data from redis ok
		fmt.Println("Get Data from redis")
	} else {
		pA.Handler.HandleData(data)
		//store ret data data to redis
	}
	return nil, nil
}
