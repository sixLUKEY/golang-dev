// Concrete implementation

package main

import "fmt"

type HP struct{}

func (p *HP) PrintFile() {
	fmt.Println("Printing from HP printer!")
}
