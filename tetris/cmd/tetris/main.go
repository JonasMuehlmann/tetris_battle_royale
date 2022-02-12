package main

import (
	"os"
	"os/exec"
	"tetris/internal"
	"time"
)

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	buffer := make([]byte, 1)

	playfield := internal.MakePlayField()

	playfield.StartGame()

	printTicker := time.NewTicker(100 * time.Millisecond)

	go func() {
		for {
			os.Stdin.Read(buffer)

			switch buffer[0] {
			case 'q':
				playfield.RotateBlockCounterClockwise()
			case 'e':
				playfield.RotateBlockClockwise()
			case 'a':
				playfield.MoveBlockLeft()
			case 'd':
				playfield.MoveBlockRight()
			case 's':
				playfield.MoveBlockDown()
			case 'c':
				playfield.HardDropBlock()
			case 'v':
				playfield.ToggleSoftDrop()
			}
		}
	}()

	go func() {
		for {
			playfield.PrintPlayField()
			<-printTicker.C
			// Terminal escape code go brrrrr
			print("\033[H\033[2J")
		}
	}()

	<-playfield.GameStop
	printTicker.Stop()

	print("\033[H\033[2J")
	print("Game over!\n")
}
