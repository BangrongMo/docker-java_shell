package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func MonitorMem() {
	var ms runtime.MemStats
	for i := 0; i < 50; i++ {
		runtime.ReadMemStats(&ms)
		fmt.Printf("Alloc = %v KiB loop time %v \n", ms.Alloc/1024, i)
		time.Sleep(time.Millisecond * 20)
	}
	wg.Done()

}
func main() {
	wg.Add(1)
	go MonitorMem()
	time.Sleep(time.Millisecond * 10)
	//bytes, err := os.ReadFile("bigfile")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//os.WriteFile("bigfile2", bytes, 0666)
	srcFile, _ := os.Open("bigfile")
	defer srcFile.Close()
	destFile, _ := os.Create("bigfile755")
	defer destFile.Close()
	mybuffer := make([]byte, 65535)
	bytesCopy, _ := io.CopyBuffer(destFile, srcFile, mybuffer)
	os.Chmod("bigfile755", 0755)
	fmt.Printf("bytes copied = %v\n", bytesCopy)
	wg.Wait()

}
