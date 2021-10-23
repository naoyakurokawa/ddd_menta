FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/naoyakurokawa/ddd_menta
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build ./cmd/main.go

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/naoyakurokawa/ddd_menta/main /main
EXPOSE 8080
CMD ["./main"]
