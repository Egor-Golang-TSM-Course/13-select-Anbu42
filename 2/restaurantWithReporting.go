package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Restaurant struct {
	Name  string
	Sales int
}

func (r *Restaurant) generateSales(stopCh <-chan struct{}, salesCh chan<- Restaurant) {

	for {
		select {
		case <-stopCh:
			return
		default:
			r.Sales = rand.Intn(10) + 1

			salesCh <- *r

			time.Sleep(time.Second)
		}
	}
}

func centralOffice(salesCh <-chan Restaurant, stopCh <-chan struct{}) {

	totalSales := 0

	for {
		select {
		case restaurant := <-salesCh:
			totalSales += restaurant.Sales
			fmt.Printf("Центральный офис: полученыо %v продаж от ресторана %s. Общее количество продаж: %d\n", restaurant.Sales, restaurant.Name, totalSales)
		case <-stopCh:
			return
		}
	}
}

func main() {
	restaurant1 := Restaurant{Name: "Ресторан 1"}
	restaurant2 := Restaurant{Name: "Ресторан 2"}

	stopCh := make(chan struct{})
	salesCh := make(chan Restaurant)

	go restaurant1.generateSales(stopCh, salesCh)
	go restaurant2.generateSales(stopCh, salesCh)

	go centralOffice(salesCh, stopCh)

	time.Sleep(time.Second * 4)

	close(stopCh)

	close(salesCh)
}
