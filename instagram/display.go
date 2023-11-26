package instagram

import (
	"context"
	"fmt"
	"time"
)

type DisplayService service

// Media represents an image, video, or album
type Media struct {
	Caption        string    `json:"caption"`           // Caption is the caption of the media
	ID             string    `json:"id"`                // ID is the media UUID (unique identifier)
	IsSharedToFeed bool      `json:"is_shared_to_feed"` // IsSharedToFeed is whether the media is shared to the feed
	MediaURL       string    `json:"media_url"`         // MediaURL is the url of the media
	PermaLink      string    `json:"permalink"`         // PermaLink is the url of the media
	ThumbnailURL   string    `json:"thumbnail_url"`     // ThumbnailURL is the url of the thumbnail
	Timestamp      time.Time `json:"timestamp"`         // Timestamp is the time the media was posted
	Username       string    `json:"username"`          // Username is the username of the user who posted the media
	MediaType      string    `json:"media_type"`        // MediaType is the type of media (IMAGE, VIDEO, or CAROUSEL_ALBUM)
}

// User represents an instagram user
type User struct {
	AccountType string  `json:"account_type"` // AccountType is the type of account (personal or business)
	ID          string  `json:"id"`           // ID is the users UUID (unique identifier)
	MediaCount  int     `json:"media_count"`  // MediaCount is the number of media the user has
	Username    string  `json:"username"`     // Username is the username of the user
	Edges       []Media `json:"media"`        // Edges is a list of media
}

// GetUser will get a single user by their id
func (u *DisplayService) GetUser(ctx context.Context, userID string) (*User, error) {
	req, err := u.client.NewRequest("GET", fmt.Sprintf("%s/%s", u.client.InstagramVersion, userID), nil)
	if err != nil {
		return nil, err
	}

	var user *User
	_, err = u.client.Do(ctx, req, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
