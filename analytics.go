package main;

import(
    "fmt"
)

type Analytics struct {
    key string
    total int
}

func NewAnalytics(key string) *Analytics {
    a := new(Analytics);

    a.key = key;
    a.total = 0;

    fmt.Println(key);

    return a;
}

func (a *Analytics) NewRow(data map[string]interface{}) {
    a.total += 1;

    //TODO: mine out desired values
}

func (a *Analytics) Display() {
    fmt.Println("Key: " + a.key);
    fmt.Println(fmt.Sprintf("Total Rows: %d", a.total));
}
