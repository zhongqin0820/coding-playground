package builder

import (
	"fmt"
	"testing"
)

func TestCar(t *testing.T) {
	var m = Manufactor{}
	var b = &CarBuilder{}
	m.SetBuilder(b)
	m.Build()
	var c = b.GetVehicle()
	if c.names != "Car" {
		fmt.Println(c.names)
		t.Error("Name error:", c.names)
	} else {
		fmt.Println("Correct names: ", c.names)
	}

}
