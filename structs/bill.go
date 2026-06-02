package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
// we can also use a constructor function to create new bills. A constructor function is a function that returns a new instance of a struct. This can be useful for initializing the struct with default values or for performing any necessary setup before returning the struct. Here is an example of a constructor function for the bill struct:
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b

}
