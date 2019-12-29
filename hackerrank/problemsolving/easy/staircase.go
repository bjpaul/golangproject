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
	size := int(parseInt(readLine(reader)))
	buildStairCase(size)
}

func buildStairCase(size int){
	for i := 1; i <= size; i++{
		fmt.Println(fmt.Sprintf("%*v",size, strings.Repeat("#",i)))
	}
}

func readLine(reader *bufio.Reader)string{
	str,_, err := reader.ReadLine()

	if err == io.EOF{
		return ""
	}

	return strings.TrimRight(string(str), "/r/n")
}

func parseInt(str string)int64{
	item, err := strconv.ParseInt(str, 10, 64)
	if err != nil{
		panic("Bad Input")
	}
	return item
}
