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

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			if arrItem < -9 || arrItem > 9{
				panic("Bad input")
			}
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	fmt.Println(sumOfHourGlass(arr))
}

func sumOfHourGlass(arr [][]int32)int32{
	var max int32
	for i := 1; i <= 4; i++{
		for j := 1; j <= 4; j++{
			//fmt.Println(arr[i - 1][j - 1],  arr[i - 1][j], arr[i - 1][j + 1])
			row1:= arr[i - 1][j - 1] + arr[i - 1][j] + arr[i - 1][j + 1]
			//fmt.Println(" ",arr[i][j])
			row2:=  arr[i][j]
			//fmt.Println(arr[i + 1][j - 1], arr[i + 1][j], arr[i + 1][j + 1])
			row3:=  arr[i + 1][j - 1] + arr[i + 1][j] + arr[i + 1][j + 1]

			sum := row1 + row2 + row3
			if (i == 1 && j == 1) || sum > max{
				max = sum
			}
		}
	}
	return max
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
