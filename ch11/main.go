package main

import "fmt"

func main() {
	player := TapeRecorder{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}

func playList(device TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("TapePlayer Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("TapePlayer Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Record() {
	fmt.Println("TapeRecorder Recording!")
}
func (t TapeRecorder) Play(song string) {
	fmt.Println("TapeRecorder Playing", song)
}
func (t TapeRecorder) Stop() {
	fmt.Println("TapeRecorder Stopped!")
}
