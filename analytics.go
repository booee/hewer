package main;

import(
    "fmt"
    "sort"
    "strings"
)

type Analytics struct {
    key string
    rows int
    encounters int
    sum int
    max int
    min int
    subkeys map[string]struct{} // map used as a set
}

func NewAnalytics(key string) *Analytics {
    a := new(Analytics);

    a.key = key;
    a.rows = 0;
    a.encounters = 0;
    a.sum = 0;
    a.max = 0;
    a.min = 9999999999; // TODO: find a better way to represent max/min

    return a;
}

func (a *Analytics) NewRow(data map[string]interface{}) {
    a.rows += 1;

    watchedValue := nestedGet(a.key, data);

    switch casted := watchedValue.(type) {
        case int:
            a.numberEncountered(casted);
        case float64:
            a.numberEncountered(int(casted));
        case map[string]interface{}:
            a.mapEncountered(casted);
        case []interface{}:
            a.sliceEncountered(casted)
    }
}

func (a *Analytics) numberEncountered(encountered int) {
    a.encounters += 1;
    a.sum += encountered;

    if(a.max < encountered) {
        a.max = encountered;
    }

    if(a.min > encountered) {
        a.min = encountered;
    }
}

func (a *Analytics) mapEncountered(encountered map[string]interface{}) {
    a.encounters += 1;
    fmt.Sprintf("%q", encountered);
    // fmt.Println("not a leaf node!")

    if len(a.subkeys) == 0 {
        // make map and populate
        a.subkeys = make(map[string]struct{}, len(encountered))
        for k := range encountered {
            a.subkeys[k] = struct{}{}
        }
    } else {
        // iterate map, append to subkeys as necessary
        for k := range encountered {
            // check if subkeys contains k
            if _, inSet := a.subkeys[k]; !inSet {
                // fmt.Println("Found new subkey:", k)
                a.subkeys[k] = struct{}{}
            }
        }
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

func nestedGet(key string, data interface{}) (value interface{}) {
    if key == "" {
        return data
    }

    keys := strings.Split(key, ".");

    value = data;

    for _, k := range keys {
        // fmt.Println(data);

        switch t := value.(type) {
            case map[string]interface{}:
                value = t[k];

            default:
                // value doesn't exist b/c parent value isn't a map
                panic(fmt.Sprintf("err mer gerrrd, %s doesn't exist in %s", k, key));
        }
    }

    return value;
}

func getSortedKeys(m map[string]struct{}) ([]string) {
    sortedKeys := make([]string, len(m))
    i := 0
    for k, _ := range m {
        sortedKeys[i] = k
        i++
    }
    sort.Strings(sortedKeys)
    return sortedKeys
}

func (a *Analytics) Print() {
    fmt.Println("Key: " + a.key);
    fmt.Printf("Total Rows: %d\n", a.rows);
    fmt.Printf("Total Encounters: %d\n", a.encounters);
    if a.encounters > 0 {
        fmt.Printf("Average: %d\n", (a.sum / a.encounters));
    }
    fmt.Printf("Max: %d\n", a.max);
    fmt.Printf("Min: %d\n", a.min);
    fmt.Printf("Subkeys: %v\n", getSortedKeys(a.subkeys))
}
