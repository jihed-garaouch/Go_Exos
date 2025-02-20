package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var dicts = map[string]string{"Test1": "this is test1 ", "Test2": "this is test1 "}
	reader := bufio.NewReader(os.Stdin)
	var exit = false
	for !exit {
		if dicts == nil || len(dicts) == 0 {
			/// TODO if the dicts is null must only show func Add
			fmt.Print("Enter your dictionary: ")
		}
		fmt.Println("Select Options :")
		fmt.Println("1 -Add ")
		fmt.Println("2 -Remove ")
		fmt.Println("3 -List")
		fmt.Println("5 -Exit")

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
	fmt.Println("Printing Dictonaary")
	for k, v := range dicts {

		fmt.Printf("%s -> %s\n", k, v)
	}

}
