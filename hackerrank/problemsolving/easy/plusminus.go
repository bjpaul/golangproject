package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main()  {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	arrLength := int(parseInt(readLine(reader)))
	data := strings.Split(readLine(reader), " ")
	if arrLength != len(data){
		panic("Bad Input")
	}
	arr := make([]int64, arrLength)

	for i, v := range data {
		arr[i] = parseInt(v)
	}

	findFractions(arr)
}

func findFractions(data []int64){
	var pos, neg, zero int
	for _, item := range data{
		switch {
		case item > 0:
			pos ++
		case item < 0:
			neg ++
		default:
			zero ++
		}
	}
	length := len(data)
	fmt.Printf("%.6f\n",float64(pos)/float64(length))
	fmt.Printf("%.6f\n",float64(neg)/float64(length))
	fmt.Printf("%.6f\n",float64(zero)/float64(length))
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
