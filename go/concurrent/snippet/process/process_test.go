package process

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"
)

// https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter10/10.1.md

func MyPrint(p []byte, err error) {
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The output: %q is %s\n", os.Args[1], p)
}

func FillStd(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out = new(bytes.Buffer)
	cmd.Stdout = out
	cmd.Stderr = out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), err
}

func UseOutput(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}

func UsePipe(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	var out = make([]byte, 0, 1024)
	for {
		tmp := make([]byte, 128)
		//从控制台读数据到缓存中，再将缓存内容放入out中
		n, err := stdout.Read(tmp)
		out = append(out, tmp[:n]...)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		return nil, err
	}
	return out, err
}

func TestProcessOutput(t *testing.T) {
	argNum := len(os.Args)
	t.Log(argNum)
	if argNum < 2 {
		log.Printf("Usage:%s command\n", os.Args[0])
		os.Exit(1) // exit the program
	}
	arg := []string{}
	if argNum > 2 {
		arg = os.Args[2:]
	}
	t.Log(arg)
	//
	t.Run("FillStd", func(t *testing.T) {
		MyPrint(FillStd(os.Args[1], arg...))
	})
	//
	t.Run("UseOutput", func(t *testing.T) {
		MyPrint(UseOutput(os.Args[1], arg...))
	})
	//
	t.Run("UsePipe", func(t *testing.T) {
		MyPrint(UsePipe(os.Args[1], arg...))
	})
}
