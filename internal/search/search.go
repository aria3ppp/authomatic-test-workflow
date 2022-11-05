package search

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aria3ppp/watch-server/internal/models"
	"github.com/elastic/go-elasticsearch/v8"
)

type Service interface {
	SearchSerieses(
		ctx context.Context,
		query string,
		from, size int,
	) (hits []*models.Series, totalHits int, err error)
	SearchMovies(
		ctx context.Context,
		query string,
		from, size int,
	) (hits []*models.Film, totalHits int, err error)
}

type ElasticSearch struct {
	client *elasticsearch.Client
}

var _ Service = &ElasticSearch{}

func NewElasticSearch(client *elasticsearch.Client) (*ElasticSearch, error) {
	// check series index exists
	resp, err := client.Indices.Exists([]string{"series"})
	if err != nil {
		return nil, err
	}
	// create series index with mappings
	if resp.StatusCode == http.StatusNotFound {
		resp, err = client.Indices.Create(
			"series",
			client.Indices.Create.WithBody(seriesMappings),
		)
		if err != nil {
			return nil, err
		}
		if resp.IsError() {
			return nil, fmt.Errorf(
				"search.NewElasticSearch: failed creating series index: %s",
				resp,
			)
		}
	}
	// check movie index exists
	resp, err = client.Indices.Exists([]string{"movie"})
	if err != nil {
		return nil, err
	}
	// create movie index with mappings
	if resp.StatusCode == http.StatusNotFound {
		resp, err = client.Indices.Create(
			"movie",
			client.Indices.Create.WithBody(movieMappings),
		)
		if err != nil {
			return nil, err
		}
		if resp.IsError() {
			return nil, fmt.Errorf(
				"search.NewElasticSearch: failed creating movie index: %s",
				resp,
			)
		}
	}
	return &ElasticSearch{client: client}, nil
}

func (e *ElasticSearch) SearchSerieses(
	ctx context.Context,
	query string,
	from, size int,
) (hits []*models.Series, totalHits int, err error) {
	// prepare search query
	searchQuery := fmt.Sprintf(
		`{"query": {"multi_match": {"query": "%s", "fields": ["title", "descriptions"], "fuzziness": "AUTO"}}}`,
		query,
	)
	// search query
	resp, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex("series"),
		e.client.Search.WithBody(strings.NewReader(searchQuery)),
		e.client.Search.WithTrackTotalHits(true),
		e.client.Search.WithFrom(from),
		e.client.Search.WithSize(size),
	)
	if err != nil {
		return nil, 0, err
	}
	if resp.IsError() {
		var em map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&em); err != nil {
			return nil, 0, err
		}
		return nil, 0, fmt.Errorf(
			"[%s] %s: %s",
			resp.Status(),
			em["error"].(map[string]interface{})["type"],
			em["error"].(map[string]interface{})["reason"],
		)
	}
	// decode response body
	type R struct {
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				Source *models.Series `json:"_source"`
			}
		}
	}
	var r R
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, 0, err
	}
	hits = make([]*models.Series, min(size, r.Hits.Total.Value))
	for i, h := range r.Hits.Hits {
		hits[i] = h.Source
	}
	return hits, r.Hits.Total.Value, nil
}

func (e *ElasticSearch) SearchMovies(
	ctx context.Context,
	query string,
	from, size int,
) (hits []*models.Film, totalHits int, err error) {
	// prepare search query
	searchQuery := fmt.Sprintf(
		`{"query": {"multi_match": {"query": "%s", "fields": ["title", "descriptions"], "fuzziness": "AUTO"}}}`,
		query,
	)
	// search query
	resp, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex("movie"),
		e.client.Search.WithBody(strings.NewReader(searchQuery)),
		e.client.Search.WithTrackTotalHits(true),
		e.client.Search.WithFrom(from),
		e.client.Search.WithSize(size),
	)
	if err != nil {
		return nil, 0, err
	}
	if resp.IsError() {
		var em map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&em); err != nil {
			return nil, 0, err
		}
		return nil, 0, fmt.Errorf(
			"[%s] %s: %s",
			resp.Status(),
			em["error"].(map[string]interface{})["type"],
			em["error"].(map[string]interface{})["reason"],
		)
	}
	// decode response body
	type R struct {
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				Source *models.Film `json:"_source"`
			}
		}
	}
	var r R
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, 0, err
	}
	hits = make([]*models.Film, min(size, r.Hits.Total.Value))
	for i, h := range r.Hits.Hits {
		hits[i] = h.Source
	}
	return hits, r.Hits.Total.Value, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// /*
func (e *ElasticSearch) deleteMe() {
	client := e.client

	// create a new index
	client.Indices.Create(
		"new_index",
		// options
		client.Indices.Create.WithErrorTrace(),
		client.Indices.Create.WithTimeout(time.Second),
	)

	// create/update a document in index
	client.Index(
		"index",
		strings.NewReader("document body"),
		// options
		client.Index.WithDocumentID("id"),
		client.Index.WithContext(context.Background()),
	)

	// create a document in index
	client.Create(
		"index",
		"id",
		strings.NewReader("document body"),
		// options
		client.Create.WithContext(context.Background()),
		client.Create.WithHuman(),
	)

	// update a document in index
	client.Update(
		"index",
		"id",
		strings.NewReader("document body"),
		// options
		client.Update.WithContext(context.Background()),
	)

	// search documents
	client.Search(
		client.Search.WithBody(nil),
		client.Search.WithFilterPath([]string{}...),
		client.Search.WithFrom(0),
		client.Search.WithSize(100),
		client.Search.WithTrackTotalHits(true),
	)

	client.Reindex(
		strings.NewReader("body"),
		client.Reindex.WithRequestsPerSecond(0),
		client.Reindex.WithWaitForCompletion(true),
	)
}

// */
