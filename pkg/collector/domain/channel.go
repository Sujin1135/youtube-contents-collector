package domain

import "time"

type Channel struct {
	Id          *int32             `json:"id"`
	ExternalId  string             `json:"external_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Published   time.Time          `json:"published"`
	Thumbnails  *ChannelThumbnails `json:"thumbnails"`
	Created     time.Time          `json:"created"`
	Modified    time.Time          `json:"modified"`
	Deleted     *time.Time         `json:"deleted"`
}

type ChannelThumbnails struct {
	Default *ThumbnailDetail `json:"default"`
	Medium  *ThumbnailDetail `json:"medium"`
	High    *ThumbnailDetail `json:"high"`
}
