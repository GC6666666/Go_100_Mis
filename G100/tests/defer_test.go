package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "C:\\ZPMC\\DB\\mytest.txt"
	close(ch)
	if err := readFiles(ch); err != nil {
		fmt.Println("error:", err)
	}
}

func readFiles(ch <-chan string) error {
	for path := range ch {
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			// 对文件进行操作
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// 对文件进行操作
	return nil
}
