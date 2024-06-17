// Client code

package main

import "fmt"

func main() {
	pizza := &VeggieMania{}

	// Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of VeggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
