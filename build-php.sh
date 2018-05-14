apt-get update && apt-get -y upgrade
apt-get install -y \
    git \
    make \
    autoconf \
    gcc \
    re2c \
    bison \
    libxml2-dev \
    vim

git clone --depth 1 https://github.com/php/php-src.git
cd php-src

./buildconf
./configure
make

