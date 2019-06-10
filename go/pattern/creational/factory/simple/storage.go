package simple

import "log"

type StorageType int

// define the options for caller of the simple factory
const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
	NotExistedStorage
)

// NewStore is the factory function, which return a concrete product at runtime
func NewStore(t StorageType) Store {
	switch t {
	case DiskStorage:
		return newDiskStorage()
	case TempStorage:
		return newTempStorage()
	case MemoryStorage:
		return newMemoryStorage()
	default:
		return nil
	}
}

// Store is an abstract product interface
type Store interface {
	Save(s string)
}

// disk is a concrete product realizes from Store interface
type disk struct{}

func newDiskStorage() disk   { return disk{} }
func (d disk) Save(s string) { log.Println(s) }

// temp is a concrete product realizes from Store interface
type temp struct{}

func newTempStorage() temp   { return temp{} }
func (t temp) Save(s string) { log.Println(s) }

// memory is a concrete product realizes from Store interface
type memory struct{}

func newMemoryStorage() memory { return memory{} }
func (m memory) Save(s string) { log.Println(s) }
