package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMinCost' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY crew_id
 *  2. INTEGER_ARRAY job_id
 */

func getMinCost(crew_id []int32, job_id []int32) int64 {
	var minCost int64
	sort.Slice(crew_id, func(i, j int) bool {
		return crew_id[i] < crew_id[j]
	})
	sort.Slice(job_id, func(i, j int) bool {
		return job_id[i] < job_id[j]
	})
	for i := 0; i < len(crew_id); i++ {
		minCost += int64(abs(crew_id[i] - job_id[i]))
	}
	return minCost
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	crew_idCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var crew_id []int32

	for i := 0; i < int(crew_idCount); i++ {
		crew_idItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		crew_idItem := int32(crew_idItemTemp)
		crew_id = append(crew_id, crew_idItem)
	}

	job_idCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var job_id []int32

	for i := 0; i < int(job_idCount); i++ {
		job_idItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		job_idItem := int32(job_idItemTemp)
		job_id = append(job_id, job_idItem)
	}

	result := getMinCost(crew_id, job_id)

	fmt.Fprintf(writer, "%d\n", result)

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
