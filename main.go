package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func main() {
	//first file
	DownloadFile("https://thumbs.dreamstime.com/z/die-fantacy-immage-woman-s-hand-draws-31994627.jpg", "./")
	//second file
	DownloadFile("https://wordpress.org/latest.tar.gz", "./")
}

//Print: file name, percent of copying file, time
func calculateAndPrintProgress(done chan int64, destination string, total int64) {
	//true when downloading done
	var stop bool = false

	start := time.Now()

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(destination)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100

			printProgress(path.Base(destination), percent, total)
		}

		if stop {
			//if file small we should display 100 persent at the end
			printProgress(path.Base(destination), 100, total)

			elapsed := time.Since(start)
			fmt.Printf("\nDownload completed in %s\n", elapsed)

			break
		}

		time.Sleep(time.Millisecond * 100)
	}
}

//download file from url to local file dest
func DownloadFile(url string, dest string) {

	file := path.Base(url)

	fmt.Printf("Downloading:\n")

	path := dest + "/" + file

	out, err := os.Create(path)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}

	defer out.Close()

	headResp, err := http.Head(url)
	if err != nil {
		panic(err)
	}

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		panic(err)
	}

	done := make(chan int64)

	go calculateAndPrintProgress(done, path, int64(size))

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	//copy done return file size to chanel
	done <- n
}

func printProgress(fileName string, percent float64, total int64) {
	fmt.Printf("\r"+fileName+"(%s)       %.0f%%", humanize.Bytes(uint64(total)), percent)
}
