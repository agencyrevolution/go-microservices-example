#Repo Service
Currently having `GET /users/:username/repos` and `GET /users/:username/repos/:reponame`.

##Prerequisites
- etcd (`go install github.com/coreos/etcd`)
- vulcand (`go install github.com/mailgun/vulcand`)

##Run
```
$ VULCAND_ADDR=http://localhost:8182 HOST=localhost [PORT=<port>] go run *.go
```


##Run with Docker
Build:

```
$ docker build -t service.repo .
```

Run:

```
$ docker run -e VULCAND_ADDR=http://172.17.42.1:8182 -e HOST=172.17.42.1 [-e PORT=<port>] -p <port>:<port> service.repo
```
