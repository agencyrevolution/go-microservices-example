#Repo Service
Currently having `GET /users/:username/repos` and `GET /users/:username/repos/:reponame`.

##Prerequisites
- etcd (`go install github.com/coreos/etcd`)
- vulcand (`go install github.com/mailgun/vulcand`)

##Run
```
VULCAND_ADDR=http://localhost:8182 VULCAND_VER=v2 HOST=localhost go run *.go [--port <port>]
```
