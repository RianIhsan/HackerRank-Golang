package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	s = strings.ReplaceAll(s, " ", "")

	L := len(s)
	rows := int(math.Floor(math.Sqrt(float64(L))))
	cols := int(math.Ceil(math.Sqrt(float64(L))))

	if rows*cols < L {
		rows++
	}

	var result strings.Builder
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			index := c + r*cols
			if index < L {
				result.WriteByte(s[index])
			}
		}
		result.WriteByte(' ')
	}

	return result.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
