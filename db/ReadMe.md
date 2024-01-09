```shell
goctl model mysql datasource --dir ./internal/model --table user --cache true --url "root:9012@tcp(127.0.0.1:3306)/infinite_user"

goctl model mysql datasource --dir ./internal/model --table user --cache true --url "root:9012@tcp(127.0.0.1:3306)/infinite_article"

goctl model mysql datasource --dir ./internal/model --table user --cache true --url "root:9012@tcp(127.0.0.1:3306)/infinite_like"

goctl model mysql datasource --dir ./internal/model --table user --cache true --url "root:9012@tcp(127.0.0.1:3306)/infinite_tag"

goctl model mysql datasource --dir ./internal/model --table user --cache true --url "root:9012@tcp(127.0.0.1:3306)/infinite_comment"
```
