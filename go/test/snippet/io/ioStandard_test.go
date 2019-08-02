package snippet

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

func TestFileWriter(t *testing.T) {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	// 除了自己创建一个文件，还可以用os.Stdout，os.Stderr作为写出对象
	file, err := os.Create("./proverbs.txt") //打开一个文件写，返回*os.File实现了io.Rearder和io.Writer
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close() //确保文件最后被关闭

	for _, p := range proverbs {
		n, err := file.Write([]byte(p)) //向*os.File中写入数据
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			log.Println("failed to write data.")
			os.Exit(1)
		}
	}
	log.Println("file write done.")
}

func TestFileReader(t *testing.T) {
	// 除了自己打开一个文件，还可以用os.Stdin作为读入对象
	file, err := os.Open("./proverbs")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p := make([]byte, 4) //缓冲[]byte
	for {
		n, err := file.Read(p) //将文件中的数据读入缓冲[]byte
		if err == io.EOF {     //读到end of file，退出
			break
		}
		log.Println(string(p)) //打印结果
	}
}

func TestIOCopy(t *testing.T) {
	//将reader读入的数据，写出到writer中
	// reader，可以直接打开一个已有文件作为读入源
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	// writer，可以不用打开一个文件，直接使用os.Stdout作为写出对象
	file, err := os.Create("./proverbs.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// copy from reader data into writer file
	if _, err := io.Copy(file, proverbs); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("file created")
}
