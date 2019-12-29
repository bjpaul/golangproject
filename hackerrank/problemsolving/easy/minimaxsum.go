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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	data := strings.Split(readLine(reader), " ")

	length := len(data)
	if length != 5 {
		panic("Bad Input")
	}
	arr := make([]int64, length)

	for i, v := range data{
		arr[i] = parsePositiveInt(v)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var max, min int64

	for i, j:=0,length - 1 ; i< length - 1 && j>0; {
		min = min + arr[i]
		max = max + arr[j]
		i++
		j--
	}
	fmt.Println(min,max)
}

func readLine(reader *bufio.Reader)string{
	str,_, err := reader.ReadLine()

	if err == io.EOF{
		return ""
	}

	return strings.TrimRight(string(str), "/r/n")
}

func parsePositiveInt(str string)int64{
	item, err := strconv.ParseInt(str, 10, 64)
	if err != nil{
		panic("Bad Input")
	}
	if item < 0{
		panic("Bad Input")
	}
	return item
}

