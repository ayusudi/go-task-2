package main

import (
	"fmt"
	"sync"
)

type Calculator struct {
	mutex sync.Mutex
	result int
}

func (c *Calculator) Add(x int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.result += x
}

func (c *Calculator) Subtract(x int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.result -= x
}

func (c *Calculator) Multiply(x int) {	
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.result *= x
}

func (c *Calculator) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.result
}

func (c *Calculator) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.result = 0
}

func main() {
	calc := &Calculator{}

	fmt.Println("Initial value:", calc.GetValue())

	calc.Add(10)
	fmt.Println("Value after adding 10:", calc.GetValue())

	calc.Subtract(5)
	fmt.Println("Value after subtracting 5:", calc.GetValue())

	calc.Multiply(2)
	fmt.Println("Value after multiplying by 2:", calc.GetValue())

	calc.Reset()
	fmt.Println("Value after reset:", calc.GetValue())
}