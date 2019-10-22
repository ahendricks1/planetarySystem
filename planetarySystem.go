package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

//Creating an array struct to store planets
type Planets struct {
	Planets []PlanetDesc `json:"planets"`
}

//Creating a struct for planets names and descriptions
type PlanetDesc struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {

	//Initial stdout to user with name request
	personName := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to the Solar System!\n")
	fmt.Printf("There are 9 planets to explore.\n")
	fmt.Printf("What is your name?\n")
	text, _ := personName.ReadString('\n')

	//Stdout with name and start of planet selection
	fmt.Printf("What's up " + text)
	fmt.Printf("My name is Eliza, I'm Alexas cooler older sister.\n")
	fmt.Printf("Shall I randomly choose a planet for you to visit? (Y or N)\n")
	travel()
}

func travel() {

	//Creating the response variable
	var response string
	var planets Planets
	fmt.Scanln(&response)

	//If Y, boolean set to true
	if response == "Y" {
		planetChoice(planets, response, true)

		//If N, boolean remains false
	} else if response == "N" {
		fmt.Println("What planet would you like to learn about?")
		planetChoice(planets, response, false)

		//Else repeat question
	} else {
		fmt.Printf("Could you repeat that?\n")
		travel()
	}
}

func planetChoice(planets Planets, userResponse string, randomizePlanet bool) {

	//Creating some variables, and random number for randomized planet
	seed := rand.NewSource(time.Now().UnixNano())
	newrand := rand.New(seed)
	var x = newrand.Intn(9)
	var specificPlanet string
	var index int
	var planetFound bool

	//Opening the planets json file
	jsonFile, err := os.Open("planets.json")

	//Checking if no errors occur, if any do print error
	if err != nil {
		fmt.Println(err)
	}

	//Defer closing of the json file so it can be parsed
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &planets)

	//If the user wants a random planet, give it to them
	if randomizePlanet {
		fmt.Println(planets.Planets[x].Name + ": " + planets.Planets[x].Description)

		//Else if they want a specific planet, have them type in the name
	} else {
		fmt.Scanln(&specificPlanet)

		//Iterate through planets to find specific planet chosen
		for i := 0; i < len(planets.Planets); i++ {
			//If the users entry exists in the array, print name and description
			if specificPlanet == planets.Planets[i].Name {
				index = i
				planetFound = true
			}
		}
		//Check if planet exists in the array, if so, print out name and description
		if planetFound {
			fmt.Println(planets.Planets[index].Name + ": " + planets.Planets[index].Description)

			//Else repeat question
		} else {
			fmt.Println("That's not a planet...try again...")
			planetChoice(planets, specificPlanet, false)
		}

	}
}
