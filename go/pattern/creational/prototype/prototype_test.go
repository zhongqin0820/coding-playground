package prototype

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestPrototype(t *testing.T) {
	var role Role
	role = NewRole()
	role.Show()
	//
	var another Role = role.Clone()
	another.SetTall(160)
	another.Show()
	role.Show()
	t.Run("same type", func(t *testing.T) {
		// correct
		if reflect.DeepEqual(role, another) {
			t.Logf("same: %p, %p\n", role, another)
		} else {
			t.Logf("different: %p, %p\n", &role, &another)
		}
		// wrong
		if unsafe.Pointer(&role) == unsafe.Pointer(&another) {
			t.Logf("same: %p,%p\n", &role, &another)
		} else {
			t.Logf("different: %p,%p\n", &role, &another)
		}
	})
}
