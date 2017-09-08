FROM golang:1.9
WORKDIR /go/src/github.com/schiob/AAAS/
RUN go get -v github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/schiob/AAAS/main .
COPY --from=0 /go/src/github.com/schiob/AAAS/swagger/swagger.json ./swagger/swagger.json
EXPOSE 8080
CMD ["./main"]
