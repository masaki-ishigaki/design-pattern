package singleton

import (
	"sync"
)

type singleton struct{}

var singleInstance *singleton

var lock sync.Mutex

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		singleInstance = &singleton{}
	}

	return singleInstance
}
