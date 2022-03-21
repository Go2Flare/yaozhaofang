package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//可执行文件中获取当前文件的路径目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//调试中可以获取当前的目录，但是可执行文件中可能不行
	runDir,_ := os.Getwd()
	fmt.Println("filepath.Abs :",dir)
	fmt.Println("os.Getwd :",runDir)

	time.Sleep(10*time.Second)
}
