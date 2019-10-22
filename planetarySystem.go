package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	personName := bufio.NewReader(os.Stdin)

	fmt.Printf("Welcome to the Solar System!\n")
	fmt.Printf("There are 9 planets to explore.\n")
	fmt.Printf("What is your name?\n")
	text, _ := personName.ReadString('\n')

	fmt.Printf("What's up " + text)
	fmt.Printf("Nice to meet you, dog. My name is Eliza, I'm Alexas cooler older sister.\n")
	fmt.Printf("Shall I randomly choose a planet for you to visit? (Y or N)\n")
	travel()

}

func randomPlanet(){
	fmt.Printf("Random JSON struct etc.")
}

func choosePlanet(){
	var 
	fmt.Printf("What planet would you like to learn about?")
}

func travel(){
	//Creating the response variable
	var response string
	fmt.Scanln(&response)

	//If Y, goto randomPlanet function
	if response == "Y"{
		randomPlanet()
	//If N, goto choosePlanet function
	} else if response == "N"{
		choosePlanet()
	//Else repeat question
	} else{
		fmt.Printf("Could you repeat that?\n")
		travel()
	}
}