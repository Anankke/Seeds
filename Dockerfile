FROM golang as builder
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/Seeds
COPY . .
RUN  dep ensure \
   && go build -o seeds-linux-amd64 src/main/main.go
FROM debian:stable
COPY --from=builder /go/src/Seeds/seeds-linux-amd64 /bin/seeds
ENTRYPOINT ["/bin/seeds"]
