package graph

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/xuqingfeng/TopHackerNews/graph/model"
)

const HN_ITEMS_API = "https://hacker-news.firebaseio.com/v0/item/"

var hnClient = &http.Client{Timeout: 10 * time.Second}

func fetchStoryDetail(id int) (*model.Story, error) {
	resp, err := hnClient.Get(HN_ITEMS_API + strconv.Itoa(id) + ".json?print=pretty")
	if err != nil {
		log.Printf("err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	s := new(model.Story)
	err = json.NewDecoder(resp.Body).Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
