package elasticx

import (
	"context"
	"encoding/json"
	"io"

	"github.com/olivere/elastic/v7"
)

// NewWriter 创建ES日志写入器
func NewWriter(source, index string) io.Writer {
	if Initialized() {
		client := GetESClient(source)
		ctx := context.Background()
		if exist, err := client.IndexExists(index).Do(ctx); err != nil || !exist {
			_, _ = client.CreateIndex(index).Do(ctx)
		}
		return &Writer{
			index:  index,
			client: client,
		}
	}
	return nil
}

// Writer 日志写入
type Writer struct {
	index  string
	client *elastic.Client
}

func (w *Writer) Write(bytes []byte) (int, error) {
	go func() {
		var body interface{}
		if err := json.Unmarshal(bytes, &body); err == nil {
			_, _ = w.client.Index().Index(w.index).BodyJson(body).Do(context.Background())
		}
	}()
	return 0, nil
}
