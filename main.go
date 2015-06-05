package main


func main() {
	a := NewAnalytics("key.from.cli");
	ParseFile("/Users/bradbowie/Desktop/metrics-calamp-blue-20501.log", a);
	a.Print();
}
