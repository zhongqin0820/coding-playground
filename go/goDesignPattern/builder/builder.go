package builder

// define the build process of vehicles
type BuildProcess interface {
	SetSeats() BuildProcess
	SetWheels() BuildProcess
	SetNames() BuildProcess
	GetVehicle() VehicleProduct
}

// define the manufator of build
type Manufactor struct {
	builder BuildProcess
}

func (m *Manufactor) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *Manufactor) Build() {
	m.builder.SetNames().SetSeats().SetWheels()
}

// define the vehicle struct
type VehicleProduct struct {
	seat   int
	wheels int
	names  string
}

// define the CarBuilder, has a member of abstract VehicleProduct and implements the buildprocess interface
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.seat = 5
	return c
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.wheels = 4
	return c
}

func (c *CarBuilder) SetNames() BuildProcess {
	c.v.names = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// Another product of vehicle
type BikeBuilder struct {
	v VehicleProduct
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.seat = 2
	return c
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.wheels = 2
	return c
}

func (c *BikeBuilder) SetNames() BuildProcess {
	c.v.names = "Bike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}
