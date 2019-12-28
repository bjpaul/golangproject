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
	var newArr [5]int64

	for i := 0; i < len(newArr); i ++{
		temp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		newArr[i] = temp
	}

	for i := 0; i < len(newArr); i ++{
		fmt.Println(newArr[i])
	}

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