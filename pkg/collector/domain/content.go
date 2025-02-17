package domain

import "time"

type Content struct {
	Id               int32              `json:"id"`
	ChannelId        int32              `json:"channel_id"`
	ExternalId       string             `json:"external_id"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	Published        time.Time          `json:"published"`
	ContentThumbnail *ContentThumbnails `json:"content_thumbnail"`
	Created          time.Time          `json:"created"`
	Modified         time.Time          `json:"modified"`
	Deleted          *time.Time         `json:"deleted"`
}

type ContentThumbnails struct {
	Default *ThumbnailDetail `json:"default"`
	Medium  *ThumbnailDetail `json:"medium"`
	High    *ThumbnailDetail `json:"high"`
}
