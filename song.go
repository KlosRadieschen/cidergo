package main

type Artwork struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type PlayParams struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
}

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
