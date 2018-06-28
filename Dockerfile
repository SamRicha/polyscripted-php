FROM golang
WORKDIR /go/src/github.com/polyverse/polyscripting-php
COPY . .
WORKDIR /go/src/github.com/polyverse/polyscripting-php/transformer
RUN go get -v ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
WORKDIR ../scrambler
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM ubuntu

RUN apt-get update && apt-get -y upgrade
RUN apt-get install -y \
      git \
      make \
      autoconf \
      gcc \
      re2c \
      bison \
      libxml2-dev \
      vim \
      ccache

COPY scripts /php/
COPY --from=0 /go/src/github.com/polyverse/polyscripting-php/transformer/transformer /php/
COPY --from=0 /go/src/github.com/polyverse/polyscripting-php/scrambler/scrambler /php/
WORKDIR /php
RUN git clone https://github.com/php/php-src.git

