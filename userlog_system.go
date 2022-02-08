package main

import (
	"image/color"
	"log"

	"github.com/RAshkettle/rrogue/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var userLogImg *ebiten.Image = nil
var err error = nil
var mplusNormalFont font.Face = nil
var lastText []string = make([]string, 0, 5)

func ProcessUserLog(g *Game, screen *ebiten.Image) {
	if userLogImg == nil {
		userLogImg, _, err = ebitenutil.NewImageFromFile("assets/UIPanel.png")
		if err != nil {
			log.Fatal(err)
		}
	}
	if mplusNormalFont == nil {
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Fatal(err)
		}

		const dpi = 72
		mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    16,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	gd := NewGameData()

	uiLocation := (gd.ScreenHeight - gd.UIHeight) * gd.TileHeight
	var fontX = 16
	var fontY = uiLocation + 24
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0.), float64(uiLocation))
	screen.DrawImage(userLogImg, op)
	tmpMessages := make([]string, 0, 5)
	anyMessages := false

	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.AttackMessage != "" {
			tmpMessages = append(tmpMessages, messages.AttackMessage)
			anyMessages = true
			//fmt.Printf(messages.AttackMessage)
			messages.AttackMessage = ""
		}
	}
	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.DeadMessage != "" {
			tmpMessages = append(tmpMessages, messages.DeadMessage)
			anyMessages = true
			//fmt.Printf(messages.DeadMessage)
			messages.DeadMessage = ""
			g.World.DisposeEntity(m.Entity)
		}
		if messages.GameStateMessage != "" {
			tmpMessages = append(tmpMessages, messages.GameStateMessage)
			anyMessages = true
			//No need to clear, it's all over
		}

	}
	if anyMessages {
		lastText = tmpMessages
	}
	for _, msg := range lastText {
		if msg != "" {
			text.Draw(screen, msg, mplusNormalFont, fontX, fontY, color.White)
			fontY += 16
		}
	}

}
