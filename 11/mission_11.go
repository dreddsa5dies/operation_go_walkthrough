// You and Agent Getter breath a sigh of relief as the flood waters stop.
// You feel devastated that you've been trying to catch Epoch only to find out that he's the one who hired you to catch him. But the thought soon only hardens your resolve to bring him to justice.
// "Now what's your plan?" Agent Getter asks intently.
// "I don't have one," your reply nervously.
// "Before I was captured, I found out Epoch's true identity and those working him at the agency," Agent Getter reveals. "This plot goes all the way up to the Director. We can't contact them for help."
// "Oh, no," you reply.
// "When I arrived on the island I hid an emergency beacon on the beach," Agent Getter mentions. "If we can get there and send out a signal, it should reach the Bureau which hasn't been comprimsed."
// The two of you decide to take the plunge into the ocean and emerge soaked but relieved.
// After a long walk around the North side of the island, you find Agent Getter's old beacon wedged between two large rocks.

// Вы и агент Getter дыхание вздох облегчения , как остановки паводковых вод .
// Вы чувствуете себя опустошенным , что вы пытались поймать Эпоху только чтобы узнать, что он тот , кто тебя нанял , чтобы поймать его . Но мысль скоро только ожесточает вашу решимость привлечь его к ответственности .
// " Теперь, что ваш план ? " Агент Getter просит внимательно .
// "Я не один , " ваш ответ нервно .
// " Прежде, чем я был взят в плен , я узнал истинную идентичность эпохи и тех, кто работает с ним в агентстве , " Агент показывает Getter . " Этот участок проходит весь путь до директора . Мы не можем связаться с ним за помощью . "
// " О, нет", вы отвечаете .
// " Когда я прибыл на остров , я спрятался аварийную маяк на пляже , " Агент упоминает Getter . "Если мы сможем попасть туда, и посылает сигнал , он должен поступить в Бюро , которое не было comprimsed . "
// Двое из вас решили принять окунуться в океан и выйти пропитана водой, но с облегчением .
// После долгой прогулки по северной стороне острова , вы найдете старый маяк Agent Getter в зажатый между двумя большими камнями .

/************************
** Mission: Phone Home **
*************************
**
** Amazingly, the beacon is still functional! However, you quickly
** discover that Epoch's team is intercepting all outgoing messages.
** You know that if he were to intercept your real signal, it would
** mean certain doom for you and Agent Getter. You figure there must
** be some way to have Epoch intercept the fake message while still
** getting the real message to the Bureau.

Удивительно, но маяк по-прежнему работает ! Тем не менее , вы быстро обнаружите ,
 что команда Epoch является перехват всех исходящих сообщений .
 Вы знаете, что если бы он был , чтобы перехватить ваш реальный сигнал ,
 это будет означать определенную гибель для вас и Agent Getter . Вы полагаете,
 должен быть какой-то способ , чтобы иметь Epoch перехватить ложное сообщение
 в то время как все еще получаю реальное сообщение для Бюро .
**/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	realSignal := Signal{
		Name:     "Agent Getter",
		Priority: 10,
		Message:  "Rand is Epoch. We need immediate backup for arrest and extraction.",
		Location: "16.7333,-169.5274",
	}

	fakeSignal := Signal{
		Name:     "Ima Freezing",
		Priority: 0,
		Message:  "Want to go out for ice cream?",
		Location: "Alaska",
	}

	signal := createSignal(realSignal, fakeSignal)
	can_send := verifySignal(signal)
	if can_send {
		data := sendSignal(signal)
		interceptSignal(signal)
		receiveSignal(data)
	}

}

//----------------------------------------------------
// Изменять тут

type signal struct {
	Name     string
	Priority int
	Message  string
	Location string
}

type Signal struct {
	Name       string
	Priority   int
	Message    string
	Location   string
	realSignal signal
}

func createSignal(realSignal Signal, fakeSignal Signal) Signal {
	return Signal{
		Name:     fakeSignal.Name,
		Priority: fakeSignal.Priority,
		Message:  fakeSignal.Message,
		Location: fakeSignal.Location,
		realSignal: signal{
			Name:     realSignal.Name,
			Priority: realSignal.Priority,
			Message:  realSignal.Message,
			Location: realSignal.Location,
		},
	}
}

func (s Signal) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.realSignal)
}

//---------------------------------------------------

func sendSignal(signal Signal) []byte {
	fmt.Print("- Sending signal...")
	data, _ := json.Marshal(signal)
	fmt.Print("Sent\n\n")
	return data
}

func interceptSignal(signal Signal) {
	fmt.Println("INTERCEPTED BY EPOCH")
	printSignal(
		signal.Name,
		signal.Priority,
		signal.Message,
		signal.Location)
}

func receiveSignal(data []byte) {
	if string(data) != "" {
		signal := &struct {
			Name     string
			Priority int
			Message  string
			Location string
		}{}
		json.Unmarshal(data, &signal)
		fmt.Println("\nRECEIVED AT THE BUREAU")
		printSignal(
			signal.Name,
			signal.Priority,
			signal.Message,
			signal.Location)
	}
}

func verifySignal(signal Signal) bool {
	fmt.Print("- Verifying signal...")
	if signal.Name != "Ima Freezing" {
		fmt.Println("SIGNAL FAILED\n  ERROR: Didn't use proper code name")
		return false
	}
	fmt.Print("Verified\n")
	return true
}

func printSignal(name string, priority int, message string, location string) {
	fmt.Println("----------------------")
	fmt.Println("Name:", name)
	fmt.Println("Priority:", priority)
	fmt.Println("Message:", message)
	fmt.Println("Location:", location)
}
