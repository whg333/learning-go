package main

import "fmt"

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = TapePlayer{} // 此处必须声明为Player接口，不可使用:=短赋值
	playList(player, mixtape)
	player = TapeRecorder{}
	playList(player, mixtape)
}

func playList(device Player, songs []string) {
	recorder, ok := device.(TapeRecorder) // 类型断言，即接口强转为具体类型
	if ok {
		recorder.Start()
	}
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

// Player 不需要额外语法声明TapePlayer和TapeRecorder为Player接口，只需要实现Player接口的方法即可
type Player interface {
	Play(song string)
	Stop()
}

// TapePlayer 在Go语言中没有类的定义，只有struct结构体和func函数定义，二者分开定义并用接收器关联即为方法
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

func (t TapeRecorder) Start() {
	fmt.Println("TapeRecorder Started!")
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
