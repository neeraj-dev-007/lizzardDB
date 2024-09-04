package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Start the REPL loop
	reader := bufio.NewReader((os.Stdin))

	fmt.Println("Welcome to LizzardDB!")
	fmt.Println("Type 'EXIT' to quit")

	//Infinite loop 
	for{
		//Display the prompt
		fmt.Print("lizzardDB > ")

		//Read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		//Handle the 'EXIT' command
		if(input == "EXIT") {
			fmt.Println("Thanks for using LizzardDB!")
			break
		}

		//For now, just echo back the command
		fmt.Println("You entered: ", input)
	}
}