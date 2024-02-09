FROM golang:1.21-alpine

ADD . /tow_service
WORKDIR /tow_service

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY .env ./

RUN go build -o /go-docker-demo

EXPOSE 8080

CMD [ "/tow_service api" ]