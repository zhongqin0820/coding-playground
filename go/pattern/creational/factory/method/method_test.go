package method

import (
	"testing"
)

func TestMethod(t *testing.T) {
	var factory Factory
	var server Server
	// call an implementation of Factory interface
	factory = &SoaFactory{}
	// get correspongding product
	server = factory.CreateServer()
	name := server.GetServerType()
	if name != "SOA" {
		t.Error("Wrong type!")
	}

}
