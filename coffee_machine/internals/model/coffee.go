package model

type Coffee struct {
	Water uint
	Milk  uint
	Sugar uint
	Beans uint
	Price float64
}

//type drinks struct {
//	drinks map[string]Coffee
//}

func NewCoffee(water, milk, sugar, beans uint, price float64) *Coffee {
	return &Coffee{
		Water: water,
		Milk:  milk,
		Sugar: sugar,
		Beans: beans,
		Price: price,
	}
}
