package main

import (
	"github.com/zhongqin0820/coding-playground/go/concurrent/classic/sorts/external"
	"github.com/zhongqin0820/coding-playground/go/concurrent/classic/sorts/pipeline"
)

func main() {
	const (
		fileName   = "large.in"
		fileSize   = 80000000
		chunkCount = 4
		outfile1   = "large1.out"
		outfile2   = "large2.out"
	)
	pipeline.GenData(fileName, fileSize/8)

	//
	p := external.LaunchPipeline(fileName, fileSize, chunkCount)
	pipeline.WriteToFile(p, outfile1)
	pipeline.PrintFile(outfile1)

	//
	p = external.LaunchNetworkPipeline(fileName, fileSize, chunkCount)
	pipeline.WriteToFile(p, outfile2)
	pipeline.PrintFile(outfile2)
}
