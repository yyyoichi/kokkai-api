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
	// 「会議単位簡易出力」を取得するクライアント
	KaniClient = Client[*KaniResult]{
		HttpClient:    http.DefaultClient,
		Interval:      defaultInterval,
		NewResultFunc: func() *KaniResult { return new(KaniResult) },
	}
	// 「会議単位出力」を取得するクライアント
	KaigiClient = Client[*KaigiResult]{
		HttpClient:    http.DefaultClient,
		Interval:      defaultInterval,
		NewResultFunc: func() *KaigiResult { return new(KaigiResult) },
	}
	// 「発言単位出力」を取得するクライアント
	HatsugenClient = Client[*HatsugenResult]{
		HttpClient:    http.DefaultClient,
		Interval:      defaultInterval,
		NewResultFunc: func() *HatsugenResult { return new(HatsugenResult) },
	}
)

// 国会議事録APIにリクエストを行う構造体。
type Client[T result] struct {
	// httpリクエストを実行する構造体。
	//
	// 下記の、`httpClient`を満たす構造体が必要。
	//
	// type httpClient interface {
	// 	Get(string) (*http.Response, error)
	// }
	//
	HttpClient httpClient
	// APIにアクセスするインターバル。
	//
	// [国会会議録検索システム　検索用APIの仕様](https://kokkai.ndl.go.jp/api.html)では、アクセスの間隔に数秒間を置くことが推奨されている。
	//
	// デフォルトで1秒間としている。
	Interval time.Duration
	// レスポンス構造体を初期化する関数。
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
