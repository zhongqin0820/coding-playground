package external

import (
	"bufio"
	"os"
	"strconv"

	"github.com/zhongqin0820/coding-playground/go/concurrent/classic/sorts/pipeline"
)

func LaunchPipeline(fileName string, fileSize int, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()
	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func LaunchNetworkPipeline(fileName string, fileSize int, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()
	sortAddr := make([]string, 0)
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(i+7000)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}
