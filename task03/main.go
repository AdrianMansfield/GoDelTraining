package main

import "fmt"

type dataInDBCreater interface {
	create() bool
}

type outputer interface {
	perfectOutput() string
}

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

func (customer Customer) create() (result bool){
	fmt.Println("Start transaction...")
	result = false
	//проверка на отсутсвие такого заказчика в базе
	if true {
		fmt.Println("Customer saved successful. Transaction committed!")
		result = true
		//логика по добавлению в базу заказчика
	} else {
		fmt.Println("Customer saved failed!")
	}
	return
}

func (order Order) create() (result bool){
	fmt.Println("Start transaction...")
	result = false
	//проверка на отсутсвие такого заказа в базе
	if true {
		fmt.Println("Order saved successful. Transaction committed!")
		result = true
		//логика по добавлению в базу заказа
	} else {
		fmt.Println("Order saved failed!")
	}
	return
}

func (order Order) perfectOutput() string{
	return fmt.Sprintf("имя: %s; фамилия: %s; адрес: %s; индекс: %d; номер заказа: %d; имя товара: %s; цена: %.2f;",order.firstName, order.lastName, order.address, order.zipCode, order.id, order.name, order.cost)
}

func main() {
	//переделанный первый таск, с другими функциями
	laptopOrder := Order{Customer{"Andrey", "Mayorau", "Belarus Gomel Belitca", 111111}, 228, "laptop", 789.23}
	laptopOrder.create()

	laptopCustomer := Customer{"Andrey", "Mayorau", "Belarus Gomel Belitca", 111111}
	laptopCustomer.create()

	//второй таск, с интерфейсом сделанным по примеру stringer
	ordersArray := [4]Order{
		{Customer{"Andrey", "Mayorau", "Belarus Gomel Belitca", 111111}, 228, "laptop", 789.23},
		{Customer{"Julia", "Pritychenko", "Belarus Gomel Makaenka", 100156}, 456, "cellphone", 290.78},
		{Customer{"Oleg", "Dedik", "Lithuania Kaunas Dworaka", 200987}, 470, "ticket", 50.56},
		{Customer{"Zaretskiy", "Roman", "Belarus Minsk Nezaleznasti", 101567}, 681, "crisps", 10.45},
	}
	for i := 0; i < len(ordersArray); i++ {
		fmt.Println(ordersArray[i].perfectOutput())
	}
}
