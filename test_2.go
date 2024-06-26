package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


func convertText(input string) string {
	output := ""
	for i := 1; i < len(input); i++ {
		num1, _ := strconv.Atoi(string(input[i-1]))
		num2, _ := strconv.Atoi(string(input[i]))
		if num1 == num2 {
			output += "="
		} else if num1 > num2 {
			output += "L"
		} else {
			output += "R"
		}
	}
	return output
}


func main() {
	fmt.Println("input value: ")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)
		if _, err := strconv.Atoi(input); err == nil {
			fmt.Println("Output: ", convertText(input))
		}
	}
}
