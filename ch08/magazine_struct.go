package main

import (
	"fmt"
	"learning-go/whg/magazine"
)

func main() {
	address := magazine.Address{}
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"

	employee := magazine.Employee{}
	employee.HomeAddress = address
	fmt.Println(employee)
}
