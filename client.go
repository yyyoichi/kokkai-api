package kokkaiapi

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"
	"time"
)

var (
	defaultInterval = time.Duration(1 * time.Second)
	KaniClient      = Client[*KaniResult]{
		HttpClient:    http.DefaultClient,
		Interval:      defaultInterval,
		NewResultFunc: func() *KaniResult { return new(KaniResult) },
	}
	KaigiClient = Client[*KaigiResult]{
		HttpClient:    http.DefaultClient,
		Interval:      defaultInterval,
		NewResultFunc: func() *KaigiResult { return new(KaigiResult) },
	}
	HatsugenClient = Client[*HatsugenResult]{
		HttpClient:    http.DefaultClient,
		Interval:      defaultInterval,
		NewResultFunc: func() *HatsugenResult { return new(HatsugenResult) },
	}
)

type Client[T result] struct {
	HttpClient    httpClient
	Interval      time.Duration
	NewResultFunc func() T
}

type result interface {
	getNextRecordPosition() int
}

func (c *Client[T]) IterRequest(ctx context.Context, uriIter iter.Seq[string]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for uri := range uriIter {
			select {
			case <-ctx.Done():
				return
			case <-time.After(c.Interval):
				var val = c.NewResultFunc()
				err := requestApi(c.HttpClient, uri, val)
				if ok := yield(val, err); !ok {
					return
				}
			}
		}
	}
}

type httpClient interface {
	Get(string) (*http.Response, error)
}

func requestApi(client httpClient, uri string, val interface{}) error {
	resp, err := client.Get(uri)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRequestFailed, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: got='%s'", ErrNonOKResponse, resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(val); err != nil {
		return fmt.Errorf("%w: %w", ErrParsingResponse, err)
	}
	return nil
}
