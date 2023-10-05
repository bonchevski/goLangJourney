package main

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntildrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntildrinking = 21 - age

	if yearsUntildrinking < 0 {
		yearsUntildrinking = 0
	}

	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilAdult, yearsUntildrinking, yearsUntilCarRental
}

func main() {
	yearsUntilAdult, yearsUntildrinking, yearsUntilCarRental := yearsUntilEvents(44)
	println("You will be an adult in", yearsUntilAdult, "years")
	println("You will be able to drink in", yearsUntildrinking, "years")
	println("You will be able to rent a car in", yearsUntilCarRental, "years")
}
