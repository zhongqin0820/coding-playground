package json

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

// 定义mongo字段
type MongoConfig struct {
	MongoAddr      string
	MongoPoolLimit int
	MongoDb        string
	MongoCol       string
}

//定义config字段
type Config struct {
	Addr  string
	Mongo MongoConfig
}

// 定义读取结构
type JsonStruct struct{}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
func (js *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//转码json文件
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func TestJson(t *testing.T) {
	// init
	JsonParse := NewJsonStruct()
	v := Config{}
	// read
	JsonParse.Load("./demo.json", &v)
	t.Log(v.Addr)
	t.Log(v.Mongo.MongoDb)
}
