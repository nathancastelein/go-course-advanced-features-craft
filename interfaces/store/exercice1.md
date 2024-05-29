# A store of Compute and Storage

You're in charge to maintain a piece of software for your new online store to sell compute and storage.

You need to add a new product in your store, and you start to check the complexity of adding this feature.

Your predecessor already wrote some code, and you're not really happy with this implementation. 

It looks like things can become easier with interfaces, so you start by this small refactoring.

## Step 1: a less-anemic model

For now, your application provides three structures:
- Basket
- Compute
- Storage

Basket contains all the responsabilities: computing the price for each type of product, then computing the total price.

You might notice different rules to get the price of storage and compute. 

Storage has a business rule: the higher storage you have, the lower the unit price is.

Compute and Storage objects have no responsabilities at all for now.

Let's change this!

Add a new method `GetPrice() float64` on the Compute and Storage structure. Move the business rule to compute the price for each product in their respective structure.

Modify the code inside the `GetTotalPrice()` function to use the new methods.

## Step 2: a simple interface

Storage and Compute now share the same method `GetPrice() float64`.

Create a new interface `Pricer` with this method.

## Step 3: basket with pricer

Rewrite the Basket structure to use a slice of Pricer `[]Pricer` instead of `[]Compute` and `[]Storage`.

Remove the functions `AddStorage` and `AddCompute`. Add a new method `Add(pricer Pricer)` to add a product in the new slice.

Create a `basket_test.go` file to add unit tests and add this test:

```go
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
```