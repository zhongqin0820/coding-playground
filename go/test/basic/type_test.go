package basic

import (
	"log"
	"reflect"
	"testing"
)

type instanceType struct {
	fisrtname string
	lastname  string
}

func TestTypeAssertionUpdate(t *testing.T) {
	// 已知数据类型进行断言

	// 定义及初始化接口实例：
	var instance interface{} = instanceType{"xx", "yy"}
	// 接口实例进行类型断言得到接口中的类型实例：
	var value instanceType = instance.(instanceType)
	// 类型实例调用（获取具体值）：
	log.Println(value)

	// 未知数据类型进行断言(reflect包)

	// 获取接口中的实例的类型
	log.Println(reflect.ValueOf(instance).Type(), reflect.TypeOf(instance))
	// 获取接口中的实例内容：
	log.Println(reflect.ValueOf(instance).Interface())
	for i := 0; i < reflect.TypeOf(instance).NumField(); i++ {
		log.Print(reflect.TypeOf(instance).Field(i).Name)
		log.Print(reflect.TypeOf(instance).Field(i).Type) //或reflect.ValueOf(instance).Field(i).Type()
		log.Print(reflect.ValueOf(instance).Field(i))
	}
}

//
func TestTypeAssertion(t *testing.T) {
	type xxx int
	type yyy = int
	var a xxx = 3
	log.Println(int(a))
	// var b yyy = 3
	var c interface{} = a
	switch c.(type) {
	case xxx:
		log.Println("xxx", int(c.(xxx)))
	// case yyy: //case yyy conflicts with case int
	// log.Println("yyy")
	case int:
		log.Println("int")
	}
	log.Println(reflect.ValueOf(c).Type(), reflect.TypeOf(c).Kind())
}

// ref: https://juejin.im/post/5a75a4fb5188257a82110544
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	log.Println("ReflectCallFunc")
}

func TestTypeStruct(t *testing.T) {
	user := User{1, "a", 3}
	// definition
	GetFiledAndMethod := func(input interface{}) {
		getType := reflect.TypeOf(input)
		log.Println(getType.Name())        //User
		getValue := reflect.ValueOf(input) //{1 a 3}
		log.Println(getValue)
		// iterate to get the corresponding field type and value
		for i := 0; i < getType.NumField(); i++ {
			field := getType.Field(i)              //reflect.Type ==> type
			value := getValue.Field(i).Interface() //reflect.Value ==> value
			log.Println(field.Name, field.Type, value)
			// Id int 1
			// Name string a
			// Age int 3
		}
		// iterate to get the corresponding method of the concrete type
		for i := 0; i < getType.NumMethod(); i++ {
			m := getType.Method(i)
			log.Println(m.Name, m.Type)
			// ReflectCallFunc func(basic.User)
		}
	}
	// call GetFiledAndMethod
	GetFiledAndMethod(user)
}

// ref: https://blog.csdn.net/Tovids/article/details/77880789
func TestTypeAndReflect(t *testing.T) {
	// type
	t.Run("Type", func(t *testing.T) {
		f := func(i interface{}) {
			switch i.(type) {
			case int, int16, int32, int64:
				log.Println("int")
			case uint, uint16, uint32, uint64:
				log.Println("uint")
			case float32, float64:
				log.Println("float")
			case complex64, complex128:
				log.Println("complex")
			case byte:
				log.Println("char")
			case string:
				log.Println("string")
			case error:
				log.Println("error")
			default:
				log.Println("a unknown type")
			}
		}
		v := 3
		f(v)
	})
	t.Run("Reflect", func(t *testing.T) {
		v := 3
		log.Println("type:", reflect.TypeOf(v))
	})
	// value
	t.Run("Type", func(t *testing.T) {
		f := func(i interface{}) {
			switch i.(type) {
			case int:
				log.Println(i.(int))
			case uint64:
				log.Println(i.(uint64))
			case float64:
				log.Println(i.(float64))
			case complex128:
				log.Println(i.(complex128))
			case byte:
				log.Println(i.(byte))
			case string:
				log.Println(i.(string))
			case error:
				log.Println(i.(error))
			default:
				log.Println("a unknown type")
			}
		}
		var v int = 1<<63 - 1
		// log.Println(math.MaxInt32)
		f(v)
	})
	t.Run("Reflect", func(t *testing.T) {
		v := 3
		log.Println("value:", reflect.ValueOf(v))
	})
}

func TestPointers(t *testing.T) {
	// var i uintptr
	//  cannot use &j (type *uint) as type uintptr in assignment
	var i *uint
	var j uint = 3
	i = &j
	log.Println(i)
	log.Println(&i)
	log.Println(*i)
}

type Duck interface {
	Quack()
	SetName(string)
}
type Rubber struct {
	name string
}

func (r *Rubber) Quack() {
	log.Println(r.name)
}

func (r *Rubber) SetName(s string) {
	r.name = s
}

func TestInterfaceStruct(t *testing.T) {
	var d Duck = &Rubber{name: "hi"}
	d.SetName("hai")
	// if comment SetName related codes
	// will raise a compile error
	// *Rubber does not implement Duck (missing SetName method)
	d.Quack()
}
