FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/hugomd/cloudflare-ddns/
ADD . /go/src/github.com/hugomd/cloudflare-ddns/
RUN cd /go/src/github.com/hugomd/cloudflare-ddns && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=build-env /go/src/github.com/hugomd/cloudflare-ddns/main /
CMD ["/main"]
