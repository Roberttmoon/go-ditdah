package player

type textPlayer struct{}

var _ codePlayer = &textPlayer{}

func (textPlayer) Play(string) error {

	return nil
}
