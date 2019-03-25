package abstract

import (
	"testing"
)

func TestAbstract(t *testing.T) {
	var gateway Gateway
	var parser Parser
	var encoder Encoder
	gateway = &GatewayHuawei{}
	parser = gateway.CreateInputParser()
	encoder = gateway.CreateOutputEncoder()
	t.Log(parser.ParseData())
	t.Log(encoder.EncodeData())
}
