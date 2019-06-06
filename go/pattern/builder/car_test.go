package builder

import (
	"log"
	"testing"
)

// the implementation of Builder interface
type Car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (c Car) Paint(color Color) Builder {
	c.color = color
	return c
}
func (c Car) Wheels(w Wheels) Builder {
	c.wheels = w
	return c
}
func (c Car) TopSpeed(s Speed) Builder {
	c.speed = s
	return c
}
func (c Car) Build() Interface {
	switch c.wheels {
	case SportsWheels:
		return SportsCar{}
	case SteelWheels:
		return FamilyCar{}
	}
	return nil
}

// the implementation of Interface interface
type FamilyCar struct{}

func (f FamilyCar) Drive() error {
	log.Println("FamilyCar Drive()")
	return nil
}
func (f FamilyCar) Stop() error {
	log.Println("FamilyCar Stop()")
	return nil
}

// the implementation of Interface interface
type SportsCar struct{}

func (s SportsCar) Drive() error {
	log.Println("SportsCar Drive()")
	return nil
}

func (s SportsCar) Stop() error {
	log.Println("SportsCar Stop()")
	return nil
}

// the usage
func TestCarBuilder(t *testing.T) {
	assembly := NewBuilder(Car{}).Paint(RedColor)

	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()
	familyCar.Stop()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
	sportsCar.Stop()

}
