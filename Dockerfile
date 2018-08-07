FROM golang:latest

ENV APP_ROOT /go/src/github.com/dinobi/postit-go
ENV PORT 3006

RUN mkdir -p $APP_ROOT
WORKDIR $APP_ROOT

COPY . $APP_ROOT

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go install

EXPOSE $PORT

ENTRYPOINT /go/bin/postit-go