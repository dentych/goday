package main

import (
	"sync"

	"github.com/dentych/goday/apprication/api"
	"github.com/dentych/goday/james"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go api.StartAPI()
	wg.Done()
	go james.StartAPI()
	wg.Done()

	wg.Wait()
}
