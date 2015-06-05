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
    harvestFile(fileName);
    // data := analyzeData(harvestedData)
}

func harvestFile(fileName string) {
    fmt.Println("Mining JSON file: " + fileName);
    file, err := os.Open(fileName)

    if err != nil {
        panic(err.Error())
    }

    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    var jsonData map[string][]json.RawMessage
    // var abstracted RawData := nil;

    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {

        // fmt.Println(scanner.Text());

        err := json.Unmarshal(scanner.Bytes(), &jsonData)

        if err != nil {
            panic(err);
        }

        fmt.Println(jsonData)

        // TODO: marshall

        // TODO: abstract value(s) based on key

        // return

	}
}

func analyzeData(data *RawData) {
	// TODO: crunch and display statistics
}
