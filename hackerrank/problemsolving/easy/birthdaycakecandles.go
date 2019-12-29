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
	count := int(parseInt(readLine(reader)))
	data := strings.Split(readLine(reader), " ")
	length := len(data)
	if length != count{
		panic("Bad Input")
	}

	max := parseInt(data[0])
	maxCandleCount := 1
	for i:=1; i < length;i++{
		if max == parseInt(data[i]){
			maxCandleCount ++
		}else if max < parseInt(data[i]){
			max = parseInt(data[i])
		}
	}
	fmt.Println(maxCandleCount)
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