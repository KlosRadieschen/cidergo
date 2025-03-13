package cidergo

import "encoding/json"

// QueueItem represents one item inside the queue
type QueueItem struct {
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	AssetURL     string          `json:"assetURL"`
	HlsMetadata  json.RawMessage `json:"hlsMetadata"`
	Flavor       string          `json:"flavor"`
	Song         Song            `json:"attributes"`
	PlaybackType int             `json:"playbackType"`
	Container    Container       `json:"_container"`
	Context      Context         `json:"_context"`
	State        State           `json:"_state"`
	SongID       string          `json:"_songId"`
	Assets       []Asset         `json:"assets"`
	KeyURLs      KeyURLs         `json:"keyURLs"`
}

// Preview holds preview URL info.
type Preview struct {
	URL string `json:"url"`
}

// Container holds container info like radio station details
type Container struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Href       string              `json:"href"`
	Attributes ContainerAttributes `json:"attributes"`
	Name       string              `json:"name"`
}

// ContainerAttributes holds all the extra attributes for the container
type ContainerAttributes struct {
	RequiresSubscription bool       `json:"requiresSubscription"`
	IsLive               bool       `json:"isLive"`
	Kind                 string     `json:"kind"`
	RadioUrl             string     `json:"radioUrl"`
	MediaKind            string     `json:"mediaKind"`
	Name                 string     `json:"name"`
	Artwork              Artwork    `json:"artwork"`
	URL                  string     `json:"url"`
	PlayParams           PlayParams `json:"playParams"`
}

// Context holds context info
type Context struct {
	FeatureName string `json:"featureName"`
}

// State holds the current state
type State struct {
	Current int `json:"current"`
}

// Asset holds detailed asset data
type Asset struct {
	Flavor      string   `json:"flavor"`
	URL         string   `json:"URL"`
	DownloadKey string   `json:"downloadKey"`
	ArtworkURL  string   `json:"artworkURL"`
	FileSize    int      `json:"file-size"`
	MD5         string   `json:"md5"`
	Chunks      Chunks   `json:"chunks"`
	Metadata    Metadata `json:"metadata"`
	PreviewURL  string   `json:"previewURL,omitempty"`
}

// Chunks holds info about chunk sizes and hashes
type Chunks struct {
	ChunkSize int      `json:"chunkSize"`
	Hashes    []string `json:"hashes"`
}

// Metadata holds additional metadata for each item
type Metadata struct {
	ComposerId          string `json:"composerId"`
	GenreId             int    `json:"genreId"`
	Copyright           string `json:"copyright"`
	Year                int    `json:"year"`
	SortArtist          string `json:"sort-artist"`
	IsMasteredForItunes bool   `json:"isMasteredForItunes"`
	VendorId            int    `json:"vendorId"`
	ArtistId            string `json:"artistId"`
	Duration            int    `json:"duration"`
	DiscNumber          int    `json:"discNumber"`
	ItemName            string `json:"itemName"`
	TrackCount          int    `json:"trackCount"`
	Xid                 string `json:"xid"`
	BitRate             int    `json:"bitRate"`
	FileExtension       string `json:"fileExtension"`
	SortAlbum           string `json:"sort-album"`
	Genre               string `json:"genre"`
	Rank                int    `json:"rank"`
	SortName            string `json:"sort-name"`
	PlaylistId          string `json:"playlistId"`
	SortComposer        string `json:"sort-composer"`
	Comments            string `json:"comments"`
	TrackNumber         int    `json:"trackNumber"`
	ReleaseDate         string `json:"releaseDate"`
	Kind                string `json:"kind"`
	PlaylistArtistName  string `json:"playlistArtistName"`
	Gapless             bool   `json:"gapless"`
	ComposerName        string `json:"composerName"`
	DiscCount           int    `json:"discCount"`
	SampleRate          int    `json:"sampleRate"`
	PlaylistName        string `json:"playlistName"`
	Explicit            int    `json:"explicit"`
	ItemId              string `json:"itemId"`
	S                   int    `json:"s"`
	Compilation         bool   `json:"compilation"`
	ArtistName          string `json:"artistName"`
}

// KeyURLs holds the key URLs for playback
type KeyURLs struct {
	HLSKeyCertURL   string `json:"hls-key-cert-url"`
	HLSKeyServerURL string `json:"hls-key-server-url"`
	WidevineCertURL string `json:"widevine-cert-url"`
}
