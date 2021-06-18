# Echo server

Http server for testing.

# Environment variables

## PORT

configure on which port the service must listen (default 8080)

## COLOR

define the color of the app that will appear on the root page **/**

(default: white)

## KIND

define the kind of the app that will appear on the root page **/**

(default: dog)


# Dev

```
mkdir -p $GOPATH/github.com/nvogel
cd $GOPATH/github.com/nvogel
git clone https://github.com/nvogel/echo.git
cd echo
dep ensure
go build
```

# CI

https://docs.travis-ci.com/user/deployment/releases/

```
docker run --rm -ti -v "$PWD":/usr/src/myapp ruby:2.5 /bin/bash
gem install travis
```

# Notes

- https://www.alexedwards.net/blog/golang-response-snippets
- https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/