package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Analytics struct {
	rows       int
	jsonRows   int
	key        string
	encounters int

	// numerical analytics
	sum int
	max int
	min int

	// non-leaf node analytics
	subkeys   map[string]struct{} // map used as a set
	datatypes map[string]struct{} // map used as a set
}

func NewAnalytics(key string) *Analytics {
	a := new(Analytics)

	a.key = key
	a.rows = 0
	a.encounters = 0

	a.sum = 0
	a.max = -1        // TODO: find a better way to represent max
	a.min = 1<<63 - 1 // TODO: find a better way to represent min

	return a
}

func (a *Analytics) OnRow() {
	a.rows++
}

func (a *Analytics) OnData(data map[string]interface{}) {
	a.jsonRows++
	watchedValue := nestedGet(a.key, data)

	if watchedValue != nil {
		a.encounters += 1

		if len(a.datatypes) == 0 {
			a.datatypes = make(map[string]struct{}, 1)
		}
		addStringToSet(fmt.Sprintf("%v", reflect.TypeOf(watchedValue)), a.datatypes)

		switch casted := watchedValue.(type) {
		case int:
			a.numberEncountered(casted)
		case float64:
			a.numberEncountered(int(casted))
		case map[string]interface{}:
			a.mapEncountered(casted)
		case []interface{}:
			a.sliceEncountered(casted)
		}
	}
}

func nestedGet(key string, data interface{}) (value interface{}) {
	if key == "" {
		return data
	}

	keys := strings.Split(key, ".")

	value = data

	for _, k := range keys {
		// fmt.Println(data);

		switch t := value.(type) {
		case map[string]interface{}:
			value = t[k]
			// ignore any other types/cases
		}
	}

	return value
}

func addStringToSet(str string, mapAsSet map[string]struct{}) {
	if _, inSet := mapAsSet[str]; !inSet {
		mapAsSet[str] = struct{}{}
	}
}

func (a *Analytics) numberEncountered(encountered int) {
	a.sum += encountered

	if a.max < encountered {
		a.max = encountered
	}

	if a.min > encountered {
		a.min = encountered
	}
}

func (a *Analytics) mapEncountered(encountered map[string]interface{}) {
	if len(a.subkeys) == 0 {
		a.subkeys = make(map[string]struct{}, len(encountered))
	}

	for k := range encountered {
		addStringToSet(k, a.subkeys)
	}
}

func (a *Analytics) sliceEncountered(encountered []interface{}) {
	//TODO if slice contains strings, then optionally group by value
	//     e.g separate enc/sum/min/max per value
	//     would require a struct as subset of Analytics
	//     could either group string values by default or require a new cli flag
	//     also, what about json objects with properties in slice?
	a.numberEncountered(len(encountered))
}

func (a *Analytics) Print() {
	fmt.Printf("Total Rows: %d\n", a.rows)

	if a.jsonRows > 0 {
		printDataAnalytics(a)
	}
}

func printDataAnalytics(a *Analytics) {
	fmt.Printf("JSON Rows: %d\n", a.rows)

	printDatatypes := true

	if a.key != "" {
		fmt.Println("Key: " + a.key)
		fmt.Printf("Total Encounters: %d\n", a.encounters)

		if a.encounters > 0 && a.max != -1 {
			printDatatypes = false
			fmt.Printf("Average: %d\n", (a.sum / a.encounters))
			fmt.Printf("Max: %d\n", a.max)
			fmt.Printf("Min: %d\n", a.min)
		}
	}

	if len(a.subkeys) > 0 {
		printDatatypes = false
		fmt.Printf("Subkeys: %v\n", getSortedKeys(a.subkeys))
	}

	if printDatatypes && len(a.datatypes) > 0 {
		fmt.Printf("Types: %v\n", getSortedKeys(a.datatypes))
	}
}

func getSortedKeys(m map[string]struct{}) []string {
	sortedKeys := make([]string, len(m))
	i := 0
	for k, _ := range m {
		sortedKeys[i] = k
		i++
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}
