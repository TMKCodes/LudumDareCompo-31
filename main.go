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

type bullet struct {
	bulletSprite *sf.RectangleShape
	direction float64
}

func getAngle(a sf.Vector2f, b sf.Vector2f) float64 {
	delta := sf.Vector2f{b.X - a.X, b.Y - a.Y};
	return math.Atan2(float64(delta.Y), float64(delta.X));
}

func main() {
	ticker := time.NewTicker(time.Second / time.Duration(60)); // 60 ticks per second
	lastShot := time.Now();
	gravity := float32(5);
	shootingRate := time.Second / time.Duration(10); // 30 times per second
	window := sf.NewRenderWindow(sf.VideoMode{1920, 1080, 32}, "LD31 preparation", sf.StyleDefault, sf.DefaultContextSettings());
	window.SetVSyncEnabled(true);
	snowmanSprite, err := sf.NewRectangleShape();
	if err != nil {
		return;
	}
	snowmanSprite.SetSize(sf.Vector2f{100, 100});
	snowmanSprite.SetPosition(sf.Vector2f{0, 250});
	snowmanSprite.SetFillColor(sf.Color{128,128,128,255});

	bullets := make([]*bullet, 0);
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

		if sf.KeyboardIsKeyPressed(sf.KeyW) == true {
			snowmanSprite.Move(sf.Vector2f{0, -10});
		}
		if sf.KeyboardIsKeyPressed(sf.KeyS) == true {
			snowmanSprite.Move(sf.Vector2f{0, 10});
		}
		if sf.KeyboardIsKeyPressed(sf.KeyA) == true {
			snowmanSprite.Move(sf.Vector2f{-10, 0});
		}
		if sf.KeyboardIsKeyPressed(sf.KeyD) == true {
			snowmanSprite.Move(sf.Vector2f{10, 0});
		}
		if sf.IsMouseButtonPressed(sf.MouseLeft) == true {
			if (lastShot.Add(shootingRate)).UnixNano() < time.Now().UnixNano() {
				lastShot = time.Now();
				newBullet := new(bullet);
				newBullet.bulletSprite, _ = sf.NewRectangleShape();
				newBullet.bulletSprite.SetSize(sf.Vector2f{5, 5});
				snowmanPosition := snowmanSprite.GetPosition();
				snowmanSize := snowmanSprite.GetSize();
				newBulletPosition := sf.Vector2f{snowmanPosition.X + (snowmanSize.X / 2), snowmanPosition.Y + (snowmanSize.Y / 2)};
				newBullet.bulletSprite.SetPosition(newBulletPosition);
				newBullet.bulletSprite.SetFillColor(sf.Color{0, 0, 0, 255});
				mousePosition := window.MapPixelToCoords(sf.MouseGetPosition(window), window.GetDefaultView());
				newBullet.direction = getAngle(newBulletPosition, sf.Vector2f{float32(mousePosition.X), float32(mousePosition.Y)});
				if len(bullets) < 100 {
					bullets = append(bullets, newBullet);
				} else {
					bullets = append(bullets[1:100], newBullet);
				}
			}
		}
		fmt.Printf("len(bullets) = %v\n", len(bullets));
		for i := 0; i < len(bullets); i++ {
			x := float32(math.Cos(bullets[i].direction)) * 20;
			y := float32(math.Sin(bullets[i].direction)) * (20 - gravity);
			bullets[i].bulletSprite.Move(sf.Vector2f{x, y});
		}

		snowmanPosition := snowmanSprite.GetPosition();
		if snowmanPosition.Y+gravity < 800 {
			snowmanSprite.Move(sf.Vector2f{0, gravity});
		}
		window.Clear(sf.Color{255,255,255,0});
		for i := 0; i < len(bullets); i++ {
			window.Draw(bullets[i].bulletSprite, sf.DefaultRenderStates());
		}
		window.Draw(snowmanSprite, sf.DefaultRenderStates());
		window.Display();
	}
}
