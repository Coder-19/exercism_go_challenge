package main

import (
	"fmt"
)

// constant for declaring the time  for the lasagana recipie preparation
const OvenTime = 40 // here 40 refers to 40 minutes

// the code below is used to create a function to return the time left
// for the lasagan recipie to be in the oven
func remainingOvenTime(currentTime int) int {
	// the code below is used to return the time left for the lasagana to be cooked
	return OvenTime - currentTime
}

// the code below is used to create a function that takes the number of layers
// (of type int) present in the lasagan dish as input and returns how many minutes
// we have spent initially preparing the lasagan
func preparationTime(numberOfLayers int) int {
	// the code below is used to return the preparation time for creating the
	// lasagana
	return 2 * numberOfLayers
}

// the code below is used to create a function to return the time taken for preparing
// the lasaga by adding the remainingTime with the preparation time and returning it

// the method below takes number of layers of lasagana as input and the time taken
// by the lasagana in the oven as input and returns an integer variable
func elapsedTime(numberOfLayersOfLasagana int, timeTakenByLasaganInOven int) int {
	// the code below is used to add the data returned by teh preparationTime() method and the timeTakenByLasaganInOven
	// for getting the total time taken by the lasagana for preparation

	// the preparationTime() function takes the time taken by the lasagana in the
	// oven as input 
	var totalTimeTakenByLasagana int = preparationTime(numberOfLayersOfLasagana) + timeTakenByLasaganInOven

	// the code below is used to return the totalTimeTakenByLasagan for preparation
	return totalTimeTakenByLasagana
}

func main() {
	// calling the remaingOvenTime() function and passing the currentTime to get
	// the time left for the lasagan to be made and saving the result in a
	// variable of type int
	remainingTime := remainingOvenTime(30)

	// the code below is used to print the remaining time left to make lasagan to the
	// console using the fmt.Println() method
	fmt.Println("The time remaining is:", remainingTime)

	// the code below is used to call the method for getting the preparation time
	// and then passing the number of layers as input for getting the preparation
	// time for preparing the lasagan and saving that value in a variable
	preparationTimeRequired := preparationTime(2)

	// the code below is to use the fmt.Println() method to print the time
	// required for the preparation of lasagan
	fmt.Printf("The preparation time required is: %v \n", preparationTimeRequired)

	// the code below is used to call the method elapsedTime() and then passing the
	// numberOfLayersOfLasagana and time taken by lasagana in the oven as input for getting
	// the total time taken by the lasagan for preparation
	var totalTimeRequiredInPreparation int = elapsedTime(1, 30)

	// the code below is used to print the total time required to the console
	fmt.Println("The total time required is:", totalTimeRequiredInPreparation)

}
