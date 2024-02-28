package main

import (
	"fmt"
	"learningsession/car"
)

func main() {

	// carmodel := car.NewCar("KIA", "Soul", 2024, 12000.53)

	// fmt.Println("Print car data #1:")
	// fmt.Println(carmodel)

	// carmodel = car.NewCar("KIA", "Soul", 2024, 12000.53)
	// fmt.Println("Print car data #2:")
	// fmt.Println(carmodel)

	fmt.Println(" car map:")
	mapcar := car.NewListCar()
	fmt.Println(mapcar)

	carresponse, ok, msn := car.GetItem(mapcar, 8)

	fmt.Println("Get item from car map **********************************")
	if ok {
		fmt.Println(carresponse)
	} else {
		fmt.Println(msn)
	}

	addrespose, okAdd, msnAdd := car.AddItem(mapcar, 9, car.NewCar("Nissan", "Versa", 2020, 90000))
	fmt.Println("Add item to car map **********************************")
	if okAdd {
		for k, v := range addrespose {
			fmt.Println("**************************************************")
			fmt.Printf("key[%v] value[%v]\n", k, v)
		}
	} else {
		fmt.Println(msnAdd)
	}

	fmt.Println("Remove item to car map **********************************")
	removeresponse, okRemove, msnRemove := car.DeleteItem(mapcar, 1)
	if !okRemove {
		fmt.Println(msnRemove)
	} else {
		fmt.Println("DeleteItem() Updated Map: ", removeresponse)
	}

}
