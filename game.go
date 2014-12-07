package main

import (
	"fmt";
	"time";
	"strconv";
	"math/rand";
	sf "bitbucket.org/krepa098/gosfml2";
)


type snowball struct {
	Sprite *sf.RectangleShape;
	Location sf.Vector2f;
	Exists bool;
}

type levelone struct {
	ForegroundSprite *sf.RectangleShape;
	BackgroundSprite *sf.RectangleShape;
	Exists bool;
}

type levelzero struct {
	WelcomeText pointstext;
	Exists bool;
}

type leveltwo struct {
	EndText pointstext;
	Exists bool;
}

type snowman struct {
	Sprite *sf.RectangleShape;
	Up bool;
	Exists bool;
}

type pointstext struct {
	Text *sf.Text;
	Exists bool;
}

type music struct {
	Explosions []*sf.Music
	Throws []*sf.Music
}

type game struct {
	Gear int;
	RenderWindow *sf.RenderWindow;
	LevelZero levelzero;
	LevelTwo leveltwo;
	Snowball snowball;
	Level levelone;
	Snowman snowman;
	SnowmanBall snowball;
	Time int64;
	Speed int64;
	Points int;
	SnowmanPoints int;
	Font *sf.Font;
	PointsText pointstext;
	Music music;
}

func NewGame(title string, width uint, height uint, bpp uint, vsync bool) *game {
	game := new(game);
	game.RenderWindow = sf.NewRenderWindow(sf.VideoMode{width, height, bpp}, title, sf.StyleDefault, sf.DefaultContextSettings());
	game.RenderWindow.SetVSyncEnabled(vsync);
	game.Gear = 0;
	// do some other initialization..
	font, err := sf.NewFontFromFile("deja.ttf");
	if err != nil {
		fmt.Printf("sf.NewFontFromFile() call failed: %v\n", err);
	}
	game.Font = font;
	game.InitGear();
	rand.Seed(time.Now().UnixNano());
	var Music music;
	Music.Explosions = make([]*sf.Music, 0);
	explosion, _ := sf.NewMusicFromFile("explosion1.wav");
	Music.Explosions = append(Music.Explosions, explosion);
	explosion, _ = sf.NewMusicFromFile("explosion2.wav");
	Music.Explosions = append(Music.Explosions, explosion);
	explosion, _ = sf.NewMusicFromFile("explosion3.wav");
	Music.Explosions = append(Music.Explosions, explosion);
	explosion, _ = sf.NewMusicFromFile("explosion4.wav");
	Music.Explosions = append(Music.Explosions, explosion);
	Music.Throws  = make([]*sf.Music, 0);
	throw, _ := sf.NewMusicFromFile("throw1.wav");
	Music.Throws = append(Music.Throws, throw);
	throw, _ = sf.NewMusicFromFile("throw2.wav");
	Music.Throws = append(Music.Throws, throw);
	throw, _ = sf.NewMusicFromFile("throw3.wav");
	Music.Throws = append(Music.Throws, throw);
	throw, _ = sf.NewMusicFromFile("throw4.wav");
	Music.Throws = append(Music.Throws, throw);
	game.Music = Music;
	return game;
}

func (this *game) ChangeGearUp() {
	this.Gear += 1;
	fmt.Printf("Gear changed up to %v\n", this.Gear);
	this.InitGear();
}

func (this *game) ChangeGearDown() {
	this.Gear -= 1;
	fmt.Printf("Gear changed down to %v\n", this.Gear);
	this.InitGear();
}

func (this *game) InitGear() {
	if this.Gear == 0 { // the menu gear.
		var l levelzero;
		text, _ := sf.NewText(this.Font);
		l.WelcomeText.Text = text;
		l.WelcomeText.Text.SetString("Welcome to the Snowball fights! Your duty is to win snowmen.\n\nYou can throw snowballs with your mouse!\nRemember to hit the snowmen or they will throw snowballs at you.\n\n\nPress enter to start the game.");
		l.WelcomeText.Text.SetCharacterSize(24);
		l.WelcomeText.Text.SetColor(sf.Color{0,0,0,255});
		l.WelcomeText.Text.SetPosition(sf.Vector2f{300,300});
		l.WelcomeText.Exists = true;
		l.Exists = true;
		this.LevelZero = l;
	} else if this.Gear == 1 { // the playable gear.
		fmt.Printf("Initializing gear 1\n");
		var l levelone;
		lForegroundSprite, err := sf.NewRectangleShape();
		if err != nil {
			fmt.Printf("sf.newRectangleShape call failed: %v\n", err);
		}
		l.ForegroundSprite = lForegroundSprite;
		l.ForegroundSprite.SetSize(sf.Vector2f{float32(1920), float32(1080)});
		l.ForegroundSprite.SetPosition(sf.Vector2f{0.0, 0.0});
		lForegroundTexture, err := sf.NewTextureFromFile("first-level-foreground.png", nil);
		if err != nil {
			fmt.Printf("sf.NewTexturefromFile call failed: %v\n", err);
		}
		l.ForegroundSprite.SetTexture(lForegroundTexture, false);
		lBackgroundSprite, err := sf.NewRectangleShape();
		if err != nil {
			fmt.Printf("sf.newRectangleShape call failed: %v\n", err);
		}
		l.BackgroundSprite = lBackgroundSprite;
		l.BackgroundSprite.SetSize(sf.Vector2f{float32(1920), float32(1080)});
		l.BackgroundSprite.SetPosition(sf.Vector2f{0.0, 0.0});
		lBackgroundTexture, err := sf.NewTextureFromFile("first-level-background.png", nil);
		if err != nil {
			fmt.Printf("sf.NewTexturefromFile call failed: %v\n", err);
		}
		l.BackgroundSprite.SetTexture(lBackgroundTexture, false);
		l.Exists = true;
		this.Level = l;

		var s snowman;
		sSprite, err := sf.NewRectangleShape();
		if err != nil {
			fmt.Printf("sf.NewRectangleShape call failed: %v\n", err);
		}
		s.Sprite = sSprite;
		s.Sprite.SetSize(sf.Vector2f{184.0, 237.5});
		s.Sprite.SetPosition(sf.Vector2f{1307.12, 624.84+50});
		sTexture, err := sf.NewTextureFromFile("snowman.png", nil);
		if err != nil {
			fmt.Printf("sf.NewTextureFromFile call failed. %v\n", err);
		}
		s.Sprite.SetTexture(sTexture, false);
		s.Up = false;
		s.Exists = true;
		this.Snowman = s;
		this.Time = time.Now().UnixNano();
		this.Points = 0;
		this.SnowmanPoints = 0;
		this.Speed = 4000000000;
		fmt.Printf("Initialized gear 1\n");
	} else if this.Gear == 2 {
		var l leveltwo;
		text, _ := sf.NewText(this.Font);
		l.EndText.Text = text;
		if this.Points > this.SnowmanPoints {
			l.EndText.Text.SetString("You have fallen. You made " + strconv.Itoa(this.Points) + " points\nwhile the snowman made only " + strconv.Itoa(this.SnowmanPoints) + " points.\n\nYou won the game!\n\nPress Enter if you want to play again.\nPress Escape to close the game.");
		} else {
			l.EndText.Text.SetString("You have fallen. You made only " + strconv.Itoa(this.Points) + " points\nwhile the snowman made " + strconv.Itoa(this.SnowmanPoints) + " points.\n\nYou lost the game!\n\nPress Enter if you want to play again.\nPress Escape to close the game.");
		}
		l.EndText.Text.SetCharacterSize(24);
		l.EndText.Text.SetColor(sf.Color{0,0,0,255});
		l.EndText.Text.SetPosition(sf.Vector2f{300,300});
		l.EndText.Exists = true;
		l.Exists = true;
		this.LevelTwo = l;
	}
}

func (this *game) Update() {
	for event := this.RenderWindow.PollEvent(); event != nil; event = this.RenderWindow.PollEvent() {
		switch ev := event.(type) {
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyEscape {
					this.RenderWindow.Close();
				}
				if ev.Code == sf.KeyReturn && this.Gear == 0 {
					this.ChangeGearUp();
				}
				if ev.Code == sf.KeyReturn && this.Gear == 2 {
					this.ChangeGearDown();
				}
			case sf.EventClosed:
				this.RenderWindow.Close();
		}
	}
	if this.Gear == 1 {
		if sf.IsMouseButtonPressed(sf.MouseLeft) == true {
			if this.Snowball.Exists == false {
				play := 1 + rand.Intn(4 - 1);
				this.Music.Throws[play].Play();
				var sb snowball;
				sbSprite, err := sf.NewRectangleShape();
				if err != nil {
					fmt.Printf("sf.NewRectangleShape call failed: %v\n", err);
				}
				sb.Sprite = sbSprite;
				sb.Location = this.RenderWindow.MapPixelToCoords(sf.MouseGetPosition(this.RenderWindow), this.RenderWindow.GetDefaultView());
				fmt.Printf("Mouse Location: %#s\n", sb.Location);
				sb.Sprite.SetSize(sf.Vector2f{float32(600), float32(600)});
				sb.Sprite.SetPosition(sf.Vector2f{float32(sb.Location.X-300), float32(sb.Location.Y-300)});
				sbTexture, err := sf.NewTextureFromFile("snowball.png", nil);
				if err != nil {
					fmt.Printf("sf.NewTextureFromFile call failed: %v\n", err);
				}
				sb.Sprite.SetTexture(sbTexture, false);
				sb.Exists = true;
				this.Snowball = sb;
				fmt.Printf("Snowball thrown!\n");
			}
		}
		if this.Snowball.Exists == true {
			size := this.Snowball.Sprite.GetSize();
			newSize := sf.Vector2f{size.X-size.X*0.075, size.Y-size.Y*0.075};
			this.Snowball.Sprite.SetSize(newSize);
			this.Snowball.Sprite.SetOrigin(sf.Vector2f{newSize.X/2, newSize.Y/2});
			this.Snowball.Sprite.Rotate(30.0);
			this.Snowball.Sprite.SetPosition(sf.Vector2f{this.Snowball.Location.X, this.Snowball.Location.Y});
			if size.X < 5 && size.Y < 5 {
				this.Snowball.Exists = false;
			}
			if size.X < 50 && size.Y < 50 {
				sbRect := this.Snowball.Sprite.GetGlobalBounds();
				smRect := this.Snowman.Sprite.GetGlobalBounds();
				test, _ := sbRect.Intersects(smRect);
				if test == true {
					play := 1 + rand.Intn(4 - 1);
					this.Music.Explosions[play].Play();
					sSize := this.Snowman.Sprite.GetSize();
					sPos := this.Snowman.Sprite.GetPosition();
					this.Snowman.Sprite.Move(sf.Vector2f{0, sSize.Y});
					this.Snowman.Up = false;
					fmt.Printf("Snowman shot!\n");
					// add points to the player!
					points := 500 + rand.Intn(1000 - 500);
					fmt.Printf("Points received = %v\n", points);
					this.Points += points;
					fmt.Printf("Total points = %v\n", this.Points);
					this.Speed -= int64(float64(this.Speed) * 0.05);

					this.Snowball.Exists = false;
					text, err := sf.NewText(this.Font);
					if err != nil {
						fmt.Printf("sf.NewText call failed: %v\n", err);
					}
					this.PointsText.Text = text;
					this.PointsText.Text.SetString("+" + strconv.Itoa(points));
					this.PointsText.Text.SetCharacterSize(24);
					this.PointsText.Text.SetColor(sf.Color{0,0,0,255});
					this.PointsText.Text.SetPosition(sPos);
					this.PointsText.Exists = true;

				}
			}
		}
		if this.SnowmanBall.Exists == true {
			size := this.SnowmanBall.Sprite.GetSize();
			newSize := sf.Vector2f{size.X+size.X*0.075, size.Y+size.Y*0.075};
			this.SnowmanBall.Sprite.SetSize(newSize);
			this.SnowmanBall.Sprite.SetOrigin(sf.Vector2f{newSize.X/2, newSize.Y/2});
			this.SnowmanBall.Sprite.Rotate(30.0);
			if newSize.X > 1200 && newSize.Y > 1200 {
				play := 1 + rand.Intn(4 - 1);
				this.Music.Explosions[play].Play();
				this.SnowmanBall.Exists = false;
				// Add snowman hits at you.
				this.SnowmanPoints += 500 + rand.Intn(1000 - 500);
				if this.SnowmanPoints > 5500 {
					// SNOWMAN WON THE GAME YOU FUCKING LOSER!
					this.ChangeGearUp();
					fmt.Printf("SNOWMAN WON THE GAME YOU FUCKING LOSER!\n");
					this.Snowman.Exists = false;
				}
			}
		}
		if this.PointsText.Exists == true {
			OldPosition := this.PointsText.Text.GetPosition();
			if OldPosition.Y > -100 {
				this.PointsText.Text.SetPosition(sf.Vector2f{OldPosition.X, OldPosition.Y-10});
			} else {
				this.PointsText.Exists = false;
			}
		}
		if (this.Time + this.Speed) < time.Now().UnixNano() {
			if this.Snowman.Up == true {
				play := 1 + rand.Intn(4 - 1);
				this.Music.Throws[play].Play();
				// Drop the snowman
				var sb snowball;
				sbSprite, err := sf.NewRectangleShape();
				if err != nil {
					fmt.Printf("sf.NewRectangleShape call failed: %v\n", err);
				}
				sb.Sprite = sbSprite;
				sb.Location = this.RenderWindow.MapPixelToCoords(sf.MouseGetPosition(this.RenderWindow), this.RenderWindow.GetDefaultView());
				fmt.Printf("Mouse Location: %#s\n", sb.Location);
				sb.Sprite.SetSize(sf.Vector2f{float32(50), float32(50)});
				snowmanPosition := this.Snowman.Sprite.GetPosition();
				sb.Sprite.SetPosition(sf.Vector2f{snowmanPosition.X, snowmanPosition.Y});
				sbTexture, err := sf.NewTextureFromFile("snowball.png", nil);
				if err != nil {
					fmt.Printf("sf.NewTextureFromFile call failed: %v\n", err);
				}
				sb.Sprite.SetTexture(sbTexture, false);
				sb.Exists = true;
				this.SnowmanBall = sb;
				fmt.Printf("SnowmanBall thrown!\n");
				sSize := this.Snowman.Sprite.GetSize();
				this.Snowman.Sprite.Move(sf.Vector2f{0, sSize.Y});
				this.Snowman.Up = false;
				fmt.Printf("Drop snowman\n");
			} else {
				fmt.Printf("Raise snowman\n");
				// Change the snowman position.
				random := 1 + rand.Intn(8 - 1);
				fmt.Printf("Snowman random location: %v\n", random);
				switch random {
					case 1:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{112.03, 711.39+50});
					case 2:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{480.16, 757.38+50});
					case 3:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{654.89, 806.07+50});
					case 4:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{886.97, 765.49+50});
					case 5:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{1099.04, 684.34+50});
					case 6:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{1307.12, 624.84+50});
					case 7:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{1399.15, 627.54+50});
					case 8:
						this.Snowman.Sprite.SetPosition(sf.Vector2f{1588.55, 691.11+50});
				}
				// Raise the snowman 
				sSize := this.Snowman.Sprite.GetSize();
				this.Snowman.Sprite.Move(sf.Vector2f{0, -sSize.Y});
				this.Snowman.Up = true;
			}
			this.Time = time.Now().UnixNano();
		}
	}
}

func (this *game) Draw() {
	this.RenderWindow.Clear(sf.Color{255, 255, 255, 255});
	renderStates := sf.DefaultRenderStates();
	// drawing code here.
	if this.Gear == 0 {
		if this.LevelZero.Exists == true {
			if this.LevelZero.WelcomeText.Exists == true {
				this.LevelZero.WelcomeText.Text.Draw(this.RenderWindow, renderStates);
			}
		}
	} else if this.Gear == 1 {
		if this.Level.Exists == true {
			this.RenderWindow.Draw(this.Level.BackgroundSprite, renderStates);
		}
		if this.Snowman.Exists == true {
			this.RenderWindow.Draw(this.Snowman.Sprite, renderStates);
		}
		if this.Level.Exists == true {
			this.RenderWindow.Draw(this.Level.ForegroundSprite, renderStates);
		}
		if this.Snowball.Exists == true {
			this.RenderWindow.Draw(this.Snowball.Sprite, renderStates);
		}
		if this.SnowmanBall.Exists == true {
			this.RenderWindow.Draw(this.SnowmanBall.Sprite, renderStates);
		}
		if this.PointsText.Exists == true {
			this.PointsText.Text.Draw(this.RenderWindow, renderStates);
		}
	} else if this.Gear == 2 {
		if this.LevelTwo.Exists == true {
			if this.LevelTwo.EndText.Exists == true {
				this.LevelTwo.EndText.Text.Draw(this.RenderWindow, renderStates);
			}
		}
	}
	this.RenderWindow.Display();
}

