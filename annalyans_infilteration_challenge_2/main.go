package main

import "fmt"

// the code below is used to create a method to check that if a fast attack
// can be made or not if the knight is awake

// the method below takes a boolean value named isKnightAwake as input and
// if the value is true then return false else return true

// the method below returns boolean value as output
func canFastAttack(knightIsAwake bool) bool {
	// the code below is used to check that if the value of isKnightAwake is true
	// then return false else return true
	if knightIsAwake {
		return false
	} else {
		return true
	}
}

// the code below is used to create a method called canSpy() that takes three
// integer values i.e. whether the knight is awake or whether the prisoner is awake
// or whether the archer is awake
// the method checks that if the archer or the knight is awake then the group
// can be spied that is true else false

// the method below returns a boolean value as output
func canSpy(knightIsAwake bool, archerIsAwake bool, prisonerIsAwake bool) bool {
	// the code below is used to check that if the prisoner is awake or not
	if prisonerIsAwake {
		// the code below is used to check that the knight or the archer is awake
		if knightIsAwake || archerIsAwake {
			return true

		} else {
			// if both the knight and the archer is not awake then return false
			return true
		}
	} else {
		// the code below is used to check that the knight or the archer is awake
		if knightIsAwake || archerIsAwake {
			return true
		} else {
			// if both the knight and the archer is not awake then return false
			return false
		}
	}
}

// the code below is used to create a method to check that whether the prisoner can
// be signled or not

// the method below takes two inputs first one is the value that the archer is
// awake or not and the second one is that the prisoner is awake or not

// the method returns a boolean value
func canSignalPrisoner(archerIsAwake bool, prisonerIsAwake bool) bool {
	// the code below is used to check that if the archer is awake and the
	// prisoner is sleeping then the prisoner cannot be signled
	if archerIsAwake && !prisonerIsAwake {
		return false

		// the code below is used to check that if the archer is sleeping and
		// the prisoner is awake then the prisoner can be signled
	} else if !archerIsAwake && prisonerIsAwake {
		return true

		// the code below is used to check that if both the archer and the
		// prisoner are sleeping then the prisoner cannot be signled
	} else if !archerIsAwake && !prisonerIsAwake {
		return false

		// the code below is used to check that if the archer and the
		// prisoner both are awake that is the last case then also the prisoner
		// cannot be signled
	} else {
		return false
	}
}

// the code below is used to create a method for checking that whether the prisoner
// can be freed or not

// the method below takes four inputs that is
// 1. the knight is awake or not
// 2. archer is awake or not
// 3. prisoner is awake or not
// 4. pet dog is present

// the method below checks the above four conditions and returns true or false
// based on the condition

// the method below returns a boolean value
func canFreePrisoner(knightIsAwake bool, archerIsAwake bool, prisonerIsAwake bool, petDogIsPresent bool) bool {
	// the code below is used to check that whether the prisoner is awake or not
	if prisonerIsAwake {
		// the code below is used to check that if the pet dog is present
		// or not
		if petDogIsPresent {
			// if the pet dog is present then we need to check that if the archer is
			// sleeping or not
			if !archerIsAwake && knightIsAwake {
				return true

				// if the archer and the knight are sleeping then also the prisoner
				// can be freed
			} else if !archerIsAwake && !knightIsAwake {
				return true

				// the code below is used to check that if the archer is awake but
				// the knight is sleeping then also the prisoenr cannot be freed
			} else if archerIsAwake && !knightIsAwake {
				return false

				// the code below is used to check that if both the archer and the
				// knight are awake then also the prisoner cannot be freed
			} else {
				return false
			}

			// the code below is used to check the conditions that if the dog is not
			// present
		} else {
			// if the archer and the knight are sleeping then also the prisoner
			// can be freed
			if !archerIsAwake && !knightIsAwake {
				return true

				// the code below is used to check that if the archer is awake or
				// the knight is awake then also the prisoenr cannot be freed
			} else {
				return false

			}
		}

		// the code below is used to check that if the prisoner is not awake
		// then the prisoner cannot be rescued
	} else {

		// the code below is used to check that if the dog is present
		// but the prisoner, archer and the knight are sleeping then
		// also the prisoner can be freed else the prisoner cannot be freed
		if petDogIsPresent && !knightIsAwake && !archerIsAwake {
			return true

			// the code below is used to check that if the archer and the 
			// prisoner are sleeping but the knight is awake and the 
			// dog is present then also we can free the prisoner
		} else if knightIsAwake && !archerIsAwake && petDogIsPresent{
			return true
		}else {
			return false
		}
	}
}

func main() {
	// the code below is used to create a property for getting the value that
	// the knight is awake or not
	var knightIsAwake bool = true

	// the code below is used to call the canFastAttack() method and pass the
	// value of the knightIsAwake variable and save the output returned by the
	// canFastAttack() method in a variable called fastAttackPossible
	fastAttackPossible := canFastAttack(knightIsAwake)

	// the code below is used to print the value of fastAttackPossible to console
	fmt.Printf("Fast attack Possible: %v \n", fastAttackPossible)

	// the code below is used to create a property for checking that whether the knight is awake
	// or not
	var knightAwake bool = false

	// the code below is used to create property for checking that whether the
	// archer is awake or not
	var archerAwake bool = true

	// the code below is used to check that whether the prisoner is awake or not
	var prisonerAwake = false

	// the code below is used to create a property for getting the value from the
	// canSpy() method for checking that the prisoner can be spied or not
	canSpyOnPrisoner := canSpy(knightAwake, archerAwake, prisonerAwake)

	// the code below is used to print the value of the canSpyOnPrisoner property
	// for checking that whether the prisoner can be spied or not
	fmt.Println("Can Spy on Prisoner:", canSpyOnPrisoner)

	// the code below is to use the fmt.Println() for printing the value that
	// whether the prisoner can be signled or not
	fmt.Println("The Prisoner can be signled:", canSignalPrisoner(archerAwake, prisonerAwake))

	// the code below is to use the fmt.Println() for printing the value that
	// whether the prisoner can be freed or not
	fmt.Println("The Prisoner can be freed:", canFreePrisoner(false, true, false, false))
}
