package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	urls := []string{
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_2mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_5mb.mp4",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			err := downloadFile(url, "downloads/")
			if err != nil {
				fmt.Printf("Error downloading file from %s: %v\n", url, err)
			} else {
				fmt.Printf("Downloaded file from %s\n", url)
			}
		}(url)
	}

	wg.Wait()
	fmt.Println("All files downloaded.")
}

func downloadFile(url string, destination string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	filePath := destination + getFileName(url)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func getFileName(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1]
}
