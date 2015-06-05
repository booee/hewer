package main


func main() {
	a := NewAnalytics("system.activeThreads");
	ParseFile("/Users/bradbowie/Desktop/metrics-calamp-blue-20501.log", a);
	a.Print();
}
