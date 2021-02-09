package main

import (
	"log"

	"github.com/norendren/go-fov/fov"
)

func UpdateMonster(game *Game) {
	l := game.Map.CurrentLevel
	playerPosition := Position{}

	for _, plr := range game.World.Query(game.WorldTags["players"]) {
		pos := plr.Components[position].(*Position)
		playerPosition.X = pos.X
		playerPosition.Y = pos.Y
	}
	for _, result := range game.World.Query(game.WorldTags["monsters"]) {
		pos := result.Components[position].(*Position)
		mon := result.Components[monster].(*Monster)

		monsterSees := fov.New()
		monsterSees.Compute(l, pos.X, pos.Y, 8)
		if monsterSees.IsVisible(playerPosition.X, playerPosition.Y) {
			log.Printf("%s shivers to its bones.", mon.Name)
		}

	}

	game.Turn = PlayerTurn
}
