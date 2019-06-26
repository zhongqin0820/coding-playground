package simple

import (
	"fmt"
	"testing"
)

type testCase struct {
	name    string
	storage StorageType
}

// go test -v -run=TestStorage
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
				t.Log("NotExistedStorage")
			} else {
				storage.Save(fmt.Sprint("Save from ", v.name))
			}
		})
	}
}
