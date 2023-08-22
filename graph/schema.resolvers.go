package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/xuqingfeng/HackerNewsTopStories/graph/model"
)

// TopStories is the resolver for the topStories field.
func (r *queryResolver) TopStories(ctx context.Context, offset *int, limit *int) ([]*model.Story, error) {

	r.topStories = make([]*model.Story, 0)

	// get all top stories' ids
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		log.Printf("err: %v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		ids := new([]int)
		err = json.NewDecoder(resp.Body).Decode(ids)
		// log.Printf("ids: %v", ids)
		if err != nil {
			return nil, err
		}

		var wg sync.WaitGroup
		for i, id := range *ids {
			if i < *offset {
				continue
			}
			if i >= *offset+*limit {
				break
			}
			// get every story's detail concurrently
			wg.Add(1)
			go fetchNewsDetail(id, r, &wg)
		}
		wg.Wait()
	}

	return r.topStories, nil
}
func fetchNewsDetail(id int, r *queryResolver, wg *sync.WaitGroup) error {

	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(id) + ".json?print=pretty")
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

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
