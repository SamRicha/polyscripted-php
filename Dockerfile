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

COPY build-php.sh /php/
COPY build-scrambled.sh /php/
COPY tests/ /php/tests
COPY phpLexAndYak/ /php/phpReset
COPY Scrambler /php/
COPY Transformer /php/
WORKDIR /php
RUN git clone https://github.com/php/php-src.git

