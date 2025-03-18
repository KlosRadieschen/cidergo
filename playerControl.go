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

// Pause is a convenience function equivalent to ControlPlayer(PlayerControlOptionPause).
// It pauses whatever is currently playing (but it doesn't unpause if playback is already paused)
func Pause() error {
	return ControlPlayer(PlayerControlOptionPause)
}

// Play is a convenience function equivalent to ControlPlayer(PlayerControlOptionPlay).
// It unpauses whatever is currently playing (but it doesn't pause if playback is already playing)
func Play() error {
	return ControlPlayer(PlayerControlOptionPlay)
}

// TogglePause is a convenience function equivalent to ControlPlayer(PlayerControlOptionTogglePause).
// It unpauses a paused song or pauses and unpaused song.
func TogglePause() error {
	return ControlPlayer(PlayerControlOptionTogglePause)
}

// Stop is a convenience function equivalent to ControlPlayer(PlayerControlOptionStop).
// It completely stops the current playback
func Stop() error {
	return ControlPlayer(PlayerControlOptionStop)
}

// Skip is a convenience function equivalent to ControlPlayer(PlayerControlOptionSkip).
// It skips the current song
func Skip() error {
	return ControlPlayer(PlayerControlOptionSkip)
}

// PlayPrevious is a convenience function equivalent to ControlPlayer(PlayerControlOptionPlayPrevious)
// It goes back to the song that played before the currently playing song
func PlayPrevious() error {
	return ControlPlayer(PlayerControlOptionPlayPrevious)
}
