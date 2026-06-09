# elasticx

Elasticsearch 客户端管理工具，基于 olivere/elastic/v7，支持多数据源与索引管理。

## 安装

```bash
go get github.com/go-xuan/elasticx
```

## 快速开始

在 `conf/elastic.yaml` 中配置：

```yaml
source: "default"
enable: true
url: "http://127.0.0.1:9200"
username: "elastic"
password: "elastic"
indices:
  - "users"
```

```go
import "github.com/go-xuan/elasticx"

func main() {
    elasticx.Initialize()
    client := elasticx.GetESClient("default")
    client.Search().Index("users").Do(ctx)
}
```

## 主要功能

- **多数据源** — 支持同时连接多个 ES 集群
- **索引管理** — 自动创建索引、Mapping 配置
- **配置驱动** — 配合 configx 自动从 nacos / 本地文件加载
