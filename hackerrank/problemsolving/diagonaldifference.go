package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	matrixLength := parseInt(readLine(reader))
	arr := readMatrix(int(matrixLength), reader)
	fmt.Println(findAbsDiagonalDiff(arr))
}

func findAbsDiagonalDiff(arr [][]int64)int64{
	length := len(arr)
	var diagonal1, diagonal2 int64
	start, end := 0,length-1

	for i := 0; i < length; i ++{
		diagonal1 += arr[i][i]
		diagonal2 += arr[start+i][end -i]
	}

	return int64(math.Abs(float64(diagonal1 - diagonal2)))
}

func readMatrix(matrixLength int, reader *bufio.Reader) [][]int64{
	arr := make([][]int64, matrixLength)
	for i :=0; i < matrixLength; i++{
		arr[i] = make([]int64, matrixLength)
		data := strings.Split(readLine(reader), " ")
		if len(data) != matrixLength{
			panic("Bad Input")
		}

		for j, v := range data{
			arr[i][j] = parseInt(v)
		}
	}
	return arr
}

func readLine(reader *bufio.Reader)string  {
	raw,_,err := reader.ReadLine()
	if err == io.EOF{
		return ""
	}
	return strings.TrimRight(string(raw),"/r/n")
}

func parseInt(str string)int64{
	item,err:= strconv.ParseInt(str, 10, 64)
	if err != nil{
		panic("Bad Input")
	}
	return item
}