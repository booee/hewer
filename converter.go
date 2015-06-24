package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Converter struct {
	key    string
	toType string
}

func NewConverter(key string, toType string) *Converter {
	c := new(Converter)

	c.key = key
	c.toType = toType

	return c
}

func (c *Converter) OnOtherData(bytes []byte) {
	os.Stdout.Write(bytes)
	fmt.Println()
}

func (c *Converter) OnJsonData(data map[string]interface{}) {
	var owner map[string]interface{}
	owner = nil
	prop := c.key
	var value interface{}
	value = data

	if c.key != "" {
		keys := strings.Split(c.key, ".")
		for _, k := range keys {
			switch t := value.(type) {
			case map[string]interface{}:
				owner = t
				prop = k
				value = t[k]
			}
		}
	}

	if value != nil {
		switch casted := value.(type) {
		case string:
			value = convertString(casted, c.toType)
		case int: // don't think this will ever happen
			value = convertInt(casted, c.toType)
		case float64:
			value = convertFloat(casted, c.toType)
			// case map[string]interface{}:
			// 	value = convertMap(casted, c.toType)
			// case []interface{}:
			// 	value = convertSlice(casted, c.toType)
		}
		if owner != nil {
			owner[prop] = value
		} else {
			data = value.(map[string]interface{})
		}
	}

	b, err := json.Marshal(data)
	if err != nil {
		panic(err) //TODO: this is not the Go way
	}
	os.Stdout.Write(b)
	fmt.Println()
}

func convertString(s string, toType string) (converted interface{}) {
	converted = s
	switch toType {
	case "number":
		if strings.Contains(s, ".") {
			converted, _ = strconv.ParseFloat(s, 64)
		} else {
			converted, _ = strconv.Atoi(s)
		}
	case "json":
		json.Unmarshal([]byte(s), &converted)
	}
	return converted
}

func convertInt(i int, toType string) (converted interface{}) {
	converted = i
	switch toType {
	case "string":
		converted = strconv.Itoa(i)
	}
	return converted
}

func convertFloat(f float64, toType string) (converted interface{}) {
	converted = f
	switch toType {
	case "string":
		converted = strconv.FormatFloat(f, 'f', -1, 64)
	}
	return converted
}
