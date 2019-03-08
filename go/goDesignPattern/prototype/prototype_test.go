package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	var role Role
	role = &Chinese{HeadColor: "black", EyesColor: "yellow", Tall: 190}
	role.Show()
	another := role.Clone()
	another.SetTall(160)
	another.Show()
	role.Show()
}
