package mapstructmodel

import "learningsession/structmodel"

func NewListCar() map[int]structmodel.Car {

	listCar := make(map[int]structmodel.Car)
	listCar[1] = structmodel.NewCar("JAC", "SEI 4 Pro", 2024, 35800.53)
	listCar[2] = structmodel.NewCar("KIA", "Soul", 2024, 12000.53)
	listCar[3] = structmodel.NewCar("Chevrolet", "Spark Classic", 2017, 10000.00)
	return listCar
}
