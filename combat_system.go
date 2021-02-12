package main

import (
	"fmt"

	"github.com/bytearena/ecs"
)

func AttackSystem(g *Game, attackerPosition *Position, defenderPosition *Position) {
	var attacker *ecs.QueryResult = nil
	var defender *ecs.QueryResult = nil

	//Get the attacker and defender if either is a player
	for _, playerCombatant := range g.World.Query(g.WorldTags["players"]) {
		pos := playerCombatant.Components[position].(*Position)

		if pos.IsEqual(attackerPosition) {
			//This is the attacker
			attacker = playerCombatant
		} else if pos.IsEqual(defenderPosition) {
			//This is the defender
			defender = playerCombatant
		}
	}

	//Get the attacker and defender if either is a monster
	for _, cbt := range g.World.Query(g.WorldTags["monsters"]) {
		pos := cbt.Components[position].(*Position)

		if pos.IsEqual(attackerPosition) {
			//This is the attacker
			attacker = cbt
		} else if pos.IsEqual(defenderPosition) {
			//This is the defender
			defender = cbt
		}

	}
	//If we somehow don't have an attacker or defender, just leave
	if attacker == nil || defender == nil {
		return
	}
	//Grab the required information
	defenderArmor := defender.Components[armor].(*Armor)
	defenderHealth := defender.Components[health].(*Health)
	defenderName := defender.Components[name].(*Name).Label

	attackerWeapon := attacker.Components[meleeWeapon].(*MeleeWeapon)
	attackerName := attacker.Components[name].(*Name).Label

	//Roll a d10 to hit
	toHitRoll := GetDiceRoll(10)

	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
		//It's a hit!
		damageRoll := GetRandomBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)

		damageDone := damageRoll - defenderArmor.Defense
		//Let's not have the weapon heal the defender
		if damageDone < 0 {
			damageDone = 0
		}
		defenderHealth.CurrentHealth -= damageDone
		fmt.Printf("%s swings %s at %s and hits for %d health.\n", attackerName, attackerWeapon.Name, defenderName, damageDone)

		if defenderHealth.CurrentHealth <= 0 {
			fmt.Printf("%s has died!\n", defenderName)
			if defenderName == "Player" {
				fmt.Printf("Game Over!\n")
				g.Turn = GameOver
			}
			g.World.DisposeEntity(defender.Entity)
		}

	} else {
		fmt.Printf("%s swings %s at %s and misses.\n", attackerName, attackerWeapon.Name, defenderName)
	}
}
