package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// User input
func exp1() {
	var firstName, lastName, s string
	var i int
	var f float32
	fmt.Scanln(&firstName, &lastName)
	log.Println(firstName, lastName)
	fmt.Scanf("%s %s", &firstName, &lastName)
	log.Println(firstName, lastName)
	n, err := fmt.Sscanf("56.12 / 5212 / Go", "%f / %d / %s", &f, &i, &s)
	log.Println(n, err)
	log.Println(f, i, s)
	return
}

// User input function 2
func exp2() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input: ")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		log.Println("Your input is : ", input)
	}
	return
}

// file io
func exp3() {
	inputFile, inputError := os.Open("test.txt")
	defer inputFile.Close()
	if inputError != nil {
		log.Println("An error occurred")
		return
	}
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		if readError == io.EOF || readError != nil {
			return
		}
		log.Println("The input was:", inputString)
	}
}

func exp4() {
	inputFile := "test.txt"
	outputFile := "test_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Println(os.Stderr, "err is: ", err)
	}
	log.Println(string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func exp5() {
	file, err := os.Open("test.txt")
	defer file.Close()
	if err != nil {
		log.Panicln("err is :", err)
	}
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	log.Println(col1)
	log.Println(col2)
	log.Println(col3)
	return
}

func exp6() {
	filename := filepath.Base(".")
	log.Println(filename)
	return
}

func main() {
	// exp1()
	// exp2()
	// exp3()
	// exp4()
	// exp5()
	// exp6()

	fmt.Println("...")
}
