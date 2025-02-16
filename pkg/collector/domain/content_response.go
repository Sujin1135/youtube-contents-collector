package domain

import "time"

type ContentResponse struct {
	Kind          string    `json:"kind"`
	Etag          string    `json:"etag"`
	NextPageToken *string   `json:"nextPageToken"`
	RegionCode    string    `json:"regionCode"`
	PageInfo      *PageInfo `json:"pageInfo"`
	Items         *[]Item   `json:"items"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type Item struct {
	Kind    string   `json:"kind"`
	Etag    string   `json:"etag"`
	Id      Id       `json:"id"`
	Snippet *Snippet `json:"snippet"`
}

type Id struct {
	Kind      string `json:"kind"`
	ChannelId string `json:"channelId"`
}

type Snippet struct {
	PublishedAt          time.Time   `json:"publishedAt"`
	ChannelId            string      `json:"channelId"`
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	Thumbnails           *Thumbnails `json:"thumbnails"`
	ChannelTitle         string      `json:"channelTitle"`
	LiveBroadcastContent string      `json:"liveBroadcastContent"`
	PublishTime          time.Time   `json:"publishTime"`
}

type ThumbnailItem struct {
	URL string `json:"url"`
}

type Thumbnails struct {
	Default *ThumbnailItem `json:"default"`
	Medium  *ThumbnailItem `json:"medium"`
	High    *ThumbnailItem `json:"high"`
}
