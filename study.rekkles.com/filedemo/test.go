package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 111
// 100 Red
// 010 Yellow
// 001 Green
const (
	Red    int = 4
	Yellow int = 2
	Green  int = 1
)

func color(arg int) {
	fmt.Printf("%b\n", arg)
}

func writedemo1() {
	fileObj, err := os.OpenFile("test1.txt", os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()
	fileObj.Write([]byte("zhoulin mengbi le\n"))
	fileObj.WriteString("hahahah")
}

func writedemo2() {
	fileObj, err := os.OpenFile("test1.txt", os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("qq")
	wr.Flush()
}

func writedemo3() {
	str := "gogog"
	err := ioutil.WriteFile("test1.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed, err:%v", err)
		return
	}
}

func Log() {
	fmt.Fprintln(os.Stdout, "this is log")
}

func fileInsert() {
	fileObj, err := os.OpenFile("test1.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()

	// tmpfile use to insert file
	tmpFile, err := os.OpenFile("sb.tmp", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	// write to tmp fie
	tmpFile.Write(ret[:n])
	fmt.Println(string(ret[:n]))
	// insert new content
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)

	// left origin file to tmpfile
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read err%v", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	fileObj.Close()
	tmpFile.Close()
	os.Rename("sb.tmp", "test1.txt")
	// fileObj.Seek(1, 0)
	// fileObj.Write([]byte{'c'})

}

func main() {
	// color(Red | Green)
	// open file

	// writedemo3()
	// Log()/
	fileInsert()

}
