package structmodel

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
