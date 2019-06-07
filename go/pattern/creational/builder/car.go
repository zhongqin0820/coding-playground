package builder

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

func NewBuilder(b Builder) Builder {
	return b
}
