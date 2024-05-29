package store

var (
	_ Pricer = &Storage{}
)

type Storage struct {
	QuantityInGB float64
}

func (s Storage) GetPrice() float64 {
	var unitPrice float64
	switch {
	case s.QuantityInGB >= 1024:
		unitPrice = 1.5
	case s.QuantityInGB >= 1024*10:
		unitPrice = 1.0
	default:
		unitPrice = 2.0
	}

	return unitPrice * s.QuantityInGB
}
