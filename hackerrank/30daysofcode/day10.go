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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	max, _ := findMaxOnes(n)
	fmt.Println(max)
}

func findMaxOnes(n int)(int, int){
	if n > 0 {
		remainder := n % 2
		n = n / 2
		max, i := findMaxOnes(n)
		if remainder == 1{
			i++
			if i > max{
				max = i
			}
		}else {
			i = 0
		}
		//fmt.Print(remainder)
		return max, i
	}else{
		return 0, 0
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