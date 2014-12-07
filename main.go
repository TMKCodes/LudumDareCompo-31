package main

import (
	"fmt";
	"time";
	"runtime";
)

func main() {
	runtime.LockOSThread();
	fmt.Printf("Theme for the game was given in Ludum Dare 31 and it was voted as Entire Game on One Screen\n");
	ticker := time.NewTicker(time.Second / time.Duration(60));
	snowball := NewGame("LD31 Snowball", 1920, 1080, 32, true);
	for snowball.RenderWindow.IsOpen() {
		<-ticker.C
		snowball.Update();
		snowball.Draw();
	}
}
