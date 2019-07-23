package basic

import (
	"log"
	"reflect"
	"testing"
)

func TestMaps(t *testing.T) {
	dictionary := Dictionay{"test": "this is just a test"}
	got, ok := dictionary.Search("test")
	want := "this is just a test"
	if ok && got != want {
		t.Errorf("want %s, but got %s", want, got)
	} else if ok == false {
		t.Errorf("didn't have this key")
	}
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		d := Dictionay{"test": "this is just a test"}
		res := d.Add("test1", "this is a test")
		if res != nil {
			t.Error(res)
		}
	})
}

func TestUpdate(t *testing.T) {
	// t.Run("not existed", func(t *testing.T) {
	// 	d := Dictionay{"test": "this is a test"}
	// 	res := d.Update("test1", "haha")
	// 	if res != nil {
	// 		t.Error(res)
	// 	}
	// })
	t.Run("existed", func(t *testing.T) {
		d := Dictionay{"test": "this is a test"}
		res := d.Update("test", "haha")
		if res != nil {
			t.Error(res)
		}
	})
}

func TestDelete(t *testing.T) {
	d := Dictionay{"test": "this is a test"}
	d.Delete("test")
}

// test map keys by using interface{} as key & value
func TestMapKeys(t *testing.T) {
	m := map[interface{}]interface{}{}
	m[1] = 1
	m["s"] = "s"
	m[0.3] = 0.3
	// m[[]int{1, 2, 3}] = []int{1, 2, 3}
	// panic: runtime error: hash of unhashable type []int
	for k, v := range m {
		switch k.(type) {
		case int:
			log.Println(k.(int), v.(int))
		case string:
			log.Println(k.(string), v.(string))
		default:
			log.Println("unknown type: ", reflect.TypeOf(k).String())
		}
	}
}

func TestMapLen(t *testing.T) {
	m := make(map[interface{}]interface{}, 20)
	m[1] = 1
	m["s"] = "s"
	m[0.3] = 0.3
	log.Println(len(m))
}
