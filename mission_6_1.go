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

func setupLasers() [7]Laser {
	var lasers [7]Laser
	lasers[0] = Laser{1, true}
	lasers[1] = Laser{2, true}
	lasers[2] = Laser{3, true}
	lasers[3] = Laser{4, true}
	lasers[4] = Laser{5, true}
	lasers[5] = Laser{6, true}
	lasers[6] = Laser{7, true}
	return lasers
}

func testGrid(laserGrid LaserGrid) {
	fmt.Println("Status:", laserGrid.status)
}

type LaserGrid struct {
	status string
	lasers [7]Laser
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
