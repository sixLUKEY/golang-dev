// Concrete implementation

package main

import "fmt"

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing from Epson printer!")
}
