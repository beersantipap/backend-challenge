package main

import (
	// "os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)


func read_json_file(file_path string) [][]int{
	file_json, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	var input [][]int

	err = json.Unmarshal(file_json, &input)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	return input
}


func sumMaxValue(input [][]int) int {
	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if j == 0 {
				input[i][j] += input[i-1][j]
			} else if j == len(input[i])-1 {
				input[i][j] += input[i-1][j-1]
			} else {
				if input[i-1][j-1] > input[i-1][j] {
					input[i][j] += input[i-1][j-1]
				} else {
					input[i][j] += input[i-1][j]
				}
			}
		}
	}

	output := 0
	output_arr := input[len(input) - 1]
	for i := 0; i < len(output_arr); i ++ {
		if output_arr[i] > output {
			output = output_arr[i]
		}
	}

	return output
}


func main() {
	// input := [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }
	input := read_json_file("files/hard.json")

	fmt.Println(sumMaxValue(input))
}
