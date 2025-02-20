package main

import (
	"bufio"
	"fmt"

	"os"
	"sort"
)

func main() {
	var dicts = map[string]string{"Test1": "this is test1 ", "Test2": "this is test1 ", "Atest": "test3"}
	reader := bufio.NewReader(os.Stdin)
	var exit = false
	for !exit {
		fmt.Println("Select Options :")
		if dicts == nil || len(dicts) == 0 {

			fmt.Println("1 -Add ")
		} else {

			fmt.Println("1 -Add ")
			fmt.Println("2 -Remove ")
			fmt.Println("3 -List")
			fmt.Println("5 -Exit")
		}

		input, _ := reader.ReadString('\n')
		switch input {
		case "1\n":
			fmt.Println("starting the add function	")
		case "2\n":
			fmt.Println("starting the remove function	")
		case "3\n":
			fmt.Println("starting the list function")
			List(dicts)
		case "5\n":
			fmt.Println("exiting")
			exit = true
		default:
			fmt.Println("invalid command")

		}
	}

}
func List(dicts map[string]string) {

	if dicts == nil || len(dicts) == 0 {
		fmt.Print("The dictionary is empty." +
			"")
	}

	var keys []string
	for ks := range dicts {
		keys = append(keys, ks)
	}
	sort.Strings(keys)
	fmt.Println("Printing Dictonaary")
	for _, element := range keys {
		fmt.Println("%s -> %s\n", element, dicts[element])
	}

}
