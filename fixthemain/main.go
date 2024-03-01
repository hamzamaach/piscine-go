package main

import (
	"github.com/01-edu/z01"
)

type Door struct {
	state int
}

const (
	OPEN  = 0
	CLOSE = 1
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(door Door) bool {
	PrintStr("Is the Door opened ?")
	return door.state == OPEN
}

func IsDoorClose(door Door) bool {
	PrintStr("Is the Door closed ?")
	return door.state == CLOSE
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(*door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
