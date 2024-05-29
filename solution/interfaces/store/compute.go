package store

var (
	_ Pricer = &Compute{}
)

type Compute struct {
	NumberOfVCPUs int
}

func (c Compute) GetPrice() float64 {
	unitPrice := 3.0
	return float64(c.NumberOfVCPUs) * unitPrice
}
