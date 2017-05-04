package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = fmt.Sprintf("%s%s", "http://", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "StatusCode: %d\n", resp.StatusCode)
		if _, err = io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "copy error: %v\n", err)
		}
		resp.Body.Close()
	}
}
