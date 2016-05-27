/****************************
** Mission: Need to Escape **
*****************************
**
** It looks like this Jet Ski would make a better rock than an escape
** plan as it's been hardwired not to start.  Good thing you still
** have a trick or two up your sleeve that could save the day yet.
** You need this engine to start, and fast.
**
**/

package main

import (
	"fmt"
)

func main() {
	//------------------------------------------------
	// Изменять только тут

	//------------------------------------------------

	canIStart := canStart("wrongpassword")
	if canIStart {
		fmt.Println("Starting engine...")
	} else {
		fmt.Println("Engine Off.")
	}
}

func canStart(password string) bool {
	if password == "3p0cHp@$$" {
		return true
	}
	return false
}
