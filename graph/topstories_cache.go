package graph

import (
	"cmp"
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/xuqingfeng/TopHackerNews/graph/model"
)

const (
	topStoriesCacheTTL       = 5 * time.Minute
	maxStoryFetchConcurrency = 32
)

type topStoriesCache struct {
	mu        sync.RWMutex
	refreshMu sync.Mutex
	stories   []*model.Story
	expiresAt time.Time
}

func (c *topStoriesCache) get() ([]*model.Story, error) {
	c.mu.RLock()
	if time.Now().Before(c.expiresAt) && c.stories != nil {
		stories := c.stories
		c.mu.RUnlock()
		return stories, nil
	}
	c.mu.RUnlock()

	c.refreshMu.Lock()
	defer c.refreshMu.Unlock()

	c.mu.RLock()
	if time.Now().Before(c.expiresAt) && c.stories != nil {
		stories := c.stories
		c.mu.RUnlock()
		return stories, nil
	}
	c.mu.RUnlock()

	stories, err := fetchAndSortTopStories()
	if err != nil {
		c.mu.RLock()
		stale := c.stories
		c.mu.RUnlock()
		if len(stale) > 0 {
			log.Printf("topstories refresh failed, serving stale cache: %v", err)
			return stale, nil
		}
		return nil, err
	}

	c.mu.Lock()
	c.stories = stories
	c.expiresAt = time.Now().Add(topStoriesCacheTTL)
	c.mu.Unlock()

	return stories, nil
}

func fetchAndSortTopStories() ([]*model.Story, error) {
	resp, err := hnClient.Get(HN_TOPSTORIES_API)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	ids := new([]int)
	if err := json.NewDecoder(resp.Body).Decode(ids); err != nil {
		return nil, err
	}

	stories := make([]*model.Story, 0, len(*ids))
	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxStoryFetchConcurrency)

	for _, id := range *ids {
		wg.Add(1)
		go func(storyID int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			story, err := fetchStoryDetail(storyID)
			if err != nil {
				return
			}
			mu.Lock()
			stories = append(stories, story)
			mu.Unlock()
		}(id)
	}
	wg.Wait()

	slices.SortFunc(stories, func(a, b *model.Story) int {
		return cmp.Compare(b.Score, a.Score)
	})

	return stories, nil
}
