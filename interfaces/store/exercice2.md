# Stub pricer

Let's now create a stub pricer to test only the complexity of GetTotalPrice, not of its products!

In your test file, create a new structure `StubPricer`.

Add a `GetPrice() float64` method on this structure, so it fits the `Pricer` interface.
This method just returns `1.0`.

Rewrite the unit test to use the `StubPricer` instead of `Compute` and `Storage`!

Example of test: given a basket with four StubPricer, when we compute the total price then the total price should be equal to four.