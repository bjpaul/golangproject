package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var T int

	fmt.Scan(&T)
	if T > 0 {
		strArray := make([]string, T)
		reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

		for i := 0; i < T; i++ {
			strArray[i] = readLine(reader)
		}
		for  _, v := range strArray  {
			printOddChars(v)
			fmt.Print(" ")
			printEvenChars(v)
			fmt.Println()
		}
	}
}

func printOddChars(str string){
	for i := 0; i < len(str); i = i + 2   {
		fmt.Print(string(str[i]))
	}
}

func printEvenChars(str string){
	for i := 1; i < len(str); i = i + 2   {
		fmt.Print(string(str[i]))
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
