package main

import (
	"bufio"
	"encoding/json"
	"fmt"
    "os"
)

func ParseFile(fileName string, analytics *Analytics) {
    fmt.Println("Mining JSON file: " + fileName);
    file, err := os.Open(fileName)

    if err != nil {
        panic(err.Error())
    }

    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    var data map[string]interface{}

    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        err := json.Unmarshal(scanner.Bytes(), &data)

        if err != nil {
            panic(err);
        }

        analytics.NewRow(data);
	}
}
