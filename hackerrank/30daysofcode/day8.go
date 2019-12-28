package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func populateMap(n int, reader *bufio.Reader) map[string]int64{

	dict := make(map[string]int64, n)

	for i := 0; i < n ;i ++  {
		str := readLine(reader)
		arr := strings.Split(str, " ")
		name := strings.TrimRight(arr[0], "\r\n")
		phone, err := strconv.ParseInt(strings.TrimRight(arr[1], "\r\n"),10, 64)
		checkError(err)
		dict[name] = phone
	}
	return dict
}

func checkIfKeyExist(s string,  m map[string]int64) {
	v, isExist := m[s]
	if isExist {
		fmt.Printf("%s=%d\n", s, v)
	} else {
		fmt.Println("Not found")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024* 1024)
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	dict :=populateMap(n, reader)

	var arr []string

	cont := true
	for cont  {
		s := strings.TrimRight(readLine(reader),"\r\n")
		if s == ""{
			cont = false
		}else {
			arr = append(arr, s)
		}
	}

	for i := 0; i < len(arr) ;i ++ {
		checkIfKeyExist(arr[i], dict)
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
