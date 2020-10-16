package model

type Coffee struct {
	water Ingredient
	milk  Ingredient
	sugar Ingredient
	beans Ingredient
	price float64
}

func NewCoffee(w, m, s, b Ingredient, price float64) *Coffee {
	return &Coffee{
		water: w,
		milk:  m,
		sugar: s,
		beans: b,
		price: price,
	}
}
