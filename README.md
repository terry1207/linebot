
# go-getting-started

A barebones Go app, which can easily be deployed to Heroku.

This application supports the [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) article - check it out.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.17 or newer and the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) installed.

```sh
$ git clone https://github.com/heroku/go-getting-started.git
$ cd go-getting-started
$ go build -o bin/go-getting-started -v . # or `go build -o bin/go-getting-started.exe -v .` in git bash
github.com/mattn/go-colorable
gopkg.in/bluesuncorp/validator.v5
golang.org/x/net/context
github.com/heroku/x/hmetrics
github.com/gin-gonic/gin/render
github.com/manucorporat/sse
github.com/heroku/x/hmetrics/onload
github.com/gin-gonic/gin/binding
github.com/gin-gonic/gin
github.com/heroku/go-getting-started
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku main
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)

CHANNEL_SECRET: 5888804cf864224c0039290df79bbef0
CHANNEL_TOKEN:  tNraN7tLdbZPoW/GE79Blv7ivJ10n2xlyrMvGvtk4DHrf9Jyb/lJV9ZSelL+PIq5qGjsHVxKr3UyFLmQUXeVvpPZ2BqByxB8XoTNGQCBsCW0vLjFfoFULsOhQRh64sPK7wnggWgRtiKZ/M5bKgRagwdB04t89/1O/w1cDnyilFU=
DATABASE_URL:   postgres://wkkckoevwhzsti:756a9057eb6ed63db56541deeb46335868aad00e5f26e0d942d5f021f7df062d@ec2-52-71-69-66.compute-1.amazonaws.com:5432/d2qfr5f81g7vlm
