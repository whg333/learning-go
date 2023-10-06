package main

import (
	"fmt"
	"learning-go/whg/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address.Street = "123 Oak St"
	subscriber.Address.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("subscriber\t=\t", subscriber)

	address := magazine.Address{
		Street: "456 Elm St",
		City:   "Portland",
	} // struct字面量
	address.State = "OR"
	address.PostalCode = "97222"

	employee := magazine.Employee{Name: "Joye Carr"}
	employee.HomeAddress = address
	fmt.Println("employee\t=\t", employee)
}
