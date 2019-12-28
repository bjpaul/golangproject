package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	firstNumber, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	secondNumber, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	fmt.Println(firstNumber, secondNumber)
	swap(&firstNumber, &secondNumber)
	fmt.Println(firstNumber, secondNumber)
}

func swap(a *int64, b *int64){
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
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