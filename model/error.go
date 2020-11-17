package model

// Error represent a instance of a error that ocurred at some point in the
// application.
type Error struct {
	ErrorCode   int    `json:"code"`
	Description string `json:"description"`
}
