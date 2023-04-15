package main

type TurnState int

const (
	BeforePlayerAction TurnState = iota
	PlayerTurn
	MonsterTurn
)

func NextState(state TurnState) TurnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn
	case PlayerTurn:
		return MonsterTurn
	case MonsterTurn:
		return BeforePlayerAction
	default:
		return PlayerTurn
	}
}
