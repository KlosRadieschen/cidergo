package cidergo

import "encoding/json"

// Artwork holds information about an item's artwork. Completeness depends on the item itself, and its container
type Artwork struct {
	URL        string `json:"url"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	TextColor1 string `json:"textColor1,omitempty"`
	TextColor2 string `json:"textColor2,omitempty"`
	TextColor3 string `json:"textColor3,omitempty"`
	BgColor    string `json:"bgColor,omitempty"`
	HasP3      bool   `json:"hasP3,omitempty"`
}

// PlayParams holds information about an item's parameters. Completeness depends on the item itself, and its container
type PlayParams struct {
	ID          string `json:"id"`
	Kind        string `json:"kind"`
	Format      string `json:"format,omitempty"`
	StationHash string `json:"stationHash,omitempty"`
	HasDrm      bool   `json:"hasDrm,omitempty"`
	MediaType   int    `json:"mediaType,omitempty"`
}

// Song holds information about a song
type Song struct {
	AlbumName  string   `json:"albumName"`
	ArtistName string   `json:"artistName"`
	Name       string   `json:"name"`
	GenreNames []string `json:"genreNames"`

	DurationInMillis int    `json:"durationInMillis"`
	ReleaseDate      string `json:"releaseDate"`
	TrackNumber      int    `json:"trackNumber"`
	DiscNumber       int    `json:"discNumber"`

	AudioLocale string   `json:"audioLocale"`
	AudioTraits []string `json:"audioTraits"`

	URL        string     `json:"url"`
	Artwork    Artwork    `json:"artwork"`
	PlayParams PlayParams `json:"playParams"`

	ISRC                      string `json:"isrc"`
	HasLyrics                 bool   `json:"hasLyrics"`
	IsAppleDigitalMaster      bool   `json:"isAppleDigitalMaster"`
	IsVocalAttenuationAllowed bool   `json:"isVocalAttenuationAllowed"`
	IsMasteredForItunes       bool   `json:"isMasteredForItunes"`

	CurrentPlaybackTime float64 `json:"currentPlaybackTime"`
	RemainingTime       float64 `json:"remainingTime"`

	InFavorites bool `json:"inFavorites"`
	InLibrary   bool `json:"inLibrary"`
	ShuffleMode int  `json:"shuffleMode"`
	RepeatMode  int  `json:"repeatMode"`
}

// CurrentSong returns the song that is currently playing
func CurrentSong() (Song, error) {
	jsonData, err := jsonRequest("now-playing")
	if err != nil {
		return Song{}, err
	}

	var data struct {
		Song Song `json:"info"`
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return Song{}, err
	}

	return data.Song, nil
}
