FROM golang:alpine
WORKDIR /usr/app

COPY . .

RUN apk update \
    && apk add git \
    && apk add bash \
    && apk add curl \
    && go mod tidy \    
    && go mod vendor \
    && go get github.com/cosmtrek/air


# RUN curl -fLo ~/.air https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air \
#     && chmod +x ~/.air

# RUN ~/.air init

# CMD [ "~/.air ", "-d" ]

CMD [ "go", "run", "." ]
