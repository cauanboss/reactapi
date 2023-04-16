FROM golang:alpine

RUN apk --update add ca-certificates git
RUN export DOCKER_BUILDKIT=0
RUN COMPOSE_DOCKER_CLI_BUILD=0

WORKDIR /usr/app

COPY . .

RUN apk update \
    && apk add git \
    && apk add bash \
    && apk add curl

RUN go mod download 

#RUN go mod tidy  

RUN go mod vendor   

CMD go run .
