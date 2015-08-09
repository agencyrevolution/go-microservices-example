#Vulcand Initializer
This is command to initialize vulcand for the whole system.

##Run
Install `etcd` and `vulcand`:

```
$ go install github.com/coreos/etcd
$ go install github.com/mailgun/vulcand
```
Start `etcd` and `vulcand`:

```
$ etcd --data-dir ~/.etcd
$ vulcand --etcd=http://localhost:4001
```
So we are having vulcand work on our machine (normally `:8182` is API Interface), now we run the initializer:

```
$ VULCAND_ADDR=http://localhost:8182 go run *.go
```
##Run with Docker
Run `etcd`:

```
$ docker run -d -p 4001:4001 coreos/etcd
```

Run `vulcand`:

In the example, we are using the lastest version of vulcand (that is currently not on docker hub). We have to build it first:

```
$ go get github.com/mailgun/vulcand
$ cd $GOPATH/src/github.com/mailgun/vulcand
$ docker build -t vulcand .
...
$ # Now we can run vulcand
$ docker run -d -p 8182:8182 -p 3004:3004 vulcand
```

Now we run the initializer:

```
$ VULCAND_ADDR=http://<docker-host>:8182 go run *.go
```
