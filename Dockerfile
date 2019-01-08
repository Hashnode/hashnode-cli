FROM golang:1.11 AS build
WORKDIR /go/src/github.com/hashnode/hashnode-cli

RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

COPY cmd cmd
COPY pkg pkg
RUN CGO_ENABLED=0 GOOS=linux GOFLAGS=-ldflags=-w go build -o /go/bin/hashnode -ldflags=-s -v github.com/hashnode/hashnode-cli

FROM alpine:3.8 AS final
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/hashnode /bin/hashnode
ENTRYPOINT ["hashnode"]
