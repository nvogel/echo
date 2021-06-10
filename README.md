# Echo server

Http server for testing.


# Environment variable

## PORT

configure on which port the service must listen (default 8080)


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

````
docker run --rm -ti -v "$PWD":/usr/src/myapp ruby:2.5 /bin/bash
gem install travis
```