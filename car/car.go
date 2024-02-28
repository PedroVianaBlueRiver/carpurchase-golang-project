package car

import "fmt"

type Car struct {
	brand string
	model string
	year  int
	price float64
}

func NewCar(brand, model string, year int, price float64) Car {
	return Car{
		brand: brand,
		model: model,
		year:  year,
		price: price,
	}
}

func NewListCar() map[int]Car {

	listCar := make(map[int]Car)
	listCar[1] = NewCar("JAC", "SEI 4 Pro", 2024, 35800.53)
	listCar[2] = NewCar("KIA", "Soul", 2024, 12000.53)
	listCar[3] = NewCar("Chevrolet", "Spark Classic", 2017, 10000.00)
	return listCar

}

func GetItem(listCar map[int]Car, idKey int) (Car, bool, string) {

	if _, ok := listCar[idKey]; !ok {
		return Car{}, false, fmt.Sprintf("Car with idKey: %v doesn't exist", idKey)
	}

	return listCar[idKey], true, "Success"
}

func AddItem(listCar map[int]Car, idKey int, item Car) (map[int]Car, bool, string) {
	if _, ok := listCar[idKey]; ok {
		return listCar, false, fmt.Sprintf("Car with idKey: %v already exists", idKey)
	}

	listCar[idKey] = item
	return listCar, true, "Success"
}

// Delete item of the map
func DeleteItem(listCar map[int]Car, id int) (map[int]Car, bool, string) {
	if _, ok := listCar[id]; !ok {
		return listCar, false, fmt.Sprintf("Car with id: %v doesn't exist", id)
	}
	delete(listCar, id)
	return listCar, true, ""
}
