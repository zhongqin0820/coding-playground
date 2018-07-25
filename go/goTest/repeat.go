package goTest

func Repeat(ch string, ind int) (res string) {
	var b []byte
	b0 := []byte(ch)
	for i := 0; i < ind; i++ {
		b = append(b, b0...)
	}
	return string(b)
}
