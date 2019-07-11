package future

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// waits for a second, then prints a message, sets the test as failed
// and lets the WaitGroup continue by calling its Done() method
func timeout(t *testing.T, wg *sync.WaitGroup) {
	// the execution time is now nearly zero, so our timeouts have not been executed
	// actually, they were executed, but the tests already finished and their result was already stated.
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}

// use closures to introduce a context into this type of function.
// it takes a string as an argument and returns an ExecuteStringFunc method
// that returns the previous argument with the suffix Closure!
func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}

func TestStringOrError_Execute(t *testing.T) {
	ms := &MaybeString{}
	var wg sync.WaitGroup
	wg.Add(1)
	// Timeout
	go timeout(t, &wg)
	//
	t.Run("Success result", func(t *testing.T) {
		ms.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		//
		ms.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})
	//
	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		ms.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		//
		ms.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})
		wg.Wait()
	})
	// a test that uses the closure
	t.Run("Closure Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		//Timeout!
		go timeout(t, &wg)
		ms.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		ms.Execute(setContext("Hello"))
		wg.Wait()
	})
}
