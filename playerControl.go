package cidergo

type PlayerControlOption string

const (
	PlayerControlOptionPlay         PlayerControlOption = "play"
	PlayerControlOptionPause        PlayerControlOption = "pause"
	PlayerControlOptionTogglePause  PlayerControlOption = "playpause"
	PlayerControlOptionStop         PlayerControlOption = "stop"
	PlayerControlOptionSkip         PlayerControlOption = "next"
	PlayerControlOptionPlayPrevious PlayerControlOption = "previous"
)

func ControlPlayer(option PlayerControlOption) error {
	return postRequestNoJson(string(option))
}
