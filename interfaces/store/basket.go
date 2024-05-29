package store

type Basket struct {
	Compute []Compute
	Storage []Storage
}

type Compute struct {
	NumberOfVCPUs int
}

type Storage struct {
	QuantityInGB float64
}

func (b Basket) GetTotalPrice() float64 {
	totalPrice := 0.0

	for _, computeProduct := range b.Compute {
		totalPrice += float64(computeProduct.NumberOfVCPUs) * 3.0
	}

	for _, storageProduct := range b.Storage {
		var unitPrice float64
		switch {
		case storageProduct.QuantityInGB >= 1024*10:
			unitPrice = 1.0
		case storageProduct.QuantityInGB >= 1024:
			unitPrice = 1.5
		default:
			unitPrice = 2.0
		}

		totalPrice += storageProduct.QuantityInGB * unitPrice
	}

	return totalPrice
}

func (b *Basket) AddStorage(storage Storage) {
	b.Storage = append(b.Storage, storage)
}

func (b *Basket) AddCompute(compute Compute) {
	b.Compute = append(b.Compute, compute)
}
