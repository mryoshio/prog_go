package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = fmt.Sprintf("%s%s", "http://", url)
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	f, err := os.OpenFile(fmt.Sprintf("file_%d.txt", time.Now().Unix()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		ch <- fmt.Sprintf("file open err: %v", err)
		return
	}
	nbytes, err := f.Write(b)
	f.Close()
	if err != nil {
		ch <- fmt.Sprintf("write file err: %v", err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
