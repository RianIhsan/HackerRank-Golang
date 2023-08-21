package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func biggerIsGreater(w string) string {
	r := []rune(w)

	i := len(r) - 2
	for i >= 0 && r[i] >= r[i+1] {
		i--
	}

	if i == -1 {
		return "no answer"
	}

	j := len(r) - 1
	for r[j] <= r[i] {
		j--
	}

	r[i], r[j] = r[j], r[i]

	left, right := i+1, len(r)-1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}

	return string(r)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
