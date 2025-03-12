package main

type Artwork struct {
	URL        string `json:"url"`
	BGColor    string `json:"bgColor"`
	Height     int    `json:"height"`
	Width      int    `json:"width"`
	TextColor1 string `json:"textColor1"`
	TextColor2 string `json:"textColor2"`
	TextColor3 string `json:"textColor3"`
}

type PlayParams struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
}

type Previews struct {
	URL string `json:"url"`
}

type Song struct {
	AlbumName                 string     `json:"albumName"`
	ArtistName                string     `json:"artistName"`
	Artwork                   Artwork    `json:"artwork"`
	AudioLocale               string     `json:"audioLocale"`
	AudioTraits               []string   `json:"audioTraits"`
	ComposerName              string     `json:"composerName"`
	CurrentPlaybackProgress   float64    `json:"currentPlaybackProgress"`
	CurrentPlaybackTime       float64    `json:"currentPlaybackTime"`
	DiscNumber                int        `json:"discNumber"`
	DurationInMillis          int        `json:"durationInMillis"`
	EndTime                   int        `json:"endTime"`
	GenreNames                []string   `json:"genreNames"`
	HasLyrics                 bool       `json:"hasLyrics"`
	HasTimeSyncedLyrics       bool       `json:"hasTimeSyncedLyrics"`
	IsAppleDigitalMaster      bool       `json:"isAppleDigitalMaster"`
	IsMasteredForItunes       bool       `json:"isMasteredForItunes"`
	IsPlaying                 bool       `json:"isPlaying"`
	IsVocalAttenuationAllowed bool       `json:"isVocalAttenuationAllowed"`
	ISRC                      string     `json:"isrc"`
	Kind                      string     `json:"kind"`
	Name                      string     `json:"name"`
	PlayParams                PlayParams `json:"playParams"`
	Previews                  []Previews `json:"previews"`
	ReleaseDate               string     `json:"releaseDate"`
	RemainingTime             float64    `json:"remainingTime"`
	SongID                    string     `json:"songId"`
	StartTime                 float64    `json:"startTime"`
	Status                    string     `json:"status"`
	TrackNumber               int        `json:"trackNumber"`
	URL                       string     `json:"url"`
}
