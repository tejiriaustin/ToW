FROM golang:1.21-alpine

RUN go install github.com/beego/bee/v2@latest

ADD . /tow_service
WORKDIR /tow_service

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY .env ./

RUN go build -o /tow_service

EXPOSE 8080

CMD [ "/tow_service api" ]