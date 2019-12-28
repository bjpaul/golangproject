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
	a := setRating(readLine(reader))
	b := setRating(readLine(reader))
	fmt.Println(comPareRating(a, b))
}

func comPareRating(a []int64, b []int64)(int, int){
	length := len(a)
	var aRating, bRating int

	for i := 0; i < length; i++ {
		if a[i] > b[i]{
			aRating++
		}else if a[i] < b[i]{
			bRating++
		}
	}
	return aRating, bRating
}

func setRating(str string)[]int64{
	arr := strings.Split(str, " ")
	if len(arr) != 3 {
		panic("Bad Input")
	}
	var ratings []int64
	for _, s := range arr{
		element, err := strconv.ParseInt(s, 10, 32)
		checkError(err)
		if element < 1 || element > 100{
			panic("Bad Input")
		}
		ratings = append(ratings, element)
	}
	return ratings
}

func readLine(reader *bufio.Reader)string{
	str,_,err := reader.ReadLine()
	if err == io.EOF{
		return ""
	}
	return strings.TrimRight(string(str), "/r/n")
}

func checkError(err error){
	if err != nil {
		panic("Bad Input")
	}
}
