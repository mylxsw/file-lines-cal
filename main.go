package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// 代码行数统计
// find .  -name "*.scala" -or -name "*.java" -or -name "*.go" -or -name "*.py" -or -name "*.js" -or -name "*.php" -or -name "*.html" -or -name "*.css" -or -name "*.less" -or -name "*.sass" -or -name "*.ts" | grep -v -e '/vendor/' -e '/node_modules/' | /Users/mylxsw/codes/examples/go-skills/lines/lines
func main() {

	files := readStdin(math.MaxInt64)

	fileCount := 0
	total := 0
	totalGroup := make(map[string]int)
	for _, f := range files {
		func() {
			filename := strings.TrimSuffix(f, "\n")
			file, err := os.Open(filename)
			if err != nil {
				panic(err)
			}

			defer file.Close()

			totalInFile, err := lineCounter(file)
			if err != nil {
				panic(err)
			}

			total += totalInFile
			fileCount++

			fileType := filepath.Ext(filename)
			totalGroup[fileType] += totalInFile
		}()
	}

	fmt.Printf("FileCount: %d, Total: %d\n", fileCount, total)

	extTotals := make(ExtTotals, 0)
	for typ, lines := range totalGroup {
		extTotals = append(extTotals, ExtTotal{
			Name:  typ,
			Total: lines,
		})

	}

	sort.Sort(sort.Reverse(extTotals))
	for _, et := range extTotals {
		fmt.Printf("🚩 %-8s: %d\n", et.Name, et.Total)
	}
}

type ExtTotal struct {
	Name  string
	Total int
}

type ExtTotals []ExtTotal

// Len is the number of elements in the collection.
func (e ExtTotals) Len() int {
	return len(e)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (e ExtTotals) Less(i int, j int) bool {
	return e[i].Total < e[j].Total
}

// Swap swaps the elements with indexes i and j.
func (e ExtTotals) Swap(i int, j int) {
	e[i], e[j] = e[j], e[i]
}

func lineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func readStdin(maxLines int) []string {
	result := make([]string, 0)

	reader := bufio.NewReader(os.Stdin)
	lineNo := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		lineNo++
		if lineNo > maxLines {
			break
		}

		result = append(result, line)
	}

	return result
}
