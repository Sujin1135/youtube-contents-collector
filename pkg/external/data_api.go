package external

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type DataAPI interface {
	Search(searchType SearchType, query string) (string, error)
}

type dataAPI struct {
	unit    int
	apiKey  string
	baseUrl string
}

type SearchType int

const (
	Channel SearchType = iota
	Content
)

func (s SearchType) String() string {
	switch s {
	case Channel:
		return "Channel"
	case Content:
		return "Content"
	default:
		return "Unknown"
	}
}

func (receiver *dataAPI) Search(searchType SearchType, query string) (string, error) {
	endpoint, err := url.Parse(receiver.baseUrl + "/search")
	if err != nil {
		return "", err
	}

	endpoint.RawQuery = receiver.generateSearchParams(searchType, query)
	response, err := http.Get(endpoint.String())

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := string(responseData)
	fmt.Println(data)

	return data, nil
}

func (receiver *dataAPI) generateSearchParams(searchType SearchType, query string) string {
	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("key", receiver.apiKey)
	params.Add("q", query)
	params.Add("type", searchType.String())
	return params.Encode()
}

func NewDataAPI() DataAPI {
	return &dataAPI{
		unit:    10000,
		apiKey:  os.Getenv("YOUTUBE_DATA_API_KEY"),
		baseUrl: os.Getenv("YOUTUBE_DATA_API_BASE_URL"),
	}
}
