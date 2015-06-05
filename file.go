package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// run via 'go build main.go && ./main'

func main() {
	fileName := "/test/file/name"

	fmt.Println(fileName)

    check(errors.New("zomg"));
}
