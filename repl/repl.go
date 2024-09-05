package repl

import (
	"bufio"
	"fmt"
	"lizzardDB/parser"
	"lizzardDB/storage"
	"os"
	"strings"
)

func Start() { //Start the REPL loop
	store :=  storage.NewKeyValueStore()
	reader := bufio.NewReader((os.Stdin))

	fmt.Println("Welcome to LizzardDB!")
	fmt.Println("Type 'EXIT' to quit")

	//Infinite loop
	for {
		//Display the prompt
		fmt.Print("LizzardDB> ")

		//Read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		command, args := parser.ParseCommand(input)

		switch command {
		case "INSERT": 
			if (len(args) == 2) {
				store.Insert(args[0], args[1])
			} else {
				fmt.Println("Usage: INSERT <key> <value>")
			}
		case "SELECT":
			if (len(args) == 1) {
				store.Select(args[0])
			} else {
				fmt.Println("Usage: SELECT <key>")
			}
		case "DELETE": 
			if (len(args) == 1) {
				store.Delete(args[0])
			} else {
				fmt.Println("Usage: DELETE <key>")
			}
		case "EXIT": 
			fmt.Println("Thanks for using LizzardDB!")
			os.Exit(0)
		default: 
			fmt.Println("Unknown command. Available commands: INSERT, SELECT, DELETE and EXIT")
		}
	}
}