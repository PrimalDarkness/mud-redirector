FROM golang:1.14 AS build

ENV GOPATH=/go
WORKDIR /proxy
COPY . .

RUN go get -d \
    && go build -o mud-redirector *.go

FROM golang:1.14

COPY --from=build /proxy/mud-redirector /usr/bin/mud-redirector

CMD ["/usr/bin/mud-redirector"]
