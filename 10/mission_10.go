// ** Mission: Epoch's Trap **
// ***************************
// **
// ** This could be it if you can't stop the flow of water quickly.
// ** Epoch's security is much improved in his lair with large chunks
// ** of code hidden from you.  Luckily, Agent Getter introduces you to
// ** the reflect package which could shed some light on the missing code,
// ** hopefully in time to stop the water from filling the room.
// **
// **

package main

import (
  "fmt"
  "reflect"
  "time"
)

func main() {
  c := initControls()
  c.gallonsPerMin = 700
//------------------------------------------------
// Изменять только тут

  v := reflect.TypeOf(c)
  fmt.Println(v)

//------------------------------------------------

  capacity := 20000
  filled := 0

  for {
    filled += c.flow()
    if filled >= capacity {
      fmt.Println("Waterline:", filled)
      fmt.Println("Room filled!")
      return
    } else {
      fmt.Println("Waterline:", filled)
      time.Sleep(20 * time.Millisecond)
    }
    if c.emergencyShutoff() {
      fmt.Println("Emergency shutoff activated")
      return
    }
  }
}


// As the door unlocks, Agent Getter rushes out to meet you and you quickly introduce yourself.
// "I'm Agent Marshal. I was on a mission to secure these USB drives but some agents have flipped and they're helping Epoch," you say looking for answers.
// "I'm afraid you're in serious danger. Epoch is not who you think he is," Agent Getter replies.
// Just then, a shadowy figure enters the far end of the lair. It looks like your boss Rand!
// "Agent Marshal, Agent Getter, it looks like you're in over your heads this time," shouts Rand.
// "We sure are!" you reply. "And we sure are glad to see you! How did you find us here?"
// "I am Epoch," Rand replies in a low booming voice. "You were getting too close to finding my true identity, so I sent you on this mission never to return. Despite your heroics so far, this is the end for you."
// You're stunned silence fills the room as your impending fate becomes ever more clear.
// As quickly as he appeared, Rand disappears behind a corner and the door slams shut. Suddenly, valves open up to the sea and start flooding the lair...