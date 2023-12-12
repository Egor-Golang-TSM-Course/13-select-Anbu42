package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Chef struct {
	Name        string
	CookingTime time.Duration
}

func (chef *Chef) Cook() {
	fmt.Printf("%s: Starts cooking\n", chef.Name)
	time.Sleep(chef.CookingTime)
	fmt.Printf("%s: Finished cooking\n", chef.Name)
}

func main() {
	chefs := []Chef{
		{Name: "Chef 1", CookingTime: time.Second * time.Duration(rand.Intn(5)+1)},
		{Name: "Chef 2", CookingTime: time.Second * time.Duration(rand.Intn(5)+1)},
		{Name: "Chef 3", CookingTime: time.Second * time.Duration(rand.Intn(5)+1)},
	}

	for i := range chefs {
		go chefs[i].Cook()
	}

	time.Sleep(time.Second * 10)
}
