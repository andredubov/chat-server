package model

// Chat model for service layer
type Chat struct {
	ID      int64
	Name    string
	UserIDs []int64
}
