package cidergo

// ItemType represents the different kinds of items you can play, which has to be specified when using PlayItem or AddItemToQueue
type ItemType string

const (
	ItemTypeSong       ItemType = "songs"
	ItemTypeAlbum      ItemType = "albums"
	ItemTypePlaylist   ItemType = "playlists"
	ItemTypeMusicVideo ItemType = "music-videos"
	ItemTypeStation    ItemType = "stations"
)
