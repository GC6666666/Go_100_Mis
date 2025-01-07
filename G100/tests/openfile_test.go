package tests

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestOpenFile(t *testing.T) {
	//emptyline, _ := countEmptyLines(strings.NewReader(
	//	`foo
	//
	//
	//            baz
	//
	//
	//            baa`))
	//fmt.Println(emptyline)
	input := "Line 1\n\nLine 3\n  \nLine 5\n"
	reader := strings.NewReader(input)

	count, err := countEmptyLines(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number of empty lines: %d\n", count)
}

//func countEmptyLinesInFile(filename string) (int, error) {
//	file, err := os.Open(filename)
//	if err != nil {
//		return 0, err
//	}
//	// 处理文件闭包
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		// 迭代每行
//	}
//}

func countEmptyLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	emptyLines := 0

	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			emptyLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return emptyLines, nil
}
