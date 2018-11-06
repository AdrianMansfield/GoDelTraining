package main

import "fmt"

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

func main() {
	//array
	fmt.Println("array")
	ordersArray := [4]Order{
		{Customer{"Andrey", "Mayorau", "Belarus Gomel Belitca", 111111}, 228, "laptop", 789.23},
		{Customer{"Julia", "Pritychenko", "Belarus Gomel Makaenka", 100156}, 456, "cellphone", 290.78},
		{Customer{"Oleg", "Dedik", "Lithuania Kaunas Dworaka", 200987}, 470, "ticket", 50.56},
		{Customer{"Zaretskiy", "Roman", "Belarus Minsk Nezaleznasti", 101567}, 681, "crisps", 10.45},
	}
	for i := 0; i < len(ordersArray); i++ {
		fmt.Printf("%d | %+v\n", i, ordersArray[i])
	}

	//slice
	fmt.Println("slice")
	ordersSlice := ordersArray[:]

	for index, element := range ordersSlice {
		fmt.Printf("%d | %+v\n", index, element)
	}
}
