package main

import "fmt"

func main() {
	map1 := map[string]string{}
	map1["one"] = "tom"
	map1["two"] = "tom1"
	map1["three"] = "tom2"
	delete(map1, "one")
	fmt.Println(map1)
}
