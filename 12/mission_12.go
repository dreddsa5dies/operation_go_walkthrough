// "Nice work!" Agent Getter says wiping the sweat off of his brow. "The field can be kind of fun, huh?"
// "I can't wait to get back to the office," you reply exhaustedly.
// As the sun begins to set, you hear the sound of a half dozen helicopters flying towards Epoch's compund. You watch from a bluff as the raid goes flawlessly. Rand is arrested along with the corrupt agents. You and Agent Getter are greeted as heros by the rescue squad.
// "I'm just happy to do my job," you reply to a seemlingly endless line of congradulatory remarks.
// "Speaking of which, I can't wait to get my hands on those USBs," Agent Getter says finally breaking into a grin.
// "Absolutely! Right after me." you smile back.
// No sooner does the flight land back at headquarters when you're at your workstation ready to get back to work.
// "OK, Epoch, time to see what you were really up to," you say as you settle into your plush work chair.

// "Хорошая работа!" Агент Getter говорит вытирая пот со лба. "Поле может быть забавно, да?"
// "Я не могу дождаться, чтобы вернуться в офис," Вы отвечаете изнеможении.
// По мере того как солнце начинает устанавливать, вы услышите звук полутора десятков вертолетов, выполняющих полеты в сторону compund эпохи. Вы смотрите с блефом, как рейд идет без сбоев. Rand арестован вместе с коррумпированными агентами. Вы и агент Getter встречают, как героев бригадой скорой помощи.
// "Я просто счастлив, чтобы делать свою работу", вы отвечаете на seemlingly бесконечной линии congradulatory замечаний.
// "Говоря об этом, я не могу ждать, чтобы мои руки на этих USBs," говорит агент Getter, наконец, врываясь в улыбке.
// "Абсолютно! Сразу же после меня." Вы улыбаетесь назад.
// Не раньше, делает полет землю обратно в штаб-квартире, когда вы на вашей рабочей станции готовы вернуться к работе.
// "ОК, Великая Эпоха, время, чтобы увидеть то, что вы были действительно до", вы говорите, как вы поселиться в вашем плюшевые рабочее кресло.

/******************************
** Mission: Secrets Revealed **
*******************************
**
** IMPORTANT: You only get one chance to run your code on this level
** before your score is recorded on the leaderboard.
**
** You download and aggregate the USB data into a single file. As you
** sit down to dive in, your initial enthusiasm turns to concern when
** you see that as soon as some files are read, other files lock.
** It's absolutely critical for the case against Epoch to get as much
** data as you can. Epoch's file system allows you to input up to ten
** switching commands on the file locks. What's worse is you can't see
** which files are locked until you attempt to read the file. Everyone
** is counting on you to deliver as many readable files as possible.
**
**/

package main

import (
	"fmt"
	"sort"
)

func main() {
	commands := setCommands()
	files := createFiles()
	files.run(commands)
	megabytesCollected := files.countMegabytes()
	fmt.Println("Megabytes Collected:", megabytesCollected)
}

type File struct {
	megabytes int
	locked    bool
}

type Files []File

func (files Files) Len() int {
	return len(files)
}

func (files Files) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

//-----------------------------------------------
// изменять здесь

func (files Files) Less(i, j int) bool {
	return files[i].megabytes < files[j].megabytes
}

func setCommands() []string {
	commands := make([]string, 10)
	for i := 0; i < len(commands)-1; i++ {
		commands[i] = "flip_all"
	}

	return commands
}

//-----------------------------------------------

func (files Files) run(commands []string) {
	for i := 0; i < 10; i++ {
		switch commands[i] {
		case "sort":
			sort.Sort(files)
		case "chain":
			files.chainLocks()
		case "flip_all":
			files.flipAllLocks()
		case "flip_intervals":
			files.flipIntervals()
		}
	}
}

func (files Files) flipIntervals() {
	for i := len(files) - 1; i > 0; i-- {
		if i%4 == 0 &&
			files[i].megabytes > files[i-1].megabytes {
			files[i].locked = !files[i].locked
		}
	}
}

func (files Files) flipAllLocks() {
	for i := range files {
		files[i].locked = !files[i].locked
	}
}

func (files Files) chainLocks() {
	chain_type := files[len(files)-1].locked
	for i := len(files) - 2; i > 0; i-- {
		if files[i].locked != chain_type &&
			files[i].megabytes > files[i-1].megabytes {
			files[i].locked = chain_type
		} else {
			return
		}
	}
}

func createFiles() Files {
	files := make(Files, 25)
	megabytes := 0

	for i := range files {
		files[i] = File{
			megabytes: fileSizes[i],
			locked:    fileLocks[i],
		}
		megabytes += fileSizes[i]
	}
	fmt.Println("Megabytes Available:", megabytes)
	return files
}

func (files Files) countMegabytes() int {
	megabytes := 0
	for i := range files {
		if !files[i].locked {
			megabytes += files[i].megabytes
		}
	}
	return megabytes
}
