package main

import (
	"fmt"
	"learningsession/mapstructmodel"
	"learningsession/structmodel"
)

func main() {

	carmodel := structmodel.NewCar("KIA", "Soul", 2024, 12000.53)

	fmt.Println("Print car data #1:")
	fmt.Println(carmodel)

	carmodel = structmodel.NewCar("KIA", "Soul", 2024, 12000.53)
	fmt.Println("Print car data #2:")
	fmt.Println(carmodel)

	fmt.Println("Print car map:")
	fmt.Println(mapstructmodel.NewListCar())

}
