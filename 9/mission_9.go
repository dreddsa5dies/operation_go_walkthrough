/*************************
** Mission: Hidden Lair **
**************************
**
** Unbelievably, the sequence of buttons to unlock the door is
** sitting right there in the code. You try your luck, but the
** sequence is getting scrambled and the door still remains locked.
** This has to be the first time you've ever seen a race condition
** double as a security feature. If you can execute the code safely,
** it looks like the door should unlock for you.
**
**/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var buttons []string = []string{"red", "blue", "green", "yellow", "purple"}

func main() {
	rand.Seed(111009)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(buttons))

	sequence := make([]Button, 0)

	for x := 0; x < len(buttons); x++ {
		go addButton(x, &sequence, &waitGroup)
	}

	waitGroup.Wait()

	fmt.Println("Sending Code Sequence...")
	for i, button := range sequence {
		fmt.Println(i+1, ":", button.color)
	}
}

//------------------------------------------------
// Изменять только тут

func setButton(x int, sequence *[]Button) {
	fmt.Println(Button{buttons[x]})
	//------------------------------------------------

	new_button := Button{buttons[x]}
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	*sequence = append(*sequence, new_button)
}

func addButton(x int, sequence *[]Button, waitGroup *sync.WaitGroup) {
	setButton(x, sequence)
	waitGroup.Done()
}

type Button struct {
	color string
}