FROM golang:1.13

WORKDIR $GOPATH/src/github.com/er230059/golang-practice
COPY . $GOPATH/src/github.com/er230059/golang-practice
RUN go build

EXPOSE 8000

ENTRYPOINT ["./golang-practice"]
