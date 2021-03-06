package main

import (
	"fmt"
	"strings"
)

type Customer struct {
	firstName string
	lastName  string
	address   string
	zipCode   int
}

type Order struct {
	Customer
	id   int
	name string
	cost float64
}

func (order Order) checkPayment(payment float64) (isEnough bool) {
	isEnough = false
	if payment >= order.cost {
		isEnough = true
	}
	return
}

func (order *Order) changeAddress(newAddress string)  {
	order.address = newAddress
}

func isBelorussian(order Order) {
	if strings.Contains(order.address,"Belarus") {
		fmt.Println("Заплати налог!!!")
	}
	
}

func main() {
	firstOrder := Order{Customer{"Andrey", "Mayorau", "Belarus Gomel Belitca", 111111}, 228, "laptop", 789.23}
	fmt.Printf("%+v\n",firstOrder)
	fmt.Println(firstOrder.checkPayment(900))
	isBelorussian(firstOrder)
	firstOrder.changeAddress("Belarus Gomel Makaenka")
	fmt.Printf("%+v\n",firstOrder)
}
