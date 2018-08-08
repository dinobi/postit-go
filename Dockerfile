# Using a multistep container build to achieve a smalller
# app image

FROM golang:1.10.3 as builder

ENV APP_ROOT /go/src/github.com/dinobi/postit-go
ENV PORT 3006

WORKDIR $APP_ROOT

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# instruct our build script to statically compile our
# app with all libraries from linux go os built in:
# ...
# We’re disabling cgo which gives us a static binary.
# We’re also setting the OS to Linux (in case someone
# builds this on a Mac or Windows) and the -a flag means
# to rebuild all the packages we’re using, which means all
# the imports will be rebuilt with cgo disabled.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

EXPOSE $PORT

# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM alpine:latest

# Security related package for handling SSL requests, good to have.
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
COPY --from=builder /go/src/github.com/dinobi/postit-go .

CMD [ "./postit-go" ]