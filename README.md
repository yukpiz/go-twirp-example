# go-twirp-example

## Quickstart

```bash
$ realize start
$ curl -H "Content-Type: application/json" -XPOST -d "{\"user_id\": 1}" "http://localhost:9991/twirp/user.UserAPI/GetUser"
```


## Development

```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go get -u github.com/twitchtv/twirp/protoc-gen-twirp

# Generates user pb
$ protoc --proto_path=./proto --twirp_out=./pb/ --go_out=./pb/ ./proto/user/*.proto
# Generates staff pb
$ protoc --proto_path=./proto --twirp_out=./pb/ --go_out=./pb/ ./proto/staff/*.proto
```
