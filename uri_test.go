package kokkaiapi

import (
	"context"
	"iter"
	"strconv"
	"strings"
	"testing"
)

// ConstructURIを繰り返し呼んでもURIが変わらないことを確認する。
func TestConstructURIDoesNotMutateParam(t *testing.T) {
	p := NewParam()
	p.Any("科学技術")
	p.SessionFrom(209)
	uri1 := ConstructURI(KaniURI, p)
	uri2 := ConstructURI(KaniURI, p)
	if uri1 != uri2 {
		t.Errorf("ConstructURI changed the Param:\n first: %s\nsecond: %s", uri1, uri2)
	}
	if n := strings.Count(uri2, "recordPacking="); n != 1 {
		t.Errorf("recordPacking appears %d times in %q, want 1", n, uri2)
	}
}

// レスポンスをjsonとしてパースする都合上、recordPackingはjsonに固定されることを確認する。
func TestConstructURIForcesRecordPackingJSON(t *testing.T) {
	p := NewParam()
	p.RecordPacking("xml")
	uri := ConstructURI(KaniURI, p)
	if strings.Contains(uri, "recordPacking=xml") {
		t.Errorf("recordPacking=xml remained in %q, but the client parses responses as json", uri)
	}
	if n := strings.Count(uri, "recordPacking=json"); n != 1 {
		t.Errorf("recordPacking=json appears %d times in %q, want 1", n, uri)
	}
}

// 全件取得でページをまたいでもrecordPackingが重複せず、開始位置が進むことを確認する。
func TestGenerateDoesNotAccumulateRecordPacking(t *testing.T) {
	p := NewParam()
	p.Any("科学技術")
	g := Generator(KaniURI)
	nextPosCh := make(chan int, 1)
	next, stop := iter.Pull(g.Generate(context.Background(), p, nextPosCh))
	defer stop()
	for i := 0; i < 3; i++ {
		uri, ok := next()
		if !ok {
			t.Fatalf("generator stopped at page %d", i+1)
		}
		if n := strings.Count(uri, "recordPacking="); n != 1 {
			t.Errorf("page %d: recordPacking appears %d times in %q, want 1", i+1, n, uri)
		}
		if i >= 1 {
			if want := "startRecord=" + strconv.Itoa(i*10+1); !strings.Contains(uri, want) {
				t.Errorf("page %d: %q does not contain %q", i+1, uri, want)
			}
		}
		nextPosCh <- (i+1)*10 + 1
	}
}
