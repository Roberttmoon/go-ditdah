package player

import "code"

type Player struct {
	CharacterSet code.CharSet

	playerType codePlayer
}

type codePlayer interface {
	Play(char string) error
}
