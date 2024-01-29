package main

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actualMinutes int) int {
	return OvenTime - actualMinutes
}

func PreparationTime(numberOfLayers int) int {
    prepareTimePerLayer := 2

	return prepareTimePerLayer * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	preperationTime := PreparationTime(numberOfLayers)

	return preperationTime + actualMinutesInOven
}

func main() {
	fmt.Println(RemainingOvenTime(30))
	fmt.Println(ElapsedTime(3, 20))
}