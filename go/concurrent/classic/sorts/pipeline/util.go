package pipeline

import (
	"bufio"
	"os"
	"time"
)

// setup timer
var startTime time.Time

func Init() {
	startTime = time.Now()
}

// print utilities
func PrintData(in <-chan int) {
	for v := range in {
		print(v, " ")
	}
	println()
}

func PrintFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		println(v)
		count++
		if count >= 3 {
			break
		}
	}
}

func GenData(filename string, count int) {
	// create the source data
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := RandomSource(count)
	writer := bufio.NewWriter(file)
	WriterSink(writer, p)
	defer writer.Flush()
}

func WriteToFile(p <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	WriterSink(writer, p)
}
