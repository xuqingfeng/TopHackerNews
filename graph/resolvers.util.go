package graph

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/xuqingfeng/TopHackerNews/graph/model"
)

const HN_ITEMS_API = "https://hacker-news.firebaseio.com/v0/item/"

func fetchStoryDetail(id int, r *queryResolver, wg *sync.WaitGroup) error {

	resp, err := http.Get(HN_ITEMS_API + strconv.Itoa(id) + ".json?print=pretty")
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	s := new(model.Story)
	err = json.NewDecoder(resp.Body).Decode(s)
	if err != nil {
		return err
	}
	r.topStories = append(r.topStories, s)
	wg.Done()

	return nil
}
