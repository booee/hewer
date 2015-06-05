package main

import (
	"bufio"
	"encoding/json"
	"fmt"
    "os"
)

type RawData struct {

}

func ParseFile(fileName string) {
    if fileName == nil {
        panic("Need a file, boss");
    }

	// fileName := "/Users/bradbowie/Desktop/metrics-calamp-blue-20501.log"; // TODO: via CLI
	harvestedData := harvestFile(fileName);
	data := analyzeData(harvestedData)
}

func harvestFile(fileName string) (abstracted) {
	fmt.Println("Mining JSON file: " + fileName);
	file, err := os.Open(fileName)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	var jsonData map[string][]json.RawMessage
    var abstracted RawData;

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

 		err := json.Unmarshal(scanner.Text(), &jsonData)

		fmt.Println(jsonData)

		// TODO: marshall

		// TODO: abstract value(s) based on key

		// return

	}
}

func analyzeData(data RawData) {
	// TODO: crunch and display statistics
}
