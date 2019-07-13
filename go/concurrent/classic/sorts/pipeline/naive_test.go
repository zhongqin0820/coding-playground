package pipeline

import (
	"bufio"
	"os"
	"testing"
)

func TestInMemSort(t *testing.T) {
	a := []int{1, 7, 9, 3, 5}
	b := []int{2, 6, 8, 4, 10}
	Init() //init the starTime
	t.Run("array source: sort pipeline", func(t *testing.T) {
		PrintData(InMemSort(ArraySource(a...)))
	})
	Init() //init the starTime
	t.Run("array source: merge pipeline", func(t *testing.T) {
		PrintData(Merge(InMemSort(ArraySource(a...)), InMemSort(ArraySource(b...))))
	})
}

func TestReaderSource(t *testing.T) {
	const fileName = "small.in"
	const outfile = "small.out"
	const n = 64
	const chunkCount = 2
	const chunkSize = n * 8 / chunkCount

	GenData(fileName, n)
	Init()

	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, InMemSort(source))
	}
	p := MergeN(sortResults...)
	WriteToFile(p, outfile)
	PrintFile(outfile)
}
