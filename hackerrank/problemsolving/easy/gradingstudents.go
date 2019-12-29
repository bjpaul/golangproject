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
	studentCount := parseInt(readLine(reader))

	studentMarks := make([]int, studentCount)

	for i:=0; i< studentCount; i++{
		studentMarks[i] = parseInt(readLine(reader))
		if studentMarks[i] >= 38{
			nearestMarksofMultipleFive := ((studentMarks[i] / 5)+1)*5
			if nearestMarksofMultipleFive - studentMarks[i] < 3{
				studentMarks[i] = nearestMarksofMultipleFive
			}
		}
	}

	for _, v := range studentMarks{
		fmt.Println(v)
	}
}

func readLine(reader *bufio.Reader)string{
	str,_, err := reader.ReadLine()

	if err == io.EOF{
		return ""
	}

	return strings.TrimRight(string(str), "/r/n")
}

func parseInt(str string) int{
	item, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("Bad Input")
	}
	return int(item)
}