package store

import "testing"

func TestGetTotalPrice(t *testing.T) {
	// Arrange
	compute := Compute{NumberOfVCPUs: 4}   // total price is 12
	storage := Storage{QuantityInGB: 2000} // total price is 3000
	basket := Basket{Products: []Pricer{compute, storage}}
	expectedTotalPrice := 3012.0

	// Act
	totalPrice := basket.GetTotalPrice()

	// Assert
	if expectedTotalPrice != totalPrice {
		t.Fatalf("expected total price %.2f, got %.2f", expectedTotalPrice, totalPrice)
	}
}

type StubPricer struct{}

func (StubPricer) GetPrice() float64 {
	return 1.0
}

func TestGetTotalPriceWithStub(t *testing.T) {
	// Arrange
	basket := Basket{Products: []Pricer{StubPricer{}, StubPricer{}, StubPricer{}, StubPricer{}}}
	expectedTotalPrice := 4.0

	// Act
	totalPrice := basket.GetTotalPrice()

	// Assert
	if expectedTotalPrice != totalPrice {
		t.Fatalf("expected total price %.2f, got %.2f", expectedTotalPrice, totalPrice)
	}
}
