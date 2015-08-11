#React Service
This is front-end service.

##Prerequisites
- node.js
- `service.repo` started
- `service.user` started

##Build
```
$ npm install
...
$ npm run build
```


##Run
```
$ API_ADDR=http://localhost:3004 [PORT=3003] npm start
```


##Run with Docker
Build:

```
$ docker build -t react .
```

Run:

```
$ docker run -e API_ADDR=http://172.17.42.1:3004 -e HOST=172.17.42.1 [-e PORT=<port>] -p <port>:<port> react
```
