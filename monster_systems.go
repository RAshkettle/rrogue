package main

import (
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
		//mon := result.Components[monster].(*Monster)

		monsterSees := fov.New()
		monsterSees.Compute(l, pos.X, pos.Y, 8)
		if monsterSees.IsVisible(playerPosition.X, playerPosition.Y) {

			if pos.GetManhattanDistance(&playerPosition) == 1 {
				//The monster is right next to the player.  Just smack him down
				AttackSystem(game, pos, &playerPosition)
				if result.Components[health].(*Health).CurrentHealth <= 0 {
					//this monster is dead
					//clear the tile
					t := l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)]
					t.Blocked = false
				}

			} else {
				astar := AStar{}
				path := astar.GetPath(l, pos, &playerPosition)
				if len(path) > 1 {
					nextTile := l.Tiles[l.GetIndexFromXY(path[1].X, path[1].Y)]
					if !nextTile.Blocked {
						l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
						pos.X = path[1].X
						pos.Y = path[1].Y
						nextTile.Blocked = true
					}
				}
			}

		}

	}

	game.Turn = PlayerTurn
}
