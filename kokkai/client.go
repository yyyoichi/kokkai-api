package kokkai

import (
	"context"
	"time"
)

type result interface {
	NextRecordPosition() int
}
type Client[T result] struct {
	HTTPClient
	Get func(HTTPClient, Params) T
	// 処理間隔
	//
	// > 機械的なアクセスを行う場合、多重リクエストは避けてください。
	// また、データを取得し終えてから数秒程度空けて次のリクエストを行うようにしてください。
	// https://kokkai.ndl.go.jp/api.html より引用
	Interval time.Duration
}

// [params]にヒットするすべてのデータを取得する。
func (c *Client[T]) GetAll(ctx context.Context, params Params) <-chan T {
	paramsCh := make(chan Params)
	resultCh := make(chan T)
	respCh := make(chan T)
	go func() {
		defer close(resultCh)
		for {
			select {
			case <-ctx.Done():
				return
			case p, ok := <-paramsCh:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case resultCh <- c.Get(c.HTTPClient, p):
					select {
					case <-ctx.Done():
						return
					case <-time.After(c.Interval):
					}
				}
			}

		}
	}()
	go func() {
		defer close(paramsCh)
		defer close(respCh)

		paramsCh <- params
		for {
			select {
			case <-ctx.Done():
				return
			case result, ok := <-resultCh:
				if !ok {
					return
				}
				respCh <- result
				next := result.NextRecordPosition()
				if next == 0 {
					// end contents
					return
				}
				// has next
				params.ResetStartRecord(next)
				paramsCh <- params
			}
		}
	}()
	return respCh
}
