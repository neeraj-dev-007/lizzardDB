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
	db :=  storage.NewDatabase()
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
		case "CREATE": 
			if (len(args)>=2 && args[0] == "TABLE") {
				name := args[1]
				columns := []storage.Column{
					{Name: "id", Type: "INT"},
					{Name: "name", Type: "VARCHAR"},
				}
				db.Create(name, columns)
			} else {
				fmt.Println("Usage: CREATE TABLE <tableName>")
			}
		case "INSERT": 
			if (len(args)>=2 && args[0] == "INTO") {
				name := args[1]
				row := map[string]interface{}{
					"id": 1,
					"name": "Harvey Specter",
				}
				db.Insert(name, row)
			} else {
				fmt.Println("Usage: INSERT INTO <tableName>")
			}
		case "SELECT":
			if (len(args)>=2 && args[0] == "*") {
				name := args[2]
				db.Select(name)
			} else {
				fmt.Println("Usage: SELECT * FROM <tableName>")
			}
		case "EXIT": 
			fmt.Println("Thanks for using LizzardDB!")
			os.Exit(0)
		default: 
			fmt.Println("Unknown command. Available commands: CREATE, INSERT, SELECT and EXIT")
		}
	}
}