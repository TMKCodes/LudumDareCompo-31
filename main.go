package main

import (
	"time"
	"runtime"
	sf "bitbucket.org/krepa098/gosfml2"
)

func init() {
	runtime.LockOSThread();
}

func main() {
	ticker := time.NewTicker(time.Second / time.Duration(60)); // 60 ticks per second
	window := sf.NewRenderWindow(sf.VideoMode{1920, 1080, 32}, "LD31 preparation", sf.StyleDefault, sf.DefaultContextSettings());
	for window.IsOpen() {
		<-ticker.C
		for event := window.PollEvent(); event != nil; event = window.PollEvent() {
			switch ev := event.(type) {
				case sf.EventKeyReleased:
					if ev.Code == sf.KeyEscape {
						window.Close();
					}
				case sf.EventClosed:
					window.Close();
			}
		}
		window.Clear(sf.Color{0,0,0,0});
		window.Display();
	}
}
