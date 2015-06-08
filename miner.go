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

    fmt.Println("Mining JSON file: " + fileName);
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        err := json.Unmarshal(scanner.Bytes(), &data)

        if err != nil {
            panic(err);
        }

        analytics.NewRow(data);
	}
}
