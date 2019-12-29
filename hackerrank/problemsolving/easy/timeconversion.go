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
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	time:=convertToTwentyfourHourFormat(readLine(reader))
	fmt.Println(time)
}

func convertToTwentyfourHourFormat(time string) string{
	twelveHourFormat := []rune(time)
	dayOrNight := string(twelveHourFormat[8:10])
	var twentyfourHourFormat string
	var hour = parseInt(string(twelveHourFormat[:2]))
	if dayOrNight == "AM"{
		if hour == 12{
			twentyfourHourFormat = fmt.Sprintf("%s:%s", "00", string(twelveHourFormat[3:8]))
		}else if hour < 12{
			twentyfourHourFormat = string(twelveHourFormat[:8])
		}else {
			panic("Bad Input")
		}
	}else if dayOrNight == "PM"{
		if hour == 12{
			twentyfourHourFormat = string(twelveHourFormat[:8])
		}else if hour < 12{
			hour = hour + 12
			twentyfourHourFormat = fmt.Sprintf("%d:%s", hour, string(twelveHourFormat[3:8]))
		}else {
			panic("Bad Input")
		}
	}else {
		panic("Bad Input")
	}
	return twentyfourHourFormat
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
