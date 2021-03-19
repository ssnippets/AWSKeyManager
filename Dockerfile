FROM golang:1.16-buster as builder
WORKDIR /go/src
COPY go.sum go.mod main.go ui.go . 
RUN go get -d -v "github.com/manifoldco/promptui" "gopkg.in/ini.v1" \
  && CGO_ENABLED=0 GOOS=linux go build -a -o app . \
  && CGO_ENABLED=0 GOOS=darwin go build -a -o app.darwin .

FROM scratch
COPY --chown=0:0 --from=builder /go/src/app app.linux
COPY --from=builder /go/src/app.darwin app.mac
CMD ["/app.linux"]
