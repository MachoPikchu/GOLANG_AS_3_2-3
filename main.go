package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func downloadFile(url, filePath string, wg *sync.WaitGroup, ch chan<- error) {
	defer wg.Done()

	file, err := os.Create(filePath)
	if err != nil {
		ch <- fmt.Errorf("error creating file %s: %v", filePath, err)
		return
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Errorf("error downloading file from %s: %v", url, err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		ch <- fmt.Errorf("error writing file %s: %v", filePath, err)
		return
	}

	ch <- nil
}

func main() {
	urls := []string{
		"https://example.com/file1.txt",
		"https://example.com/file2.txt",
		"https://example.com/file3.txt",
	}

	errorChannel := make(chan error)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go downloadFile(url, "downloaded_files/"+url[strings.LastIndex(url, "/")+1:], &wg, errorChannel)
	}

	go func() {
		wg.Wait()
		close(errorChannel)
	}()

	for err := range errorChannel {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}
