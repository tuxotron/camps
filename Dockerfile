FROM golang:1.12 as build

WORKDIR /go/src/app
COPY camp.go .

RUN go get -d -v ./...
RUN go install -v ./...

FROM gcr.io/distroless/base

COPY --from=build /go/bin/app /
COPY camps.json /
EXPOSE 8080

CMD ["/app"]