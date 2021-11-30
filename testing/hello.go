package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var path = "test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	fmt.Println("Hello World!")
	var file, err = os.Open(path)
	if isError(err) {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	fmt.Println(txtlines)
	//var a int = 0
	a, err := strconv.Atoi(txtlines[0])
	fmt.Println(a, err)

	defer file.Close()
}
