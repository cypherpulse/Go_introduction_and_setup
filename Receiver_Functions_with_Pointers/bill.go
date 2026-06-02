package main

import (
	"fmt"
	"os"
)

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

// formart the bill\
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25s.... $%.2f\n", k+":", v)
		total += v
	}
	//total
	fs += fmt.Sprintf("%-25v.... $%.2f\n", "Tip:", b.tip)
	total += b.tip
	fs += fmt.Sprintf("%-25v.... $%.2f\n", "Total:", total)
	return fs
}

//update tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bill/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
