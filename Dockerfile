FROM golang AS build-stage
LABEL build_stage=true
LABEL maintainer="Dora <me@dora1998.net>"
ADD . /go/src/feed-api

WORKDIR /go/src/feed-api
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -o /go/bin/feed-api

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build-stage /go/bin/feed-api /usr/local/bin/feed-api
ENTRYPOINT ["/usr/local/bin/feed-api"]