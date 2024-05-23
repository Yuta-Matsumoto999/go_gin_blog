FROM golang:1.22.3-alpine

WORKDIR /workspace
COPY ./backend/src .

RUN apk upgrade --update && apk --no-cache add git

RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]