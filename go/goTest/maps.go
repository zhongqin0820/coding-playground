package goTest

type Dictionay map[string]string

// Dictionary Error
type DictionayErr string

func (e DictionayErr) Error() string {
	return string(e)
}

//
const (
	ErrorFound      = DictionayErr("The keyword have already existed!")
	ErrorNotExisted = DictionayErr("The keyword didn't existed!")
)

// search a keyword in the dictionary
func (m Dictionay) Search(k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func (d Dictionay) Add(k string, v string) (err error) {
	if _, ok := d[k]; !ok {
		d[k] = v
		return nil
	}
	return ErrorFound
}

func (d Dictionay) Update(k string, v string) (err error) {
	if _, ok := d[k]; !ok {
		return ErrorNotExisted
	} else {
		d[k] = v
	}
	return nil
}

func (d Dictionay) Delete(k string) {
	delete(d, k)
}
