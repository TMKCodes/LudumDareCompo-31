package main

import (
	"fmt"
	"math"
	"time"
	"runtime"
	sf "bitbucket.org/krepa098/gosfml2"
)

func init() {
	runtime.LockOSThread();
}


func getAngle(a sf.Vector2f, b sf.Vector2f) float64 {
	delta := sf.Vector2f{b.X - a.X, b.Y - a.Y};
	return math.Atan2(float64(delta.Y), float64(delta.X));
}

func main() {
	ticker := time.NewTicker(time.Second / time.Duration(60)); // 60 ticks per second
	window := sf.NewRenderWindow(sf.VideoMode{1920, 1080, 32}, "LD31 preparation", sf.StyleDefault, sf.DefaultContextSettings());
	snowmanSprite, err := sf.NewRectangleShape();
	if err != nil {
		return;
	}
	snowmanSprite.SetSize(sf.Vector2f{10, 10});
	snowmanSprite.SetPosition(sf.Vector2f{0, 250});
	snowmanSprite.SetFillColor(sf.Color{128,128,128,255});
	extraSprite, err := sf.NewRectangleShape();
	if err != nil {
		return;
	}
	extraSprite.SetSize(sf.Vector2f{10, 10});
	extraSprite.SetPosition(sf.Vector2f{1000, 300});
	extraSprite.SetFillColor(sf.Color{0,0,0,255});
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

		angle := getAngle(extraSprite.GetPosition(), snowmanSprite.GetPosition());
		if sf.KeyboardIsKeyPressed(sf.KeyUp) == true {
			snowmanSprite.Move(sf.Vector2f{0, -15});
		}
		if sf.KeyboardIsKeyPressed(sf.KeyDown) == true {
			snowmanSprite.Move(sf.Vector2f{0, 15});
		}
		if sf.KeyboardIsKeyPressed(sf.KeyLeft) == true {
			snowmanSprite.Move(sf.Vector2f{-15, 0});
		}
		if sf.KeyboardIsKeyPressed(sf.KeyRight) == true {
			snowmanSprite.Move(sf.Vector2f{15, 0});
		}
		if sf.KeyboardIsKeyPressed(sf.KeySpace) == true {
			snowmanSprite.Move(sf.Vector2f{float32(math.Cos(angle) * 15), float32(math.Sin(angle) * 15)});
		}


		position := snowmanSprite.GetPosition();
		fmt.Printf("Angle: %v\n", angle);
		fmt.Printf("Snowman Position: %v\n", position);
		window.Clear(sf.Color{255,255,255,0});
		window.Draw(extraSprite, sf.DefaultRenderStates());
		window.Draw(snowmanSprite, sf.DefaultRenderStates());
		window.Display();
	}
}
