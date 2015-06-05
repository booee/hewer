package main;

import(
    "fmt"
    "strings"
)

type Analytics struct {
    key string
    total int
    sum int
    max int
    min int
}

func NewAnalytics(key string) *Analytics {
    a := new(Analytics);

    a.key = key;
    a.total = 0;
    a.sum = 0;
    a.max = 0;
    a.min = 0; // TODO: fix min so it doesn't default to 0

    return a;
}

func (a *Analytics) NewRow(data map[string]interface{}) {
    a.total += 1;

    watchedValue := nestedGet(a.key, data);

    switch casted := watchedValue.(type) {
        case int:
            a.numberEncountered(casted);
        case float64:
            a.numberEncountered(int(casted));
        default:
            // TODO: add properties to extraProperties Set
            fmt.Printf("%q\n", casted)
    }

    //TODO: mine out desired values
}

func (a *Analytics) numberEncountered(encountered int) {
    a.sum += encountered;

    if(a.max < encountered) {
        a.max = encountered;
    }

    if(a.min > encountered) {
        a.min = encountered;
    }

}

func nestedGet(key string, data interface{}) (value interface{}) {
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

    // value = data["processingMs"];
    return value;
}

func (a *Analytics) Print() {
    fmt.Println("Key: " + a.key);
    fmt.Printf("Total Rows: %d\n", a.total);
    fmt.Printf("Average: %d\n", (a.sum / a.total));
    fmt.Printf("Max: %d\n", a.max);
    fmt.Printf("Min: %d\n", a.min);
}
