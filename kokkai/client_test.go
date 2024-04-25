package kokkai

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	var params = NewParam()
	params.Any("国会議事録")
	params.MaximumRecords(100)
	client := &Client[*HatsugenResult]{
		HTTPClient: http.DefaultClient,
		Get:        GetHatsugen,
		Interval:   time.Duration(1) * time.Second,
	}
	for result := range client.GetAll(context.Background(), params) {
		assert.NoError(t, result.Err)
	}
}
