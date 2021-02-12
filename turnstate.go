package main

type TurnState int

const (
	BeforePlayerAction = iota
	PlayerTurn
	MonsterTurn
	GameOver
)

func GetNextState(state TurnState) TurnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn
	case PlayerTurn:
		return MonsterTurn
	case MonsterTurn:
		return BeforePlayerAction
	case GameOver:
		return GameOver
	default:
		return PlayerTurn
	}
}
