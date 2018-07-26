package testBasic

import (
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
