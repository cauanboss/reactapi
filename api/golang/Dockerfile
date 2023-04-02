FROM golang:alpine
WORKDIR /usr/app

COPY . .

RUN apk update \
    && apk add git \
    && apk add bash \
    && apk add curl \
    && go mod tidy \    
    && go mod vendor

CMD [ "go", "run", "." ]
