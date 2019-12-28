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
	arrLength := parseInt(readLine(reader))

	data := readLine(reader)

	elements := strings.Split(data, " ")

	if int64(len(elements)) != arrLength{
		panic("Bad Input")
	}
	var sum int64
	for _, element := range elements{
		sum = sum + parseInt(element)
	}
	fmt.Println(sum)

}

func readLine(reader *bufio.Reader)string{
	raw, _, err :=reader.ReadLine()
	if err == io.EOF{
		return ""
	}
	return string(strings.TrimRight(string(raw), "\r\n"))
}

func parseInt(str string) int64{
	item, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("Bad Input")
	}
	return item
}
