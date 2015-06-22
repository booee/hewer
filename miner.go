package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ParseFile(fileName string, analytics *Analytics) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	var data map[string]interface{}

	fmt.Println("Mining JSON file: " + fileName)
	skippedLines := 0
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		analytics.OnRow()
		err := json.Unmarshal(scanner.Bytes(), &data)

		if err != nil {
			skippedLines++
		} else {
			analytics.OnData(data)
		}
	}

	if skippedLines > 0 {
		fmt.Printf("Skipped %d non-JSON lines in '%s'.\n", skippedLines, fileName)
	}
}

func ParseAndConvertFile(fileName string, converter *Converter) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	var data map[string]interface{} //TODO: support other types via interface{}

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		err := json.Unmarshal(bytes, &data)

		if err == nil {
			converter.OnJsonData(data)
		} else {
			converter.OnOtherData(bytes)
		}
	}
}
