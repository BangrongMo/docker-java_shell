package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
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
	/* srcFile, _ := os.Open("bigfile")
	defer srcFile.Close()
	os.Mkdir("tmpdir", 0744)
	//os.Chmod("tmpdir", 0744)
	destFile, _ := os.Create("tmpdir/bigfile755")
	defer destFile.Close()
	mybuffer := make([]byte, 65535)
	bytesCopy, _ := io.CopyBuffer(destFile, srcFile, mybuffer)
	os.Chmod("bigfile755", 0755)
	fmt.Printf("bytes copied = %v\n", bytesCopy)*/
	url := "https://static-cdn.wotgame.cn/static/6.3.2_ee8aab/wotp_static/img/core/frontend/scss/common/blocks/video-bg/img/video-bg_cn.mp4"
	/*resp, _ := http.Get(url)
	defer resp.Body.Close()
	dst, _ := os.Create("output.mp4")
	defer dst.Close()
	io.Copy(dst, resp.Body) // this code use 2M memory ,video is 10M  */

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	dst, _ := os.Create("output.mp4")
	defer dst.Close()
	writer := bufio.NewWriterSize(dst, 65535)
	io.Copy(writer, resp.Body)
	time.Sleep(time.Second * 15)
	writer.Flush()

	// this code use 2M memory ,video is 10M

	wg.Wait()

}
