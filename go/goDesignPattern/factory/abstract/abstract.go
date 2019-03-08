package abstract

// definition of gateway interface and implementation of concrete brand of gateway
type Gateway interface {
	CreateInputParser() Parser
	CreateOutputEncoder() Encoder
}

type GatewayHuawei struct{}

func (g *GatewayHuawei) CreateInputParser() Parser {
	return &JsonParser{}
}

func (g *GatewayHuawei) CreateOutputEncoder() Encoder {
	return &JsonEncoder{}
}

type GatewayTencent struct{}

func (g *GatewayTencent) CreateInputParser() Parser {
	return &XmlParser{}
}

func (g *GatewayTencent) CreateOutputEncoder() Encoder {
	return &XmlEncoder{}
}

// definition of parser interface and implementation of concrete parser
type Parser interface {
	ParseData() string
}

type XmlParser struct{}

type JsonParser struct{}

func (x *XmlParser) ParseData() string {
	return "XmlParser parse"
}

func (j *JsonParser) ParseData() string {
	return "JsonParser parse"
}

// definition of encoder interface and implementation of concrete encoder
type Encoder interface {
	EncodeData() string
}

type XmlEncoder struct{}

type JsonEncoder struct{}

func (x *XmlEncoder) EncodeData() string {
	return "XmlEncoder encode"
}

func (j *JsonEncoder) EncodeData() string {
	return "JsonEncoder encode"
}
