package main

import (
	"fmt"
)

func main() {

	elevation := 4000
	destinationNorth := 1000
	destinationWest := 1000

	windNorth := -4
	windWest := -4

	deployChuteAtElevation := 2000

	landingLocation := simDrop(elevation,
		destinationNorth,
		destinationWest,
		windNorth,
		windWest,
		deployChuteAtElevation)

	fmt.Println("Initiate deploy at:",
		deployChuteAtElevation,
		"meters -", landingLocation)
}

func simDrop(elevation int,
	dest_north int,
	dest_west int,
	windNorth int,
	windWest int,
	deployChuteAtElevation int) string {

	locationNorth := 0
	locationWest := 0

	deployed := false
	for elevation > 0 {
		if elevation == deployChuteAtElevation {
			deployed = true
		}
		if !deployed {
			locationNorth -= windNorth / 4
			locationWest -= windWest / 4
		} else {
			locationNorth -= windNorth
			locationWest -= windWest
		}
		elevation -= 10
	}
	landingLocation := "You missed the target!"
	if locationNorth == dest_north && locationWest == dest_west {
		landingLocation = "You hit the target!"
	}

	return landingLocation
}
