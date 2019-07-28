package code

type CharSet struct {
	Chars map[string]*Char
}

type Char struct {
	ASCII           string
	MorseCodeString string

	MorseCode []MorseChar
}

type MorseChar int

const (
	UnkonwChar MorseChar = iota
	DitChar
	DahChar
	SpaceChar
)
