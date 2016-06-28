package main

import (
	"fmt"
)

const FOUND_INTRUDER bool = true

func main() {
	camera := online()
	status := "Idle"
	if FOUND_INTRUDER == true {
		status = startRecording(camera)
	}

	if status != "Idle" && status != "Recording" {
		status = "Recording"
	}
	fmt.Println("Status:", status)
}

// EDITABLE OMIT
// необязательно возвращать из функции именно тип Camera
//определяем свой тип
func online() RecordingDevice {
	return FakeCamera{name: "Fake Camera"}
}

type FakeCamera struct {
	name string
}

func (f FakeCamera) record() string {
	return "Idle"
}

// UNEDITABLE OMIT

type RecordingDevice interface {
	record() string
}

type Camera struct {
	name string
}

func startRecording(device RecordingDevice) string {
	return device.record()
}

func (c Camera) record() string {
	if FOUND_INTRUDER {
		return "Recording"
	} else {
		return "Idle"
	}
}
