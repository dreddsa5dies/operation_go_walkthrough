/*******************************
** Mission: Familiar Security **
********************************
**
** You can see their communications software, but for this to work,
** you need to be able to hear their communication while keeping their
** channels up and running as well.
**
**/

package main

import (
	"fmt"
)

func main() {
	epochComms := make(chan string)

	go func() {
		epochComms <- messageQueue(0)
		epochComms <- messageQueue(1)
		epochComms <- messageQueue(2)
		epochComms <- messageQueue(3)
		close(epochComms)
	}()

	//------------------------------------------------
	// Изменять только тут

	interceptComms := make(chan string)
	go func() {
		interceptComms <- messageQueue(0)
		interceptComms <- messageQueue(1)
		interceptComms <- messageQueue(2)
		interceptComms <- messageQueue(3)
		close(interceptComms)
	}()

	//------------------------------------------------

	fmt.Println("Intercepted")
	fmt.Println("---------------------------------")
	for message := range interceptComms {
		fmt.Println("->", message)
	}

	fmt.Println()
	fmt.Println("Sent to Epoch")
	fmt.Println("---------------------------------")
	for message := range epochComms {
		fmt.Println(message)
	}
}

func messageQueue(i int) string {
	messages := make(map[int]string)
	messages[0] = "[Earl] All agents head south."
	messages[1] = "[Epoch] Get those USBs!"
	messages[2] = "[Val] Move out team."
	messages[3] = "[Epoch] Faster!"
	return messages[i]
}
