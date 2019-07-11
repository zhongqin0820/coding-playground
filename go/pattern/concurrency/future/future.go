package future

// executed if everything went well
type SuccessFunc func(string)

// receive an error when something goes wrong
type FailFunc func(error)

// defines the operation we want to perform.
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

// accept a SuccessFunc function that will be executed if everything goes correctly
// and return the same pointer to the future object (so we can keep chaining)
func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

// accept a FailFunc function to later return the pointer.
func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

// pass an ExecuteStringFunc type (a function that accepts nothing and returns a string or an error)
func (s *MaybeString) Execute(f ExecuteStringFunc) {
	// launch the Goroutine that executes the f method (an ExecuteStringFunc)
	// and takes its result
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}
