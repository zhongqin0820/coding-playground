package prototype

import (
	"log"
)

// defition of interface: some attri modify method
type Role interface {
	Clone() Role
	SetHeadColor(string)
	SetEyesColor(string)
	SetTall(int)
	Show()
}

func NewRole() Role {
	return &Chinese{HeadColor: "black", EyesColor: "yellow", Tall: 190}
}

// concrete product: implements the attri modify method and call Clone() to have a new object at runtime
type Chinese struct {
	HeadColor string
	EyesColor string
	Tall      int
}

func (c *Chinese) Clone() Role {
	return &Chinese{HeadColor: c.HeadColor, EyesColor: c.EyesColor, Tall: c.Tall}
}

func (c *Chinese) SetHeadColor(s string) {
	c.HeadColor = s
}

func (c *Chinese) SetEyesColor(s string) {
	c.EyesColor = s
}

func (c *Chinese) SetTall(s int) {
	c.Tall = s
}

func (c *Chinese) Show() {
	log.Printf("%s,%s,%d\n", c.HeadColor, c.EyesColor, c.Tall)
}
