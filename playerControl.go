package cidergo

// PlayerControlOption represents the different actions you can perform on the music player
type PlayerControlOption string

const (
	PlayerControlOptionPlay         PlayerControlOption = "play"
	PlayerControlOptionPause        PlayerControlOption = "pause"
	PlayerControlOptionTogglePause  PlayerControlOption = "playpause"
	PlayerControlOptionStop         PlayerControlOption = "stop"
	PlayerControlOptionSkip         PlayerControlOption = "next"
	PlayerControlOptionPlayPrevious PlayerControlOption = "previous"
)

// ControlPlayer performs the given PlayerControlOption on the music player
func ControlPlayer(option PlayerControlOption) error {
	return postRequestNoJson(string(option))
}
