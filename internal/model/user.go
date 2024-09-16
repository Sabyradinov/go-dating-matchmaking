package model

// UserResponse response to user matching
type UserResponse struct {
	PotentialMatches []UserData
}

// UserData response to user data
type UserData struct {
	UserID    string
	Name      string
	Age       int
	Gender    string
	Interests []string
}
