# Go covaerage in devcontainer

```sh
go test -v -cover -coverprofile=cover.out .
go tool cover -html=cover.out -o cover.html
```

add extension

```json
{
    "extentions" : [
        "ritwickdey.liveserver"
    ]
}
```

```sh
go get -tool github.com/air-verse/air@v1.62.0
```