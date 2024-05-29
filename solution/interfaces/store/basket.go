package store

type Pricer interface {
	GetPrice() float64
}

type Basket struct {
	Products []Pricer
}

func (b Basket) GetTotalPrice() float64 {
	totalPrice := 0.0
	for _, product := range b.Products {
		totalPrice += product.GetPrice()
	}

	return totalPrice
}

func (b *Basket) AddProduct(product Pricer) {
	b.Products = append(b.Products, product)
}
