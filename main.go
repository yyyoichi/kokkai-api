package kokkaiapi

import (
	"context"
	"iter"
)

type Request[T result] struct {
	Client    Client[T]
	Generator Generator
}

var (
	//「会議単位簡易出力」を取得するリクエスト構造体
	DefaultKaniRequest = Request[*KaniResult]{
		Client:    KaniClient,
		Generator: KaniURI,
	}
	//「会議単位出力」を取得するリクエスト構造体
	DefaultKaigiRequest = Request[*KaigiResult]{
		Client:    KaigiClient,
		Generator: KaigiURI,
	}
	//「発言単位出力」を取得するリクエスト構造体
	DefaultHatsugenRequest = Request[*HatsugenResult]{
		Client:    HatsugenClient,
		Generator: HatsugenURI,
	}
)

// 「会議単位簡易出力」で議事録を再帰的に取得します。
//
// 引数に最初の検索条件を指定します。
func IterKaniResult(p Param) iter.Seq2[*KaniResult, error] {
	return IterResult(p, DefaultKaniRequest)
}

// 「会議単位出力」で議事録を再帰的に取得します。
//
// 引数に最初の検索条件を指定します。
func IterKaigiResult(p Param) iter.Seq2[*KaigiResult, error] {
	return IterResult(p, DefaultKaigiRequest)
}

// 「発言単位出力」で議事録を再帰的に取得します。
//
// 引数に最初の検索条件を指定します。
func IterHatsugenResult(p Param) iter.Seq2[*HatsugenResult, error] {
	return IterResult(p, DefaultHatsugenRequest)
}

// 任意のリクエスト構造体から再帰的に議事録を取得します。
func IterResult[T result](p Param, r Request[T]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		nextPosCh := make(chan int, 1)
		ctx, cancel := context.WithCancel(context.Background())

		defer close(nextPosCh)
		defer cancel()
		uriIter := r.Generator.Generate(ctx, p, nextPosCh)
		for val, err := range r.Client.IterRequest(ctx, uriIter) {
			if ok := yield(val, err); !ok {
				break
			}
			p := val.getNextRecordPosition()
			if p > 0 {
				nextPosCh <- p
			} else {
				return
			}
		}
	}
}
