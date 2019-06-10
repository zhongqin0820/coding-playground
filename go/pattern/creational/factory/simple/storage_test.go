package simple

import (
	"fmt"
	"testing"
)

type testCase struct {
	name    string
	storage StorageType
}

func TestStorage(t *testing.T) {
	var ts []testCase = []testCase{
		{"DiskStorage", DiskStorage},
		{"TempStorage", TempStorage},
		{"MemoryStorage", MemoryStorage},
		{"NotExistedStorage", NotExistedStorage},
	}
	// table test
	for _, v := range ts {
		t.Run(v.name, func(t *testing.T) {
			storage := NewStore(v.storage)
			if storage == nil {
				t.Errorf("Type DiskStorage not implemented yet")
			} else {
				storage.Save(fmt.Sprint("Save from ", v.name))
			}
		})
	}
}
