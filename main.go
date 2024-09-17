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
	DefaultKaniRequest = Request[*KaniResult]{
		Client:    KaniClient,
		Generator: KaniURI,
	}
	DefaultKaigiRequest = Request[*KaigiResult]{
		Client:    KaigiClient,
		Generator: KaigiURI,
	}
	DefaultHatsugenRequest = Request[*HatsugenResult]{
		Client:    HatsugenClient,
		Generator: HatsugenURI,
	}
)

func IterKaniResult(p Param) iter.Seq2[*KaniResult, error] {
	return IterResult(p, DefaultKaniRequest)
}

func IterResult[T result](p Param, r Request[T]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		nextPosCh := make(chan int)
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
