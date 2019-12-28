package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
	//
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i2, _ := strconv.Atoi(scanner.Text())

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	scanner.Scan()
	f2, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	s2 := scanner.Text()

	fmt.Println(int(i) + i2 )
	fmt.Println(fmt.Sprintf("%.1f", d + f2))
	fmt.Println(s + s2)
}
