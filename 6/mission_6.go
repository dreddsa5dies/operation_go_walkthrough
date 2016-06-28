package main

import (
	"fmt"
)

func main() {
	lasers := setupLasers()

	for i := 0; i < 7; i++ {
		if !lasers[i].isRunning {
			fmt.Println("ALERT! Security system breached.")
			return
		}
	}

	laserGrid := LaserGrid{"Operational", lasers}

	testGrid(laserGrid)

	if laserGrid.status != "Operational" {
		fmt.Println("ALERT! Security system breached.")
		return
	}

	if activeLasers(laserGrid) > 0 {
		fmt.Println(activeLasers(laserGrid), "lasers active.")
	} else {
		fmt.Println("All lasers deactivated!")
	}
}

//----------------------------------------
// Изменять только тут

func setupLasers() [7]*Laser {
	var lasers [7]*Laser
	for i := 0; i < 7; i++ {
		lasers[i] = &Laser{i, true}
	}
	return lasers
}

func testGrid(laserGrid LaserGrid) LaserGrid {
	fmt.Println("Status:", laserGrid.status)
	for _, laser := range laserGrid.lasers {
		laser.isRunning = false
	}
	return laserGrid
}

type LaserGrid struct {
	status string
	lasers [7]*Laser
}

type Laser struct {
	id        int
	isRunning bool
}

//------------------------------------------

func activeLasers(laserGrid LaserGrid) int {
	activeLasers := 0
	for _, laser := range laserGrid.lasers {
		if laser.isRunning {
			activeLasers++
		}
	}
	return activeLasers
}
